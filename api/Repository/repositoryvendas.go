package repository

import (
	db "easytrady-backend/api/DB"
	models "easytrady-backend/api/Models"
	"fmt"
	"log"
)

func InsertVenda(venda models.Venda) (id int, err error) {

	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO VENDAS (data_venda, valor_venda, id_usuario) VALUES ($1, $2, $3) RETURNING id`
	err = conn.QueryRow(sql, venda.Data_venda, venda.Valor_venda, venda.Usuarios).Scan(&id)

	if err != nil {
		fmt.Println("Erro ao inserir venda no banco de dados: ", err)
		return
	}

	for _, produto := range venda.Produtos {
		sqlProdutoVenda := `INSERT INTO produtos_venda (id_produto, id_venda, valor_unitario, quantidade) VALUES ($1, $2, $3, $4)`
		_, err = conn.Exec(sqlProdutoVenda, produto.IDProduto, id, produto.Valor_unitario, produto.Quantidade)

		if err != nil {
			fmt.Println("Erro ao inserir produto_venda no banco de dados: ", err)
			return
		}
	}

	return id, nil
}

func GetVenda() ([]models.Venda, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM vendas;")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var vendas []models.Venda
	for rows.Next() {
		var venda models.Venda
		if err := rows.Scan(&venda.ID, &venda.Data_venda, &venda.Valor_venda, &venda.Usuarios); err != nil {
			log.Fatal(err)
			return nil, err
		}
		vendas = append(vendas, venda)
	}

	return vendas, nil
}
