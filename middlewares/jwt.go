package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// jwtCustomClaims adalah struktur data yang digunakan untuk menentukan klaim khusus dalam token JWT.
type jwtCustomClaims struct {
	UserId int    `json:"userId"` // ID pengguna (anggota atau admin)
	Name   string `json:"name"`   // Nama pengguna
	jwt.RegisteredClaims
}

// GenerateMemberToken digunakan untuk menghasilkan token JWT untuk anggota (user).
func GenerateMemberToken(userId int, name string) string {
	claims := &jwtCustomClaims{
		userId,
		name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)), // Token berlaku selama 72 jam
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte("tokobuku")) // Menandatangani token dengan kunci "tokobuku"
	return t
}

// GenerateAdminToken digunakan untuk menghasilkan token JWT untuk admin.
func GenerateAdminToken(userId int, name string) string {
	claims := &jwtCustomClaims{
		userId,
		name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)), // Token berlaku selama 72 jam
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte("admintokobuku")) // Menandatangani token dengan kunci "admintokobuku"
	return t
}
