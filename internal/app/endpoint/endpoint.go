package endpoint

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Service interface {
	DaysLeft() int64
}
type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Status(c echo.Context) error {

	s := fmt.Sprintf("Days left to next year: %d", e.s.DaysLeft())

	err := c.String(http.StatusOK, s)
	if err != nil {
		return err
	}
	return nil
}
