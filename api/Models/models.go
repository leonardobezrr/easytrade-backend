package models

import "time"

// esse diretório irá armazenas as structs

type Usuarios struct {
	ID    string `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

type Produtos struct {
	ID string `json:"id"`
	Nome  string `json:"nome"`
	Descricao string `json:"descricao"`
	Preco float64 `json:"preco"`
	Qtd_estoque int64 `json:"qtd_estoque"`
	Usuarios
}

type Produto_Venda struct {
	ID string `json:"id"`
	IDProduto string `json:"idproduto"`
	IDVenda string `json:"idvenda"`
	Valor_unitario float64 `json:"valor_unitario"`
	Quantidade int64 `json:"quantidade"`
}

type Venda struct {
	ID string `json:"id"`
	Data_venda time.Time `json:"date"`
	Valor_venda float64 `json:"valor_venda"`
	Produto []Produtos 
	Usuarios
}

