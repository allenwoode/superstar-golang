package service

import (
	"awesome/dao"
	"awesome/db"
	"awesome/model"
)

type SuperstarService interface {
	GetAll() []model.Star
	Get(id int) *model.Star
	Delete(id int) error
	Update(star *model.Star, colums []string) error
	Create(star *model.Star) error

	Search(nameen string) []model.Star
}

type superstarService struct {
	dao *dao.SuperstarDao
}

func NewSuperstarService() *superstarService {
	return &superstarService{
		dao: dao.NewSuperstarDao(db.NewMasterEngine()),
	}
}

func (s *superstarService) GetAll() []model.Star {
	return s.dao.GetAll()
}

func (s *superstarService) Get(id int) *model.Star {
	return s.dao.Get(id)
}

func (s *superstarService) Update(star *model.Star, colums []string) error {
	return s.dao.Update(star, colums)
}

func (s *superstarService) Create(star *model.Star) error {
	return s.dao.Create(star)
}

func (s *superstarService) Delete(id int) error {
	return s.dao.Delete(id)
}

func (s *superstarService) Search(nameen string) []model.Star {
	return s.dao.Search(nameen)
}
