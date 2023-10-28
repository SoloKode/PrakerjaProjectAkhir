package models

import (
	"time"

	"gorm.io/gorm"
)

type Buku struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Judul     string `json:"judul" gorm:"not null"`
	Deskripsi string `json:"deskripsi" gorm:"not null"`
	Pengarang string `json:"pengarang" gorm:"not null"`
	Penerbit  string `json:"penerbit" gorm:"not null"`
	Tahun     int    `json:"tahun" gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
