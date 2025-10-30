package response

import "time"

type AuditLogResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Aksi      string    `json:"aksi"`
	Timestamp time.Time `json:"timestamp"`
	IPAddress string    `json:"ip_address"`
	OldValue  string    `json:"old_value"`
	NewValue  string    `json:"new_value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
