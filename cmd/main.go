package main

import (
	"easytrady-backend/api"
	"log"

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

	e.Use(middleware.CORS())

	api.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}
