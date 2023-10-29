package repositories

import (
	"projectakhir/configs"
	"projectakhir/models"
)

// AddBook digunakan untuk menambahkan data buku ke dalam basis data.
func AddBook(book *models.Book) error {
	result := configs.DB.Create(book) // Menyimpan data buku ke basis data
	return result.Error               // Mengembalikan kesalahan (jika ada) selama penyimpanan
}

// GetBooks digunakan untuk mengambil semua data buku dari basis data.
func GetBooks(books *[]models.Book) error {
	result := configs.DB.Find(books) // Mengambil semua data buku dari basis data
	return result.Error              // Mengembalikan kesalahan (jika ada) selama pengambilan data
}

// GetDetailBook digunakan untuk mengambil detail buku berdasarkan ID buku.
func GetDetailBook(book *[]models.Book, id string) error {
	result := configs.DB.First(book, id) // Mengambil detail buku berdasarkan ID buku
	return result.Error                  // Mengembalikan kesalahan (jika ada) selama pengambilan data
}
