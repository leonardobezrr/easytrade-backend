package main

import (
	"easytrady-backend/api"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env", err)
	}

	e := echo.New()

	// Configurar o middleware CORS com as opções desejadas
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://127.0.0.1:5500"},                                            // Adicione suas origens permitidas
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete}, // Adicione métodos permitidos
		AllowHeaders: []string{echo.HeaderContentType},                                             // Adicione cabeçalhos permitidos
	}))

	api.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))
}
