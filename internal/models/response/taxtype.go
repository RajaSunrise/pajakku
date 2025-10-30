package response

import "time"

type TaxTypeResponse struct {
	ID           uint      `json:"id"`
	KodePajak    string    `json:"kode_pajak"`
	Nama         string    `json:"nama"`
	TarifDefault float64   `json:"tarif_default"`
	Deskripsi    string    `json:"deskripsi"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
