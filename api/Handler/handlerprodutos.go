package handler

import (
	models "easytrady-backend/api/Models"
	repository "easytrady-backend/api/Repository"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllProdutos(c echo.Context) error {
	produtos, err := repository.GetProdutos()
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao obter produtos do banco de dados"})
	}
	return c.JSON(http.StatusOK, produtos)
}

func GetProdutosByUsuarioID(c echo.Context) error {
	usuarioID := c.Param("usuarioID")

	produtos, err := repository.GetProdutosByUsuarioID(usuarioID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao obter produtos do usuário"})
	}

	return c.JSON(http.StatusOK, produtos)
}

func PostProduto(c echo.Context) error {
	produto := models.Produtos{}
	err := c.Bind(&produto)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid request payload")
	}

	id, err := repository.InsertProduto(produto)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro em inserir o produto no banco de dados")
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": fmt.Sprintf("Produto inserido com sucesso ID: %d", id),
	})
}

func UpdateProduto(c echo.Context) error {
	var produto models.Produtos
	err := c.Bind(&produto)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid request payload")
	}

	id := c.Param("id")

	produto.ID = id

	err = repository.UpdateProduto(produto)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro ao atualizar produto no banco de dados")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("Produto atualizado com sucesso ID: %s", produto.ID),
	})
}

func DeleteProduto(c echo.Context) error {
	var produto models.Produtos
	err := c.Bind(&produto)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid request payload")
	}

	id := c.Param("id")

	produto.ID = id

	err = repository.DeleteProduto(produto)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro ao deletar produto no banco de dados")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("Produto deletado com sucesso ID: %s", produto.ID),
	})
}
