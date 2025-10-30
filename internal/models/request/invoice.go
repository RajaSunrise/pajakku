package request

import "time"

type CreateInvoice struct {
	NomorFaktur      string    `json:"nomor_faktur" validate:"required"`
	UserID           uint      `json:"user_id" validate:"required"`
	TanggalTransaksi time.Time `json:"tanggal_transaksi" validate:"required"`
	Jumlah           float64   `json:"jumlah" validate:"required"`
	Jenis            string    `json:"jenis" validate:"required,oneof=masuk keluar"`
	StatusVerifikasi string    `json:"status_verifikasi" validate:"omitempty,oneof=pending verified rejected"`
	TaxReturnID      *uint     `json:"tax_return_id" validate:"omitempty"`
}

type UpdateInvoice struct {
	NomorFaktur      string    `json:"nomor_faktur" validate:"omitempty"`
	TanggalTransaksi time.Time `json:"tanggal_transaksi" validate:"omitempty"`
	Jumlah           float64   `json:"jumlah" validate:"omitempty"`
	Jenis            string    `json:"jenis" validate:"omitempty,oneof=masuk keluar"`
	StatusVerifikasi string    `json:"status_verifikasi" validate:"omitempty,oneof=pending verified rejected"`
	TaxReturnID      *uint     `json:"tax_return_id" validate:"omitempty"`
}
