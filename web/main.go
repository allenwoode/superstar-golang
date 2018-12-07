package main

import (
	"awesome/bootstrap"
	"awesome/web/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Superstar database", "Quincy")
	app.Bootstrap()

	app.Configure(routes.Configure)
	return app
}
func main() {
	//api := iris.New()
	//api.Logger().SetLevel("debug")

	//api.Use(recover.New())

	//api.Handle("GET", "/", func(ctx iris.Context) {
	//	ctx.HTML("<h1>Welcome</h1>")
	//})

	//api.RegisterView(iris.HTML("D:/go/src/awesome/web/views", ".html").Reload(true))
	//
	//api.Get("/", func(ctx iris.Context) {
	//	ctx.View("index.html", map[string]string{"Title": "首页"})
	//	//return mvc.View{
	//	//	Name: "index.html",
	//	//	Data: iris.Map{
	//	//		"Title": "首页",
	//	//	}
	//	//}
	//})
	//
	//api.Get("/ping", func(ctx iris.Context) {
	//	ctx.WriteString("pong")
	//})
	//
	//api.Get("/hello", func(ctx iris.Context) {
	//	ctx.JSON(iris.Map{"message": "Hello Iris!"})
	//})
	//
	//api.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
	app := newApp()
	app.Logger().SetLevel("debug")
	app.Listen(":8080")
}
