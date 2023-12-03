package repository

import (
	db "easytrady-backend/api/DB"
	models "easytrady-backend/api/Models"
)

// esse repostitório irá lidar com as consultas ao banco de dados
func InsertUser(usuario models.Usuarios) (id int, err error) {

	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO usuarios (nome, email, senha) VALUES ($1, $2, $3) RETURNING id`

	err = conn.QueryRow(sql, usuario.Nome, usuario.Email, usuario.Senha).Scan(&id)

	return
}
