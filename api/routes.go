package api

import (
	handler "easytrady-backend/api/Handler"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {

	// usuarios

	e.POST("/criarusuario", func(c echo.Context) error {
		return handler.PostUsuario(c)
	})

	e.PUT("/atualizarusuario/:id", func(c echo.Context) error {
		return handler.UpdateUsuario(c)
	})

	e.GET("/usuarios", func(c echo.Context) error {
		return handler.GetAllUsuarios(c)
	})

	e.POST("/login", func(c echo.Context) error {
		return handler.Login(c)
	})

	// fim usuarios

	// produtos

	e.GET("/produtos", func(c echo.Context) error {
		return handler.GetAllProdutos(c)
	})

	e.POST("/criarproduto", func(c echo.Context) error {
		return handler.PostProduto(c)
	})

	e.PUT("/atualizarproduto/:id", func(c echo.Context) error {
		return handler.UpdateProduto(c)
	})

	e.DELETE("/deletarproduto/:id", func(c echo.Context) error {
		return handler.DeleteProduto(c)
	})

	// fim produtos

}
