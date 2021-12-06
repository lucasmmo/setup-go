package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lucasmmo/setup-go/pkg/hello"
)

func Start() *echo.Echo {
	e := echo.New()
	SetupRoutes(e)
	return e
}

func SetupRoutes(e *echo.Echo) {
	e.GET("/:name", func(c echo.Context) error {
		name := c.Param("name")
		msg := hello.Say(name)
		return c.String(http.StatusOK, msg)
	})
}
