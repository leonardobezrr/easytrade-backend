package models

// esse diretório irá armazenas as structs

type Usuarios struct {
	ID    string `json:"string"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

type Produtos struct {
	ID string `json:"string"`
	Nome  string `json:"nome"`
	Descricao string `json:"descricao"`
	Preco float64 `json:"preco"`
	Qtd_estoque int64 `json:"qtd_estoque"`
}