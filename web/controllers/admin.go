package controllers

import (
	"awesome/model"
	"awesome/service"
	"awesome/web/viewmodels"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
)

type AdminController struct {
	Ctx     iris.Context
	Service service.SuperstarService
}

func (c *AdminController) Get() mvc.Result {

	return mvc.View{
		Name: "admin/index.html",
		Data: iris.Map{
			"title": "管理后台",
			"data":  c.Service.GetAll(),
		},
		Layout: "admin/layout.html",
	}
}

func (c *AdminController) GetEdit() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	if err != nil {
		return mvc.View{
			Name: "404.html",
		}
	}
	return mvc.View{
		Name: "admin/edit.html",
		Data: iris.Map{
			"title": "管理后台",
			"data":  c.Service.Get(id),
		},
		Layout: "admin/layout.html",
	}
}

func (c *AdminController) PostSave() mvc.Result {
	var star viewmodels.Star
	err := c.Ctx.ReadForm(&star)
	if err != nil {
		log.Fatal(err)
	}
	if star.Id > 0 {
		c.Service.Update(&star, []string{"name_zh", "name_en"})
	} else {
		c.Service.Create(&star)
	}
	return mvc.Response{
		Path: "/admin/",
	}
}

func (c *AdminController) PostDelete() mvc.Result {
	var star model.Star
	err := c.Ctx.ReadForm(&star)
	if err != nil {
		log.Fatal(err)
	}
	if star.Id > 0 {
		c.Service.Delete(int(star.Id))
	}

	return mvc.Response{
		Path: "/admin/",
	}
}
