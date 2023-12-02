package main

import (
	"easytrady-backend/api"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	api.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}
