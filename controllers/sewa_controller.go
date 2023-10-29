package controllers

import (
	"net/http"
	"projectakhir/models"
	"projectakhir/repositories"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// SewaBukuController adalah handler untuk permintaan peminjaman buku.
func SewaBukuController(c echo.Context) error {
	var bukupilihan []models.Book
	id := c.Param("idbuku")

	// Memanggil metode untuk mendapatkan detail buku dari repositori atau database
	err := repositories.GetDetailBook(&bukupilihan, id)

	if err != nil {
		// Mengatasi kesalahan jika terjadi kesalahan dalam mengambil data buku
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Gagal mengambil data buku dari database",
			Data:    nil,
		})
	}

	var SecretKey = []byte("tokobuku")
	tokenHeader := c.Request().Header.Get("Authorization")
	if tokenHeader == "" {
		// Mengatasi akses yang tidak sah jika token tidak ditemukan
		return c.JSON(http.StatusUnauthorized, models.BaseResponse{
			Status:  false,
			Message: "Unauthorized1",
			Data:    nil,
		})
	}
	token := strings.Replace(tokenHeader, "bearer ", "", 1)
	tokenClaims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		// Mengatasi akses yang tidak sah jika terjadi kesalahan dalam parsing token
		return c.JSON(http.StatusUnauthorized, models.BaseResponse{
			Status:  false,
			Message: "Unauthorized2",
			Data:    nil,
		})
	}
	claims, ok := tokenClaims.Claims.(jwt.MapClaims)
	if !ok || !tokenClaims.Valid {
		// Mengatasi akses yang tidak sah jika token tidak valid
		return c.JSON(http.StatusUnauthorized, models.BaseResponse{
			Status:  false,
			Message: "Unauthorized3",
			Data:    nil,
		})
	}
	userIdAnggota := int(claims["userId"].(float64)) // Ubah "userId" dari token ke int
	existingMember, err := repositories.GetMemberByIdAnggota(userIdAnggota)
	if err != nil {
		// Mengatasi jika anggota tidak ditemukan dalam database
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Status:  false,
			Message: "Member not found",
			Data:    nil,
		})
	}
	sewaBuku := new(models.Sewa)
	sewaBuku.IdAnggota = existingMember.IdAnggota
	sewaBuku.NamaAnggota = existingMember.Name
	sewaBuku.IdBuku = bukupilihan[0].IdBuku
	sewaBuku.JudulBuku = bukupilihan[0].Judul
	err = repositories.SewaBuku(sewaBuku)
	if err != nil {
		// Mengatasi kesalahan jika gagal menyimpan data peminjaman buku ke database
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed create in database",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Successfully created",
		Data:    sewaBuku,
	})
}

// GetSewaController adalah handler untuk permintaan data peminjaman buku.
func GetSewaController(c echo.Context) error {
	var sewa []models.Sewa

	err := repositories.GetSewa(&sewa)

	if err != nil {
		// Mengatasi kesalahan jika terjadi kesalahan dalam mengambil data peminjaman buku dari database
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed get data from database",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Successfully get data",
		Data:    sewa,
	})
}
