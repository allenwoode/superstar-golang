package controllers

import (
	"awesome/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type IndexController struct {
	Ctx     iris.Context
	Service service.SuperstarService
}

func (c *IndexController) Get() mvc.Result {

	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"title": "球星库",
			"data":  c.Service.GetAll(),
		},
	}
}

func (c *IndexController) GetBy(id int) mvc.Result {
	if id < 1 {
		return mvc.Response{
			Path: "/",
		}
	}
	return mvc.View{
		Name: "info.html",
		Data: iris.Map{
			"title": "球星",
			"data":  c.Service.Get(id),
		},
	}
}
