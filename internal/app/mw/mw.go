package mw

import (
	"strings"

	"github.com/enchik0reo/umbrellaTest/pkg/logging"

	"github.com/labstack/echo/v4"
)

const roleAdmin = "admin"

func CheckRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		lgr := logging.InitLogger()

		val := c.Request().Header.Get("User-Role")

		if strings.Contains(strings.ToLower(val), roleAdmin) {
			lgr.Info("red button user detected")
		}

		err := next(c)
		if err != nil {
			return err
		}

		return nil
	}
}
