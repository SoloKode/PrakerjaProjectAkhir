package controllers

import (
	"net/http"
	"projectakhir/middlewares"
	"projectakhir/models"
	"projectakhir/repositories"

	"github.com/labstack/echo/v4"
)

// RegisterController adalah handler untuk mendaftarkan anggota baru ke dalam database.
func RegisterController(c echo.Context) error {
	var member models.Member
	c.Bind(&member)

	// Memanggil repositori untuk mendaftarkan anggota baru ke dalam database
	err := repositories.RegisterMember(&member)

	// Mengatasi kesalahan jika terjadi kesalahan saat mendaftarkan anggota
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed create in database",
			Data:    nil,
		})
	}

	// Mengembalikan respons berhasil dengan status 200 jika pendaftaran berhasil
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Successfully created",
		Data:    member,
	})
}

// LoginController adalah handler untuk proses login anggota.
func LoginController(c echo.Context) error {
	var member models.Member
	c.Bind(&member)

	// Memanggil repositori untuk melakukan proses login
	err := repositories.Login(&member)

	// Mengatasi kesalahan jika username atau password tidak cocok
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Status:  false,
			Message: "username or password is wrong",
			Data:    nil,
		})
	}

	var memberResponse models.MemberResponse
	memberResponse.IdAnggota = member.IdAnggota
	memberResponse.Name = member.Name
	memberResponse.Email = member.Email

	// Menghasilkan token berdasarkan tipe anggota (Admin atau Member)
	if member.MembershipType == "Admin" {
		memberResponse.Token = middlewares.GenerateAdminToken(member.IdAnggota, member.Name)
	} else {
		memberResponse.Token = middlewares.GenerateMemberToken(member.IdAnggota, member.Name)
	}

	memberResponse.CreatedAt = member.CreatedAt
	memberResponse.UpdatedAt = member.UpdatedAt

	// Mengembalikan respons berhasil dengan status 200 jika proses login berhasil
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Successfully login",
		Data:    memberResponse,
	})
}
