package main

import (
	"easytrady-backend/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.CORS())

	api.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}
