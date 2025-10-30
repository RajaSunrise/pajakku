package request

type CreateAuditLog struct {
	UserID    uint   `json:"user_id" validate:"required"`
	Aksi      string `json:"aksi" validate:"required"`
	IPAddress string `json:"ip_address" validate:"omitempty"`
	OldValue  string `json:"old_value" validate:"omitempty,json"`
	NewValue  string `json:"new_value" validate:"omitempty,json"`
}
