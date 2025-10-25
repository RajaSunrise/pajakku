package request

type CreateUsersProfile struct {
	NPWP           string `json:"npwp" validate:"required, min=8, max=100"`
	NamaWajibPajak string `json:"nama_wajib_pajak" validate:"required, min=8, max=100"`
	TipeWajibPajak string `json:"tipe_wajib_pajak" validate:"required, min=6, max=50"`
	NomorTelepon   string `json:"nomor_telepon" validate:"required, min=8, max=16"`
	AlamatLengkap  string `json:"alamat_lengkap" validate:"required"`
}

type UpdateUsersProfile struct {
	NPWP           string `json:"npwp" validate:"omitempty, min=8, max=100"`
	NamaWajibPajak string `json:"nama_wajib_pajak" validate:"omitempty, min=8, max=100"`
	TipeWajibPajak string `json:"tipe_wajib_pajak" validate:"omitempty, min=6, max=50"`
	NomorTelepon   string `json:"nomor_telepon" validate:"omitempty, min=8, max=16"`
	AlamatLengkap  string `json:"alamat_lengkap" validate:"omitempty"`
}
