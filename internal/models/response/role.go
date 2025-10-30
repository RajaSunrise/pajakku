package response

import "time"

type RoleResponse struct {
	ID          uint      `json:"id"`
	NamaRole    string    `json:"nama_role"`
	Permissions string    `json:"permissions"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
