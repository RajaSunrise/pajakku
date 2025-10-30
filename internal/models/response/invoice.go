package response

import "time"

type InvoiceResponse struct {
	ID               uint      `json:"id"`
	NomorFaktur      string    `json:"nomor_faktur"`
	UserID           uint      `json:"user_id"`
	TanggalTransaksi time.Time `json:"tanggal_transaksi"`
	Jumlah           float64   `json:"jumlah"`
	Jenis            string    `json:"jenis"`
	StatusVerifikasi string    `json:"status_verifikasi"`
	TaxReturnID      *uint     `json:"tax_return_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
