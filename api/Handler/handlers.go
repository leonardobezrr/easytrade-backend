package handler

import (
	models "easytrady-backend/api/Models"
	repository "easytrady-backend/api/Repository"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func PostUsuario(c echo.Context) error {
	usuario := models.Usuarios{}
	err := c.Bind(&usuario)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid request payload")
	}

	id, err := repository.InsertUser(usuario)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro ao inserir usuário no banco de dados")
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": fmt.Sprintf("Usuário inserido com sucesso ID: %d", id),
	})
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

func GetAllUsuarios(c echo.Context) error {
	usuarios, err := repository.GetUsuarios()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao obter usuários do banco de dados"})
	}
	return c.JSON(http.StatusOK, usuarios)
}

func Login(c echo.Context) error {
	var usuario models.Usuarios
	if err := c.Bind(&usuario); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao decodificar dados do usuário"})
	}

	usuarios, err := repository.GetUsuarios()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao obter usuários do banco de dados"})
	}

	autenticado := false
	for _, u := range usuarios {
		if u.Email == usuario.Email && u.Senha == usuario.Senha {
			autenticado = true
			break
		}
	}

	if autenticado {
		return c.JSON(http.StatusOK, map[string]string{"message": "Login bem-sucedido para o usuário: " + usuario.Email})
	}

	return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Credenciais inválidas"})
}
