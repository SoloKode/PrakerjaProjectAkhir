package repositories

import (
	"projectakhir/configs"
	"projectakhir/models"
)

// RegisterMember digunakan untuk mendaftarkan anggota baru dalam basis data.
func RegisterMember(member *models.Member) error {
	result := configs.DB.Create(member) // Menyimpan data anggota ke basis data
	return result.Error                 // Mengembalikan kesalahan (jika ada) selama penyimpanan
}

// GetMembers digunakan untuk mengambil semua data anggota dari basis data.
func GetMembers(member *[]models.Member) error {
	result := configs.DB.Find(member) // Mengambil semua data anggota dari basis data
	return result.Error               // Mengembalikan kesalahan (jika ada) selama pengambilan data
}

// GetDetailMember digunakan untuk mengambil detail anggota berdasarkan ID anggota.
func GetDetailMember(member *[]models.Member, id string) error {
	result := configs.DB.First(member, id) // Mengambil detail anggota berdasarkan ID anggota
	return result.Error                    // Mengembalikan kesalahan (jika ada) selama pengambilan data
}

// UpdateMember digunakan untuk memperbarui data anggota berdasarkan ID anggota.
func UpdateMember(member *models.Member, id string) error {
	result := configs.DB.Model(&models.Member{}).Where("id_anggota = ?", id).Updates(member) // Memperbarui data anggota
	if result.Error != nil {
		return result.Error
	}
	return result.Error
}

// GetMemberByIdAnggota digunakan untuk mengambil data anggota berdasarkan ID anggota.
func GetMemberByIdAnggota(userIdAnggota int) (*models.Member, error) {
	member := new(models.Member)
	result := configs.DB.First(member, "id_anggota = ?", userIdAnggota) // Mengambil anggota berdasarkan ID anggota
	if result.Error != nil {
		return nil, result.Error
	}
	return member, nil
}

// UpdateMemberProfile digunakan untuk memperbarui profil anggota dalam basis data.
func UpdateMemberProfile(member *models.Member) error {
	result := configs.DB.Save(member) // Memperbarui profil anggota dalam basis data
	if result.Error != nil {
		return result.Error
	}
	return nil
}
