package request

type CreateTaxType struct {
	KodePajak    string  `json:"kode_pajak" validate:"required,min=1,max=20"`
	Nama         string  `json:"nama" validate:"required,min=1,max=100"`
	TarifDefault float64 `json:"tarif_default" validate:"required"`
	Deskripsi    string  `json:"deskripsi" validate:"omitempty"`
}

type UpdateTaxType struct {
	KodePajak    string  `json:"kode_pajak" validate:"omitempty,min=1,max=20"`
	Nama         string  `json:"nama" validate:"omitempty,min=1,max=100"`
	TarifDefault float64 `json:"tarif_default" validate:"omitempty"`
	Deskripsi    string  `json:"deskripsi" validate:"omitempty"`
}
