package repository

import (
	db "easytrady-backend/api/DB"
	models "easytrady-backend/api/Models"
	"fmt"
	"log"
)

func InsertUsuario(usuario models.Usuarios) (id int, err error) {

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

func UpdateUsuario(usuario models.Usuarios)error{
	conn,err := db.OpenConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	sql := `UPDATE usuarios SET nome=$1, email=$2, senha=$3 WHERE id=$4`

	_, err = conn.Exec(sql, usuario.Nome, usuario.Email, usuario.Senha, usuario.ID)
	if err != nil {
		fmt.Println("Erro ao atualizar usuário no banco de dados:",err)
		return err
	}
	return nil
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

func UpdateProduto(produto models.Produtos)error{
	conn,err := db.OpenConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	sql := `UPDATE produtos SET nome=$1, descricao=$2, preco=$3, qtd_estoque=$4 WHERE id=$5`

	_, err = conn.Exec(sql, produto.Nome, produto.Descricao, produto.Preco, produto.Qtd_estoque,produto.ID)
	if err != nil {
		fmt.Println("Erro ao atualizar produto no banco de dados:",err)
		return err
	}
	return nil
}

func DeleteProduto(produto models.Produtos)error{
	conn,err := db.OpenConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	sql := `DELETE FROM produtos WHERE id=$1`

	_, err = conn.Exec(sql, produto.ID)
	if err != nil {
		fmt.Println("Erro ao deletar produto no banco de dados:",err)
		return err
	}
	return nil
}

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
