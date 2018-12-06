package config

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

var (
	MasterDBConfig = DBConfig{"localhost", 3306, "root", "root", "superstar"}
	SlaveDBConfig  = DBConfig{"localhost", 3306, "root", "root", ""}

	DriveName = "mysql"
)
