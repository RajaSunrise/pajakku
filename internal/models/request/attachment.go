package request

type CreateAttachment struct {
	NamaFile     string `json:"nama_file" validate:"required"`
	PathURL      string `json:"path_url" validate:"required"`
	TipeMime     string `json:"tipe_mime" validate:"required"`
	Ukuran       int64  `json:"ukuran" validate:"required"`
	RelatedModel string `json:"related_model" validate:"required"`
	UserID       *uint  `json:"user_id" validate:"omitempty"`
	TaxReturnID  *uint  `json:"tax_return_id" validate:"omitempty"`
	InvoiceID    *uint  `json:"invoice_id" validate:"omitempty"`
}

type UpdateAttachment struct {
	NamaFile     string `json:"nama_file" validate:"omitempty"`
	PathURL      string `json:"path_url" validate:"omitempty"`
	TipeMime     string `json:"tipe_mime" validate:"omitempty"`
	Ukuran       int64  `json:"ukuran" validate:"omitempty"`
	RelatedModel string `json:"related_model" validate:"omitempty"`
	UserID       *uint  `json:"user_id" validate:"omitempty"`
	TaxReturnID  *uint  `json:"tax_return_id" validate:"omitempty"`
	InvoiceID    *uint  `json:"invoice_id" validate:"omitempty"`
}
