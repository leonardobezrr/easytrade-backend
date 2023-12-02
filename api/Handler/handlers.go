package handler

import (
	models "easytrady-backend/api/Models"
	"net/http"

	"time"

	"github.com/labstack/echo/v4"
)

// esse diretório irá lidar com as requisições

func GetUsuario(c echo.Context) error {
	user := models.Usuarios{"dsdaa", "Ricardo", "ric@ric.com", "1234567"}
	user2 := models.Usuarios{"1", "Leo", "leo@leo.com", "123456"}

	users := []models.Usuarios{user, user2}

	return c.JSON(http.StatusOK, users)
}

func GetProduto(c echo.Context) error {
	produto := models.Produtos{"dfdsf", "Mouse", "Mouse Gamer", 10.5, 10}
	return c.JSON(http.StatusOK, produto)
}

func GetVenda(c echo.Context) error {
	user := models.Usuarios{"dsdaa", "Ricardo", "ric@ric.com", "1234567"}
	produto := models.Produtos{"dfdsf", "Mouse", "Mouse Gamer", 10.5, 10}
	produtosSlice := []models.Produtos{produto}
	venda := models.Venda{"dafdfds", time.Now(), 10.5, produtosSlice, user}
	return c.JSON(http.StatusOK, venda)
}

func GetProdutoVenda(c echo.Context) error {
	produto := models.Produtos{"dfdsf", "Mouse", "Mouse Gamer", 10.5, 10}
	venda := models.Venda{"dafdfds", time.Now(), 10.5, []models.Produtos{produto}, models.Usuarios{"dsdaa", "Ricardo", "ric@ric.com", "1234567"}}
	produtoVenda := models.Produto_Venda{produto.ID, venda.ID, 10.5, 1}
	return c.JSON(http.StatusOK, produtoVenda)
}
