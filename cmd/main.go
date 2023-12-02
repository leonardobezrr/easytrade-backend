package main

import (
	models "easytrady-backend/api/Models"
	"fmt"
	"time"
)

func main() {

	// e := echo.New()
	// api.SetupRoutes(e)
	// e.Logger.Fatal(e.Start(":3000"))

	user := models.Usuarios{"dsdaa", "Ricardo", "ric@ric.com", "1234567"}
	produto := models.Produtos{"dfdsf", "Mouse", "Mouse Gamer", 10.5, 10}
	produtosSlice := []models.Produtos{produto}
	venda := models.Venda{"dafdfds", time.Now(), 10.5, produtosSlice, user}
	produto_venda := models.Produto_Venda{produto.ID, venda.ID, 10.5, 1}

	fmt.Println("Usuário: ", user, "\nProduto: ", produto, "\nVenda: ", venda, "\nProduto venda: ", produto_venda)

}
