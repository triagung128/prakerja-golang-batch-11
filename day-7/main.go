package main

import (
	"day-7/configs"
	"day-7/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.InitDB()
	e := echo.New()
	routes.InitRoutes(e)
	e.Start(":8000")
}
