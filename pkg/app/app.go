package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lucasmmo/setup-go/pkg/services"
)

func Setup() *echo.Echo {
	e := echo.New()
	e.GET("/:name", func(c echo.Context) error {
		s := services.SayHello(c.Param("name"))
		return c.String(http.StatusOK, s)
	})
	return e
}
