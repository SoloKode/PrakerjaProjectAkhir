package routes

import (
	"projectakhir/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Pre(middleware.RemoveTrailingSlash())
	e.GET("/buku", controllers.GetUsersController)
	e.POST("/buku", controllers.AddUsersController)
}
