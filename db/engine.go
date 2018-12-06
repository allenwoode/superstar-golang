package db

import (
	"awesome/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"sync"
)

var (
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
	lock         sync.Mutex
)

func NewMasterEngine() *xorm.Engine {
	if masterEngine != nil {
		return masterEngine
	}
	lock.Lock()
	defer lock.Unlock()
	if masterEngine != nil {
		return masterEngine
	}

	conf := config.MasterDBConfig
	driverSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", conf.User, conf.Password, conf.Host, conf.Port, conf.DBName)
	engine, err := xorm.NewEngine(config.DriveName, driverSource)
	if err != nil {
		log.Fatal("NewMasterEngine err:", err)
		return nil
	}
	masterEngine = engine
	return masterEngine
}

func NewSlaveEngine() *xorm.Engine {
	if slaveEngine != nil {
		return slaveEngine
	}
	lock.Lock()
	defer lock.Unlock()
	if slaveEngine != nil {
		return slaveEngine
	}
	conf := config.MasterDBConfig
	driverSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", conf.User, conf.Password, conf.Host, conf.Port, conf.DBName)
	engine, err := xorm.NewEngine(config.DriveName, driverSource)
	if err != nil {
		log.Fatal("NewSlaveEngine err:", err)
		return nil
	}
	slaveEngine = engine
	return slaveEngine
}
