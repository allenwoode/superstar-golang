package service

import (
	"awesome/dao"
	"awesome/db"
	"awesome/model"
	"awesome/web/viewmodels"
)

type SuperstarService interface {
	GetAll() []viewmodels.Star
	Get(id int) *viewmodels.Star
	Delete(id int) error
	Update(star *viewmodels.Star, colums []string) error
	Create(star *viewmodels.Star) error

	Search(nameen string) []viewmodels.Star
}

type superstarService struct {
	dao *dao.SuperstarDao
}

func NewSuperstarService() *superstarService {
	return &superstarService{
		dao: dao.NewSuperstarDao(db.NewMasterEngine()),
	}
}

func (s *superstarService) GetAll() []viewmodels.Star {
	var list []viewmodels.Star
	stars := s.dao.GetAll()
	//var stars []model.Star
	//stars = append(stars, model.Star{1, "ddddd.com", "中国", "CHN", time.Now(), time.Now()})
	//stars = append(stars, model.Star{1, "ddddd.com", "中国", "CHN", time.Now(), time.Now()})
	//stars = append(stars, model.Star{1, "ddddd.com", "中国", "CHN", time.Now(), time.Now()})
	for _, star := range stars {
		data := viewmodels.Star{
			star.Id,
			star.NameZh,
			star.NameEn,
			star.Avatar,
		}
		list = append(list, data)
	}
	return list
}

func (s *superstarService) Get(id int) *viewmodels.Star {
	star := s.dao.Get(id)
	return &viewmodels.Star{star.Id, star.NameZh, star.NameEn, star.Avatar}
}

func (s *superstarService) Update(star *viewmodels.Star, colums []string) error {
	var updateStar *model.Star
	updateStar.Id = star.Id
	updateStar.Avatar = star.Avatar
	updateStar.NameEn = star.NameEn
	updateStar.NameZh = star.NameZh

	return s.dao.Update(updateStar, colums)
}

func (s *superstarService) Create(star *viewmodels.Star) error {
	var createStar *model.Star
	createStar.Avatar = star.Avatar
	createStar.NameEn = star.NameEn
	createStar.NameZh = star.NameZh

	return s.dao.Create(createStar)
}

func (s *superstarService) Delete(id int) error {
	return s.dao.Delete(id)
}

func (s *superstarService) Search(nameen string) []viewmodels.Star {
	var list []viewmodels.Star
	stars := s.dao.Search(nameen)
	for _, star := range stars {
		data := viewmodels.Star{
			star.Id,
			star.NameZh,
			star.NameEn,
			star.Avatar,
		}
		list = append(list, data)
	}
	return list
}
