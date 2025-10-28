package request

import "time"

type CreateBillingRequest struct {
	KodeBilling        string    `json:"kode_billing" validate:"required,min=1,max=30"`
	KodeAkunPajak      string    `json:"kode_akun_pajak" validate:"required,min=1,max=10"`
	KodeJenisSetoran   string    `json:"kode_jenis_setoran" validate:"required,min=1,max=10"`
	MasaPajak          int       `json:"masa_pajak" validate:"required,min=1,max=12"`
	TahunPajak         int       `json:"tahun_pajak" validate:"required,min=2000"`
	JumlahSetor        int64     `json:"jumlah_setor" validate:"required,min=1"`
	TanggalKadaluwarsa time.Time `json:"tanggal_kadaluwarsa" validate:"required"`
}

type UpdateBillingRequest struct {
	StatusPembayaran string `json:"status_pembayaran" validate:"omitempty,oneof=Belum Dibayar Dibayar"`
	NTPN             string `json:"ntpn" validate:"omitempty,min=1,max=100"`
}
