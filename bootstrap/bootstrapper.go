package bootstrap

import (
	"awesome/config"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"time"
)

type Bootstrapper struct {
	*iris.Application
	AppName      string
	AppOwner     string
	AppSpawnDate time.Time
	Sessions     *sessions.Sessions
}

type Configurator func(bootstrapper *Bootstrapper)

func New(appName, appOwner string, cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		AppName:      appName,
		AppOwner:     appOwner,
		AppSpawnDate: time.Now(),
		Application:  iris.New(),
	}
	for _, cfg := range cfgs {
		cfg(b)
	}
	return b
}

func (b *Bootstrapper) SetupViews(viewsDir string) {
	htmlEngine := iris.HTML(viewsDir, ".html")
	htmlEngine.Reload(false)
	htmlEngine.AddFunc("FromUnixtimeShort", func(t int) string {
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(config.SysTimeformShort)
	})
	htmlEngine.AddFunc("FromUnixtime", func(t int) string {
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(config.SysTimeform)
	})
	b.RegisterView(htmlEngine)
}

func (b *Bootstrapper) SetupSessions(expires time.Duration,
	cookieHashKey, cookieBlockKey []byte) {
	b.Sessions = sessions.New(sessions.Config{
		Cookie:  "SECRET_SESS_COOKIE_" + b.AppName,
		Expires: expires,
		//Encoding:securecookie.New(cookieHashKey, cookieBlockKey),
	})
}

func (b *Bootstrapper) SetupErrorHandlers() {
	b.OnAnyErrorCode(func(ctx iris.Context) {
		err := iris.Map{
			"api":     b.AppName,
			"status":  ctx.GetStatusCode(),
			"message": ctx.Values().GetString("message"),
		}
		if jsonOuput := ctx.URLParamExists("json"); jsonOuput {
			ctx.JSON(err)
			return
		}

		ctx.ViewData("Err", err)
		ctx.ViewData("title", "Error")
		ctx.View("shared/error.html")
	})
}

const (
	StaticAsserts = "./public"
	Favicon       = "favicon.ico"
)

func (b *Bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}

func (b *Bootstrapper) Bootstrap() *Bootstrapper {
	b.SetupViews("./views")
	b.SetupSessions(24*time.Hour, []byte("hashkey"), []byte("hashkey1"))
	b.SetupErrorHandlers()

	// ...
	return b
}

func (b *Bootstrapper) Listen(addr string, cfgs ...iris.Configurator) {
	b.Run(iris.Addr(addr), cfgs...)
}
