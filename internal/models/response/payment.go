package response

import "time"

type PaymentResponse struct {
	ID               uint      `json:"id"`
	UserID           uint      `json:"user_id"`
	JumlahBayar      float64   `json:"jumlah_bayar"`
	MetodePembayaran string    `json:"metode_pembayaran"`
	TanggalBayar     time.Time `json:"tanggal_bayar"`
	ReferensiSPTID   *uint     `json:"referensi_spt_id"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
