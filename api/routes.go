package api

import (
	handler "easytrady-backend/api/Handler"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {

	// usuarios

	e.POST("/usuarios/criar", func(c echo.Context) error {
		return handler.PostUsuario(c)
	})

	e.PUT("/usuarios/atualizar/:id", func(c echo.Context) error {
		return handler.UpdateUsuario(c)
	})

	e.DELETE("/usuarios/deleter/:id", func(c echo.Context) error {
		return handler.DeleteUsuario(c)
	})

	e.GET("/usuarios/listar", func(c echo.Context) error {
		return handler.GetAllUsuarios(c)
	})

	e.POST("/login", func(c echo.Context) error {
		return handler.Login(c)
	})

	// fim usuarios

	// produtos

	e.GET("/produtos/listar", func(c echo.Context) error {
		return handler.GetAllProdutos(c)
	})

	e.GET("/produtos/listar/:usuarioID", func(c echo.Context) error {
		return handler.GetProdutosByUsuarioID(c)
	})

	e.POST("/produtos/criar", func(c echo.Context) error {
		return handler.PostProduto(c)
	})

	e.PUT("/produtos/atualizar/:id", func(c echo.Context) error {
		return handler.UpdateProduto(c)
	})

	e.DELETE("/produtos/deleter/:id", func(c echo.Context) error {
		return handler.DeleteProduto(c)
	})

	// fim produtos

	// vendas

	e.POST("/vendas/criar", func(c echo.Context) error {
		return handler.PostVenda(c)
	})

	e.GET("/vendas/listar", func(c echo.Context) error {
		return handler.GetAllVenda(c)
	})

	e.GET("/vendas/listar/:usuarioID", func(c echo.Context) error {
		return handler.GetVendasByUsuarioID(c)
	})

	e.DELETE("/vendas/deletar/:id", func(c echo.Context) error {
		return handler.DeleteVenda(c)
	})

	// fim vendas

}
