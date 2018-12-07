package controllers

import (
	"awesome/model"
	"awesome/service"
	"awesome/web/viewmodels"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
)

type ApiController struct {
	Ctx     iris.Context
	Service service.SuperstarService
}

func (c *ApiController) Get() mvc.Result {

	//c.Ctx.JSON(c.Service.GetAll())
	return mvc.Response{
		Object: c.Service.GetAll(),
	}
}

func (c *ApiController) GetBy(id int) mvc.Result {
	if id < 1 {
		return mvc.Response{
			Path: "/api/",
		}
	}
	data := c.Service.Get(id)
	if data.Id == 0 {
		return mvc.Response{
			Object: viewmodels.Response{
				iris.StatusNotFound,
				"数据不存在",
				nil,
				nil,
			},
		}
	}
	return mvc.Response{
		Object: data,
	}
}

func (c *ApiController) GetSearch() mvc.Result {
	name := c.Ctx.URLParam("nameen")

	return mvc.Response{
		Object: c.Service.Search(name),
	}
}

//func (c *ApiController) GetEdit() mvc.Result {
//	id, err := c.Ctx.URLParamInt("id")
//	if err != nil {
//		return mvc.View{
//			Name: "404.html",
//		}
//	}
//	return mvc.View{
//		Name: "admin/edit.html",
//		Data: iris.Map{
//			"Title": "管理后台",
//			"Data":  c.Service.Get(id),
//		},
//		//Layout: "admin/layout.html",
//	}
//}
//
func (c *ApiController) PostSave() mvc.Result {
	var star viewmodels.Star
	err := c.Ctx.ReadJSON(&star)
	if err != nil {
		log.Fatal(err)
	}
	if star.Id > 0 {
		c.Service.Update(&star, []string{"avatar", "name_zh", "name_en"})
	} else {
		log.Println("Insert Star:", star)
		c.Service.Create(&star)
	}
	return mvc.Response{
		Object: map[string]string{"status": "ok", "message": "成功"},
	}
}

func (c *ApiController) PostDelete() mvc.Result {
	var star model.Star
	err := c.Ctx.ReadJSON(&star)
	if err != nil {
		log.Fatal(err)
	}
	if star.Id < 1 {
		return mvc.Response{
			Object: viewmodels.Response{iris.StatusInternalServerError, "操作失败",
				nil, map[string]string{"message": "数据id不存在"}},
		}
	}
	go c.Service.Delete(int(star.Id))
	return mvc.Response{
		Object: viewmodels.Response{iris.StatusOK, "成功", nil, nil},
	}
}
