package response

import "time"

type LoginResponse struct {
	Token     string `json:"token"`
	UserID    string `json:"user_id"`
	Email     string `json:"email"`
	ExpiresAt int64  `json:"expires_at"`
}

type UserProfileResponse struct {
	NIK            uint    `json:"nik"`
	NPWP           string    `json:"npwp"`
	NamaWajibPajak string    `json:"nama_wajib_pajak"`
	TipeWajibPajak string    `json:"tipe_wajib_pajak"`
	NomorTelepon   string    `json:"nomor_telepon"`
	AlamatLengkap  string    `json:"alamat_lengkap"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type UserAuthResponse struct {
	ID            string     `json:"id"`
	UserProfileID uint       `json:"user_profile_id"`
	FotoProfil    string     `json:"foto_profil"`
	Email         string     `json:"email"`
	StatusAkun    string     `json:"status_akun"`
	TerakhirLogin *time.Time `json:"terakhir_login"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

type PasswordResetResponse struct {
	Message string `json:"message"`
}
