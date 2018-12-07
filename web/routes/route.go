package routes

import (
	"awesome/bootstrap"
	"awesome/service"
	"awesome/web/controllers"
	"github.com/kataras/iris/mvc"
)

func Configure(b *bootstrap.Bootstrapper) {
	superstarSevice := service.NewSuperstarService()

	index := mvc.New(b.Party("/"))
	index.Register(superstarSevice)
	index.Handle(new(controllers.IndexController))

	admin := mvc.New(b.Party("/admin"))
	//admin.Router.Use()
	admin.Register(superstarSevice)
	admin.Handle(new(controllers.AdminController))

	api := mvc.New(b.Party("/api"))
	api.Register(superstarSevice)
	api.Handle(new(controllers.ApiController))
}
