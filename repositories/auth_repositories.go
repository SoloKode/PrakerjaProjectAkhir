package repositories

import (
	"projectakhir/configs"
	"projectakhir/models"
)

// Login digunakan untuk melakukan proses otentikasi pengguna berdasarkan email dan kata sandi.
func Login(member *models.Member) error {
	result := configs.DB.Where("email = ? AND password = ?", member.Email, member.Password).First(member)
	return result.Error // Mengembalikan kesalahan (jika ada) selama proses otentikasi
}
