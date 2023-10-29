package models

import "time"

// Sewa adalah struktur data yang merepresentasikan penyewaan buku.
type Sewa struct {
	IdSewa      int       `json:"idSewa" gorm:"primaryKey"`                         // ID unik untuk penyewaan
	IdBuku      int       `json:"idBuku" gorm:"foreignKey:id_buku; not null"`       // ID buku yang disewa (kunci asing)
	IdAnggota   int       `json:"idAnggota" gorm:"foreignKey:id_anggota; not null"` // ID anggota yang menyewa (kunci asing)
	JudulBuku   string    `json:"judul" gorm:"type:varchar(255); not null"`         // Judul buku yang disewa
	NamaAnggota string    `json:"namaAnggota" gorm:"type:varchar(255); not null"`   // Nama anggota yang menyewa
	CreatedAt   time.Time `json:"createdAt"`                                        // Waktu pembuatan penyewaan
	UpdatedAt   time.Time `json:"updatedAt"`                                        // Waktu pembaruan penyewaan
}
