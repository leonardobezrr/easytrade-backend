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
	err = conn.QueryRow(sql, venda.Data_venda, venda.Valor_venda, venda.Usuarios, venda.Produtos).Scan(&id)

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

func GetProdutoByVendaId(vendaID string) ([]models.Produto_Venda, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// Construa a consulta SQL com base no parâmetro vendaID
	query := `SELECT produtos_venda.id, produtos_venda.id_produto, vendas.id, produtos_venda.valor_unitario, produtos_venda.quantidade FROM vendas JOIN produtos_venda ON vendas.id = produtos_venda.id_venda WHERE vendas.id = $1;`

	rows, err := conn.Query(query, vendaID)
	if err != nil {
		log.Println("Erro ao executar a consulta:", err)
		return nil, err
	}
	defer rows.Close()

	// Crie um slice para armazenar os resultados
	var produtos []models.Produto_Venda

	// Itere sobre as linhas do resultado
	for rows.Next() {
		var produto models.Produto_Venda
		err := rows.Scan(
			&produto.ID,
			&produto.IDProduto,
			&produto.IDVenda,
			&produto.Valor_unitario,
			&produto.Quantidade,
		)
		if err != nil {
			log.Println("Erro ao escanear linha:", err)
			return nil, err
		}
		produtos = append(produtos, produto)
	}

	// Verifique se houve algum erro durante a iteração
	if err := rows.Err(); err != nil {
		log.Println("Erro durante a iteração das linhas:", err)
		return nil, err
	}

	// Retorne os resultados
	return produtos, nil
}

func GetVendasByUsuarioID(usuarioID string) ([]models.Venda, error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	query := "SELECT * FROM vendas WHERE id_usuario = $1"
	rows, err := conn.Query(query, usuarioID)
	if err != nil {
		log.Fatal(err)
		return nil, err
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

func GetVenda() ([]models.Venda, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM vendas")
	if err != nil {
		log.Fatal(err)
		return nil, err
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

func DeleteVenda(venda models.Venda) error {
	conn, err := db.OpenConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	sqlProdutoVenda := `DELETE FROM produtos_venda WHERE id_venda=$1`

	_, err = conn.Exec(sqlProdutoVenda, venda.ID)
	if err != nil {
		fmt.Println("Erro ao deletar venda na tabela produtos_venda no banco de dados:", err)
		return err
	}

	sql := `DELETE FROM vendas WHERE id=$1`

	_, err = conn.Exec(sql, venda.ID)
	if err != nil {
		fmt.Println("Erro ao deletar venda no banco de dados:", err)
		return err
	}

	return nil
}
