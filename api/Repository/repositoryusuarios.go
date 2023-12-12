package repository

import (
	db "easytrady-backend/api/DB"
	models "easytrady-backend/api/Models"
	service "easytrady-backend/api/Service"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsuarios() ([]models.Usuarios, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM usuarios;")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var usuarios []models.Usuarios
	for rows.Next() {
		var usuario models.Usuarios
		if err := rows.Scan(&usuario.ID, &usuario.Nome, &usuario.Email, &usuario.Senha); err != nil {
			log.Fatal(err)
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func InsertUsuario(usuario models.Usuarios) (id int, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	err = service.CheckEmailExists(conn, usuario.Email)
	if err != nil {
		return 0, err
	}

	err = service.CheckMail(usuario.Email)
	if err != nil {
		return 0, err
	}

	if len(usuario.Senha) < 6 {
		return 0, echo.NewHTTPError(http.StatusBadRequest, "A senha deve conter pelo menos 6 caracteres.")
	}

	hashedSenha, err := service.HashSenha(usuario.Senha)
	if err != nil {
		return 0, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Erro ao criar hash da senha: %s", err))
	}

	sql := `INSERT INTO usuarios (nome, email, senha) VALUES ($1, $2, $3) RETURNING id`

	err = conn.QueryRow(sql, usuario.Nome, usuario.Email, hashedSenha).Scan(&id)
	if err != nil {
		return 0, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Erro ao inserir usuário no banco de dados: %s", err))
	}

	return
}

func UpdateUsuario(usuario models.Usuarios) error {
	conn, err := db.OpenConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	sql := `UPDATE usuarios SET nome=$1, email=$2, senha=$3 WHERE id=$4`

	_, err = conn.Exec(sql, usuario.Nome, usuario.Email, usuario.Senha, usuario.ID)
	if err != nil {
		fmt.Println("Erro ao atualizar usuário no banco de dados:", err)
		return err
	}
	return nil
}
