package handler

import (
	models "easytrady-backend/api/Models"
	repository "easytrady-backend/api/Repository"
	"fmt"
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
