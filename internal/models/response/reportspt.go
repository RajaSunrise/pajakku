package response

import "time"

type TaxReturnResponse struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	JenisSPT     string    `json:"jenis_spt"`
	PeriodePajak string    `json:"periode_pajak"`
	JumlahPajak  float64   `json:"jumlah_pajak"`
	Status       string    `json:"status"`
	FileSPT      string    `json:"file_spt"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
