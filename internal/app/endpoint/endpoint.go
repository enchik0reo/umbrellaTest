package endpoint

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Service interface {
	DaysBefore() int64
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Timer(c echo.Context) error {

	d := e.s.DaysBefore()

	s := fmt.Sprintf("Days before 2025: %d", d)

	err := c.String(http.StatusOK, s)
	if err != nil {
		return err
	}

	return nil
}
