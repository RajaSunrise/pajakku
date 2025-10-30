package response

import "time"

type NotificationResponse struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	Judul        string    `json:"judul"`
	Isi          string    `json:"isi"`
	Tipe         string    `json:"tipe"`
	StatusBaca   bool      `json:"status_baca"`
	TanggalKirim time.Time `json:"tanggal_kirim"`
	TaxReturnID  *uint     `json:"tax_return_id"`
	PaymentID    *uint     `json:"payment_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
