package models

import (
	"time"

	"gorm.io/gorm"
)

// Book adalah struktur data yang merepresentasikan buku dalam aplikasi.
type Book struct {
	IdBuku    int    `json:"idbuku" gorm:"primaryKey"`                     // ID unik untuk buku
	Judul     string `json:"judul" gorm:"type:varchar(255); not null"`     // Judul buku
	Deskripsi string `json:"deskripsi" gorm:"not null"`                    // Deskripsi buku
	Pengarang string `json:"pengarang" gorm:"type:varchar(255); not null"` // Nama pengarang buku
	Penerbit  string `json:"penerbit" gorm:"type:varchar(255); not null"`  // Nama penerbit buku
	Tahun     int    `json:"tahun" gorm:"type:int(4)not null"`             // Tahun terbit buku

	CreatedAt time.Time      // Waktu pembuatan buku
	UpdatedAt time.Time      // Waktu pembaruan buku
	DeletedAt gorm.DeletedAt `gorm:"index"` // Informasi mengenai penghapusan buku (jika ada)
}
