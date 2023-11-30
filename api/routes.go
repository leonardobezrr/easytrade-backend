package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// criar todas as rotas aqui
func SetupRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
