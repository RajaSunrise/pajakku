package response

import "time"

type BillingResponse struct {
	ID                 uint      `json:"id"`
	KodeBilling        string    `json:"kode_billing"`
	JumlahSetor        int64     `json:"jumlah_setor"`
	MasaPajak          int       `json:"masa_pajak"`
	TahunPajak         int       `json:"tahun_pajak"`
	StatusPembayaran   string    `json:"status_pembayaran"`
	TanggalKadaluwarsa time.Time `json:"tanggal_kadaluwarsa"`
	CreatedAt          time.Time `json:"created_at"`
}
