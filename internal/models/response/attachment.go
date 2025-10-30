package response

import "time"

type AttachmentResponse struct {
	ID           uint      `json:"id"`
	NamaFile     string    `json:"nama_file"`
	PathURL      string    `json:"path_url"`
	TipeMime     string    `json:"tipe_mime"`
	Ukuran       int64     `json:"ukuran"`
	RelatedModel string    `json:"related_model"`
	UserID       *uint     `json:"user_id"`
	TaxReturnID  *uint     `json:"tax_return_id"`
	InvoiceID    *uint     `json:"invoice_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
