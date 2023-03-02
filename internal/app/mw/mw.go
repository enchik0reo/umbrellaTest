package mw

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

const roleAdmin = "admin"

func CheckRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := logrus.New()

		val := c.Request().Header.Get("User-Role")

		if strings.Contains(strings.ToLower(val), roleAdmin) {
			log.Println("red button user detected")
		}

		err := next(c)
		if err != nil {
			return err
		}

		return nil
	}
}
