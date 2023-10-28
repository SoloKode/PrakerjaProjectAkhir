package controllers

import (
	"projectakhir/models"
	"projectakhir/repositories"

	"github.com/labstack/echo/v4"
)

func AddUsersController(c echo.Context) error {
	var user models.Buku
	c.Bind(&user)

	err := repositories.AddUser(&user)

	if err != nil {
		return c.JSON(500, models.BaseResponse{
			Status:  false,
			Message: "Failed create in database",
			Data:    nil,
		})
	}

	return c.JSON(200, models.BaseResponse{
		Status:  true,
		Message: "Successfully created",
		Data:    user,
	})
}

func GetUsersController(c echo.Context) error {
	var users []models.Buku

	err := repositories.GetUser(&users)

	if err != nil {
		return c.JSON(500, models.BaseResponse{
			Status:  false,
			Message: "Failed get data from database",
			Data:    nil,
		})
	}
	return c.JSON(200, models.BaseResponse{
		Status:  true,
		Message: "Successfully get data",
		Data:    users,
	})
}
