package app

import (
	"log"
	"umbrellaTest/internal/app/endpoint"
	"umbrellaTest/internal/app/mw"
	"umbrellaTest/internal/app/service"

	"github.com/labstack/echo/v4"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() *App {
	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	a.echo.Use(mw.CheckRole)

	a.echo.GET("/timer", a.e.Timer)

	return a
}

func (a *App) Run() {
	err := a.echo.Start(":4000")
	if err != nil {
		log.Fatal(err)
	}
}
