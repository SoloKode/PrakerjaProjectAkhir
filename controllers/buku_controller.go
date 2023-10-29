package controllers

import (
	"net/http"
	"projectakhir/models"
	"projectakhir/repositories"

	"github.com/labstack/echo/v4"
)

// AddBookController adalah handler untuk menambahkan buku ke database.
func AddBookController(c echo.Context) error {
	var book models.Book
	c.Bind(&book)

	// Memanggil repositori untuk menambahkan buku ke database
	err := repositories.AddBook(&book)

	// Mengatasi kesalahan jika terjadi kesalahan saat menambahkan buku
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed create in database",
			Data:    nil,
		})
	}

	// Mengembalikan respons berhasil dengan status 200 jika buku berhasil ditambahkan
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Successfully created",
		Data:    book,
	})
}

// GetBookController adalah handler untuk mendapatkan daftar buku dari database.
func GetBookController(c echo.Context) error {
	var book []models.Book

	// Memanggil repositori untuk mendapatkan daftar buku dari database
	err := repositories.GetBooks(&book)

	// Mengatasi kesalahan jika terjadi kesalahan saat mengambil data buku
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed get data from database",
			Data:    nil,
		})
	}

	// Mengembalikan daftar buku dengan status 200 jika berhasil
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Successfully get data",
		Data:    book,
	})
}

// GetDetailBookController adalah handler untuk mendapatkan detail buku berdasarkan ID dari database.
func GetDetailBookController(c echo.Context) error {
	var book []models.Book
	id := c.Param("id")

	// Memanggil metode untuk mendapatkan detail buku dari repositori atau database
	err := repositories.GetDetailBook(&book, id)

	// Mengatasi kesalahan jika terjadi kesalahan dalam mengambil data buku
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Gagal mengambil data buku dari database",
			Data:    nil,
		})
	}

	// Mengembalikan data buku dengan status 200 jika berhasil
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Berhasil mendapatkan detail buku",
		Data:    book,
	})
}
