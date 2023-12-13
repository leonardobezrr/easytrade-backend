package handler

import (
	models "easytrady-backend/api/Models"
	repository "easytrady-backend/api/Repository"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func PostVenda(c echo.Context) error {

	venda := models.Venda{}
	err := c.Bind(&venda)

	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Solicitação inválida")
	}

	id, err := repository.InsertVenda(venda)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro ao inserir venda no banco de dados")
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": fmt.Sprintf("Venda inserida com sucesso ID: %d", id),
	})
}

func GetAllVenda(c echo.Context) error {
	vendas, err := repository.GetVenda()
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao obter vendas do banco de dados"})
	}
	return c.JSON(http.StatusOK, vendas)
}

func GetVendasByUsuarioID(c echo.Context) error {
	usuarioID := c.Param("usuarioID")

	vendas, err := repository.GetVendasByUsuarioID(usuarioID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao obter vendas do usuário"})
	}

	return c.JSON(http.StatusOK, vendas)
}

func GetProdutoByVendaId(c echo.Context) error {
	vendaID := c.Param("vendaID")

	produtos, err := repository.GetProdutoByVendaId(vendaID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao obter produtos das vendas"})
	}

	return c.JSON(http.StatusOK, produtos)
}

func DeleteVenda(c echo.Context) error {
	var venda models.Venda
	err := c.Bind(&venda)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid request payload")
	}

	id := c.Param("id")

	venda.ID = id

	err = repository.DeleteVenda(venda)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro ao deletar venda no banco de dados")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("Venda deletada com sucesso ID: %s", venda.ID),
	})
}
