package request

import "time"

type CreatePayment struct {
	UserID           uint      `json:"user_id" validate:"required"`
	JumlahBayar      float64   `json:"jumlah_bayar" validate:"required"`
	MetodePembayaran string    `json:"metode_pembayaran" validate:"required"`
	TanggalBayar     time.Time `json:"tanggal_bayar" validate:"required"`
	ReferensiSPTID   *uint     `json:"referensi_spt_id" validate:"omitempty"`
	Status           string    `json:"status" validate:"omitempty,oneof=pending sukses"`
}

type UpdatePayment struct {
	JumlahBayar      float64   `json:"jumlah_bayar" validate:"omitempty"`
	MetodePembayaran string    `json:"metode_pembayaran" validate:"omitempty"`
	TanggalBayar     time.Time `json:"tanggal_bayar" validate:"omitempty"`
	ReferensiSPTID   *uint     `json:"referensi_spt_id" validate:"omitempty"`
	Status           string    `json:"status" validate:"omitempty,oneof=pending sukses"`
}
