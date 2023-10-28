package repositories

import (
	"fmt"
	"projectakhir/configs"
	"projectakhir/models"
)

func AddUser(user *models.Buku) error {
	result := configs.DB.Create(user)
	return result.Error
}

func GetUser(users *[]models.Buku) error {
	result := configs.DB.Find(users)
	fmt.Println(users)
	return result.Error
}
