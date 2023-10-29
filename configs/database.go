package configs

import (
	"projectakhir/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB menginisialisasi koneksi database dan melakukan migrasi model-model yang diperlukan.
func InitDB() {
	// Konfigurasi koneksi database MySQL
	dsn := "root:@tcp(127.0.0.1:3306)/perpustakaan?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	// Membuka koneksi database menggunakan GORM
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed connect database")
	}

	// Melakukan migrasi tabel-tabel ke dalam database (membuat tabel jika belum ada)
	Migration()
}

// Migration melakukan otomatisasi migrasi model-model ke dalam database.
func Migration() {
	// Migrasi model "Book" ke dalam database
	DB.AutoMigrate(&models.Book{})

	// Migrasi model "Member" ke dalam database
	DB.AutoMigrate(&models.Member{})

	// Migrasi model "Sewa" ke dalam database
	DB.AutoMigrate(&models.Sewa{})
}
