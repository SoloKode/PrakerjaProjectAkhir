package main

import (
	"projectakhir/configs"
	"projectakhir/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.InitDB()
	e := echo.New()
	routes.InitRoute(e)
	e.Start(":8000")
}
