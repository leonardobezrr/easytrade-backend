package repository

import (
	db "easytrady-backend/api/DB"
	models "easytrady-backend/api/Models"
	"fmt"
)

func InsertUser(usuario models.Usuarios) (id int, err error) {

	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO usuarios (nome, email, senha) VALUES ($1, $2, $3) RETURNING id`

	err = conn.QueryRow(sql, usuario.Nome, usuario.Email, usuario.Senha).Scan(&id)

	if err != nil {
		fmt.Println("Erro ao inserir produto no banco de dados:", err)
	}

	return
}

func InsertProduto(produto models.Produtos) (id int, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO produtos (nome, descricao, preco, qtd_estoque, id_usuario) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err = conn.QueryRow(sql, produto.Nome, produto.Descricao, produto.Preco, produto.Qtd_estoque, produto.Usuarios).Scan(&id)

	if err != nil {
		fmt.Println("Erro ao inserir produto no banco de dados:", err)
	}

	return
}
