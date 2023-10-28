package models

import (
	"time"

	"gorm.io/gorm"
)

type Buku struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Judul     string `json:"judul" gorm:"type:varchar(255); not null"`
	Deskripsi string `json:"deskripsi" gorm:"not null"`
	Pengarang string `json:"pengarang" gorm:"type:varchar(255); not null"`
	Penerbit  string `json:"penerbit" gorm:"type:varchar(255); not null"`
	Tahun     int    `json:"tahun" gorm:"type:int(4)not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
