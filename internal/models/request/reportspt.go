package request

type CreateTaxReturn struct {
	JenisSPT     string  `json:"jenis_spt" validate:"required,oneof=tahunan masa"`
	PeriodePajak string  `json:"periode_pajak" validate:"required"`
	JumlahPajak  float64 `json:"jumlah_pajak" validate:"required"`
	Status       string  `json:"status" validate:"omitempty,oneof=draft submit disetujui"`
	FileSPT      string  `json:"file_spt" validate:"omitempty"`
}

type UpdateTaxReturn struct {
	JenisSPT     string  `json:"jenis_spt" validate:"omitempty,oneof=tahunan masa"`
	PeriodePajak string  `json:"periode_pajak" validate:"omitempty"`
	JumlahPajak  float64 `json:"jumlah_pajak" validate:"omitempty"`
	Status       string  `json:"status" validate:"omitempty,oneof=draft submit disetujui"`
	FileSPT      string  `json:"file_spt" validate:"omitempty"`
}
