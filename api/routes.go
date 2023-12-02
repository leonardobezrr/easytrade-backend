package api

import (
	handler "easytrady-backend/api/Handler"
	"net/http"

	"github.com/labstack/echo/v4"
)

// criar todas as rotas aqui
func SetupRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/usuarios", handler.GetUsuario)
	e.GET("/produto", handler.GetProduto)
	e.GET("/venda", handler.GetVenda)
	e.GET("/produto_venda", handler.GetProdutoVenda)

}
