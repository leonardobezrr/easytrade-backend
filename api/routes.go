package api

import (
	handler "easytrady-backend/api/Handler"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/criarusuario", func(c echo.Context) error {
		return handler.PostUsuario(c)
	})

	e.PUT("/atualizarusuario/:id", func(c echo.Context) error {
		return handler.UpdateUsuario(c)
	})
	
	e.POST("/criarproduto", func(c echo.Context) error {
		return handler.PostProduto(c)
	})

	e.POST("/login", func(c echo.Context) error {
		return handler.Login(c)
	})

	e.GET("/usuarios", func(c echo.Context) error {
		return handler.GetAllUsuarios(c)
	})
}
