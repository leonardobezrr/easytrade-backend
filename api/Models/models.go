package models

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
}

type Produto_Venda struct {
	IDProduto string `json:"idproduto"`
	IDVenda string `json:"idvenda"`
	Valor_unitario float64 `json:"valor_unitario"`
	Quantidade int64 `json:"quantidade"`
}
