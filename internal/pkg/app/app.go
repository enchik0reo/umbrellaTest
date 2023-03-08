package app

import (
	"github.com/enchik0reo/umbrellaTest/internal/app/endpoint"
	"github.com/enchik0reo/umbrellaTest/internal/app/mw"
	"github.com/enchik0reo/umbrellaTest/internal/app/service"
	"github.com/enchik0reo/umbrellaTest/internal/config"

	"github.com/labstack/echo/v4"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
	c    *config.Config
}

func New() *App {
	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	a.c = config.GetConfig()

	a.echo.GET("/timer", a.e.Timer, mw.CheckRole)

	return a
}

func (a *App) Run() {
	err := a.echo.Start(a.c.Port)
	if err != nil {
		panic(err)
	}
}
