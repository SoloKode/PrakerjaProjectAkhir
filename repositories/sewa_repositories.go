package repositories

import (
	"projectakhir/configs"
	"projectakhir/models"
)

// SewaBuku digunakan untuk menyimpan data penyewaan buku ke basis data.
func SewaBuku(sewa *models.Sewa) error {
	result := configs.DB.Create(sewa) // Menyimpan data sewa ke basis data
	return result.Error               // Mengembalikan kesalahan (jika ada) selama penyimpanan
}

// GetSewa digunakan untuk mengambil semua data penyewaan buku dari basis data.
func GetSewa(sewa *[]models.Sewa) error {
	result := configs.DB.Find(sewa) // Mengambil semua data penyewaan buku dari basis data
	return result.Error             // Mengembalikan kesalahan (jika ada) selama pengambilan data
}
