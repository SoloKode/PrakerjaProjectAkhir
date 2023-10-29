package models

import "time"

// Member adalah struktur data yang merepresentasikan anggota dalam aplikasi.
type Member struct {
	IdAnggota      int       `json:"idAnggota" gorm:"primaryKey"`                    // ID unik untuk anggota
	Name           string    `json:"name" gorm:"type:varchar(50);not null"`          // Nama anggota
	Email          string    `json:"email" gorm:"type:varchar(100);not null;unique"` // Alamat email anggota (unik)
	Password       string    `json:"password" gorm:"type:varchar(100);not null"`     // Kata sandi anggota
	Phone          string    `json:"phone" gorm:"type:varchar(20);not null"`         // Nomor telepon anggota
	Address        string    `json:"address" gorm:"not null"`                        // Alamat anggota
	CreatedAt      time.Time // Waktu pembuatan anggota
	UpdatedAt      time.Time // Waktu pembaruan anggota
	MembershipType string    `json:"membership_type" gorm:"type:varchar(20);not null"` // Jenis keanggotaan anggota
}

// MemberResponse adalah struktur data yang digunakan untuk merespons data anggota.
type MemberResponse struct {
	IdAnggota      int       `json:"idAnggota"` // ID anggota
	Name           string    `json:"name"`      // Nama anggota
	Email          string    `json:"email"`     // Alamat email anggota
	Password       string    `json:"password"`  // Kata sandi anggota
	Token          string    `json:"token"`     // Token
	Phone          string    `json:"phone"`     // Nomor telepon anggota
	Address        string    `json:"address"`   // Alamat anggota
	CreatedAt      time.Time // Waktu pembuatan anggota
	UpdatedAt      time.Time // Waktu pembaruan anggota
	MembershipType string    `json:"membership_type"` // Jenis keanggotaan anggota
}
