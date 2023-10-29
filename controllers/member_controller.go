package controllers

import (
	"net/http"
	"projectakhir/models"
	"projectakhir/repositories"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// GetMembersController adalah handler untuk permintaan data anggota.
func GetMembersController(c echo.Context) error {
	// Menyiapkan variabel untuk menyimpan data anggota
	var member []models.Member

	// Memanggil metode repositori untuk mengambil data anggota
	err := repositories.GetMembers(&member)

	// Memeriksa apakah terdapat kesalahan dalam mengambil data anggota
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed get data from database",
			Data:    nil,
		})
	}

	// Mengembalikan respons dengan status 200 dan data anggota jika berhasil
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Successfully get data",
		Data:    member,
	})
}

// GetDetailMemberController adalah handler untuk mendapatkan detail pengguna berdasarkan ID.
func GetDetailMemberController(c echo.Context) error {
	// Menyiapkan variabel untuk menyimpan data pengguna
	var member []models.Member
	id := c.Param("id")

	// Memanggil metode repositori untuk mendapatkan detail pengguna dari database
	err := repositories.GetDetailMember(&member, id)

	// Mengatasi kesalahan jika terjadi kesalahan dalam mengambil data pengguna
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Gagal mengambil data pengguna dari database",
			Data:    nil,
		})
	}

	// Mengembalikan data pengguna dengan status 200 jika berhasil
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Berhasil mendapatkan detail pengguna",
		Data:    member,
	})
}

// UpdateMemberController adalah handler untuk memperbarui data anggota berdasarkan ID.
func UpdateMemberController(c echo.Context) error {
	// Mendapatkan IdAnggota anggota yang akan diperbarui dari parameter URL
	id := c.Param("id")

	// Membaca data yang akan diperbarui dari permintaan HTTP
	updatedMember := new(models.Member)
	if err := c.Bind(updatedMember); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Status:  false,
			Message: "Invalid request body",
			Data:    nil,
		})
	}

	// Memanggil repositori untuk memperbarui anggota
	err := repositories.UpdateMember(updatedMember, id)

	// Mengatasi kesalahan jika terjadi kesalahan dalam proses pembaruan anggota
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed to update member",
			Data:    nil,
		})
	}

	// Mengembalikan respons dengan status 200 jika pembaruan berhasil
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Member updated successfully",
		Data:    updatedMember,
	})
}

// UpdateProfileController adalah handler untuk memperbarui profil anggota.
func UpdateProfileController(c echo.Context) error {
	// Mendapatkan token JWT dari permintaan HTTP
	var SecretKey = []byte("tokobuku")
	tokenHeader := c.Request().Header.Get("Authorization")
	if tokenHeader == "" {
		return c.JSON(http.StatusUnauthorized, models.BaseResponse{
			Status:  false,
			Message: "Unauthorized1",
			Data:    nil,
		})
	}

	token := strings.Replace(tokenHeader, "bearer ", "", 1)

	// Melakukan decoding token
	tokenClaims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Disini Anda dapat memeriksa validitas token, seperti memastikan token sesuai dengan algoritma yang digunakan
		return SecretKey, nil
	})
	// Memeriksa apakah token valid
	if err != nil {
		return c.JSON(http.StatusUnauthorized, models.BaseResponse{
			Status:  false,
			Message: "Unauthorized2",
			Data:    nil,
		})
	}

	// Mengekstrak klaim-klaim dari token yang telah di-decode
	claims, ok := tokenClaims.Claims.(jwt.MapClaims)
	if !ok || !tokenClaims.Valid {
		return c.JSON(http.StatusUnauthorized, models.BaseResponse{
			Status:  false,
			Message: "Unauthorized3",
			Data:    nil,
		})
	}

	// Mendapatkan userIdAnggota dari klaim token
	userIdAnggota := int(claims["userId"].(float64)) // Ubah "userId" ke int

	// Mendapatkan data anggota dari database berdasarkan IdAnggota
	existingMember, err := repositories.GetMemberByIdAnggota(userIdAnggota)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Status:  false,
			Message: "Member not found",
			Data:    nil,
		})
	}

	// Mendapatkan data yang akan diperbarui dari permintaan HTTP
	updatedProfile := new(models.Member)
	if err := c.Bind(updatedProfile); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Status:  false,
			Message: "Invalid request body",
			Data:    nil,
		})
	}

	// Perbarui field-field profil anggota dengan data yang baru
	existingMember.Name = updatedProfile.Name
	existingMember.Email = updatedProfile.Email
	existingMember.Phone = updatedProfile.Phone
	existingMember.Address = updatedProfile.Address

	// Memanggil repositori untuk menyimpan perubahan profil anggota
	err = repositories.UpdateMemberProfile(existingMember)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed to update profile",
			Data:    nil,
		})
	}

	// Mengembalikan respons dengan status 200 jika pembaruan profil berhasil
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Profile updated successfully",
		Data:    existingMember,
	})
}

// GetProfileController adalah handler untuk mendapatkan profil anggota berdasarkan token JWT.
func GetProfileController(c echo.Context) error {
	// Mendapatkan token JWT dari permintaan HTTP
	var SecretKey = []byte("tokobuku")
	tokenHeader := c.Request().Header.Get("Authorization")

	// Memeriksa apakah token kosong, jika iya, kembalikan respons 401 (Unauthorized)
	if tokenHeader == "" {
		return c.JSON(http.StatusUnauthorized, models.BaseResponse{
			Status:  false,
			Message: "Unauthorized1",
			Data:    nil,
		})
	}

	// Menghapus kata "bearer" dari token untuk mendapatkan token asli
	token := strings.Replace(tokenHeader, "bearer ", "", 1)

	// Melakukan decoding token
	tokenClaims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Disini Anda dapat memeriksa validitas token, seperti memastikan token sesuai dengan algoritma yang digunakan
		return SecretKey, nil
	})

	// Memeriksa apakah token valid, jika tidak valid, kembalikan respons 401 (Unauthorized)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, models.BaseResponse{
			Status:  false,
			Message: "Unauthorized2",
			Data:    nil,
		})
	}

	// Mengekstrak klaim-klaim dari token yang telah di-decode
	claims, ok := tokenClaims.Claims.(jwt.MapClaims)

	// Memeriksa apakah klaim valid, jika tidak valid, kembalikan respons 401 (Unauthorized)
	if !ok || !tokenClaims.Valid {
		return c.JSON(http.StatusUnauthorized, models.BaseResponse{
			Status:  false,
			Message: "Unauthorized3",
			Data:    nil,
		})
	}

	// Mendapatkan userIdAnggota dari klaim token dan mengonversinya menjadi tipe data int
	userIdAnggota := int(claims["userId"].(float64)) // Ubah "userId" ke int

	// Mendapatkan data anggota dari database berdasarkan IdAnggota
	existingMember, err := repositories.GetMemberByIdAnggota(userIdAnggota)

	// Mengatasi kesalahan jika terjadi kesalahan dalam mengambil data anggota dari database
	if err != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Status:  false,
			Message: "Member not found",
			Data:    nil,
		})
	}

	// Mengembalikan profil anggota dengan status 200 jika berhasil
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Profile updated successfully",
		Data:    existingMember,
	})
}
