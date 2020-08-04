package beego

import (
	"log"
	"net/http"
)

var (
	BeeApp *App
)

func init() {
	BeeApp = NewApp()
}

type App struct {
	Handlers *ControllerRegister
	Server   *http.Server
}

func NewApp() *App {
	cr := NewControllerRegister()
	app := &App{Handlers: cr, Server: &http.Server{}}
	return app
}

type MiddleWare func(http.Handler) http.Handler

func (app *App) Run(mws ...MiddleWare) {
	app.Server.Addr = ":9090"
	app.Server.Handler = app.Handlers
	err := app.Server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func Router(rootpath string, c ControllerIntrerface) *App {
	BeeApp.Handlers.Add(rootpath, c)
	return BeeApp
}
