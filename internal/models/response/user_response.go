package response

import "time"

type LoginResponse struct {
	Token     string `json:"token"`
	UserID    uint   `json:"user_id"`
	Email     string `json:"email"`
	ExpiresAt int64  `json:"expires_at"`
}

type UserResponse struct {
	ID                uint      `json:"id"`
	NIK               uint      `json:"nik"`
	NPWP              string    `json:"npwp"`
	Nama              string    `json:"nama"`
	Email             string    `json:"email"`
	Alamat            string    `json:"alamat"`
	JenisWajibPajak   string    `json:"jenis_wajib_pajak"`
	TanggalRegistrasi time.Time `json:"tanggal_registrasi"`
	StatusAktif       bool      `json:"status_aktif"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type PasswordResetResponse struct {
	Message string `json:"message"`
}
