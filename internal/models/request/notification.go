package request

import "time"

type CreateNotification struct {
	UserID       uint      `json:"user_id" validate:"required"`
	Judul        string    `json:"judul" validate:"required"`
	Isi          string    `json:"isi" validate:"required"`
	Tipe         string    `json:"tipe" validate:"required,oneof=email push"`
	StatusBaca   bool      `json:"status_baca" validate:"omitempty"`
	TanggalKirim time.Time `json:"tanggal_kirim" validate:"required"`
	TaxReturnID  *uint     `json:"tax_return_id" validate:"omitempty"`
	PaymentID    *uint     `json:"payment_id" validate:"omitempty"`
}

type UpdateNotification struct {
	Judul        string    `json:"judul" validate:"omitempty"`
	Isi          string    `json:"isi" validate:"omitempty"`
	Tipe         string    `json:"tipe" validate:"omitempty,oneof=email push"`
	StatusBaca   bool      `json:"status_baca" validate:"omitempty"`
	TanggalKirim time.Time `json:"tanggal_kirim" validate:"omitempty"`
	TaxReturnID  *uint     `json:"tax_return_id" validate:"omitempty"`
	PaymentID    *uint     `json:"payment_id" validate:"omitempty"`
}
