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
