package models

// BaseResponse adalah struktur data yang digunakan untuk merespons permintaan HTTP umum dalam aplikasi.
type BaseResponse struct {
	Status  bool        `json:"status"`  // Status respons (biasanya true untuk sukses, false untuk kesalahan)
	Message string      `json:"message"` // Pesan respons yang memberikan informasi mengenai permintaan
	Data    interface{} `json:"data"`    // Data respons yang berisi payload atau hasil dari permintaan
}
