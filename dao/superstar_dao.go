package dao

import (
	"awesome/model"
	"github.com/go-xorm/xorm"
	"log"
)

type SuperstarDao struct {
	engine *xorm.Engine
}

func NewSuperstarDao(engine *xorm.Engine) *SuperstarDao {
	return &SuperstarDao{engine: engine}
}

func (dao *SuperstarDao) Get(id int) *model.Star {
	star := &model.Star{}
	ok, err := dao.engine.ID(id).Get(star)
	if ok && err == nil {
		return star
	}
	//star.Id = 0
	return star
}

func (dao *SuperstarDao) GetAll() []model.Star {
	var stars []model.Star
	err := dao.engine.Desc("id").Find(&stars)
	if err != nil {
		log.Println(err)
		return nil
	}
	return stars
}

func (dao *SuperstarDao) Update(star *model.Star, colums []string) error {
	_, err := dao.engine.Id(star.Id).MustCols(colums...).Update(star)
	return err
}

func (dao *SuperstarDao) Create(star *model.Star) error {
	_, err := dao.engine.Insert(star)
	return err
}

func (dao *SuperstarDao) Delete(id int) error {
	//var star *model.Star
	//star.Id = int64(id)
	_, err := dao.engine.Delete(&model.Star{Id: id})
	return err
}

func (dao *SuperstarDao) Search(nameen string) []model.Star {
	var stars []model.Star
	err := dao.engine.Where("name_en=?", nameen).Find(&stars)
	if err != nil {
		log.Println(err)
		return nil
	}
	return stars
}
