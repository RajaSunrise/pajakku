package request

type CreateRole struct {
	NamaRole    string `json:"nama_role" validate:"required,min=2,max=50"`
	Permissions string `json:"permissions" validate:"required,json"`
}

type UpdateRole struct {
	NamaRole    string `json:"nama_role" validate:"omitempty,min=2,max=50"`
	Permissions string `json:"permissions" validate:"omitempty,json"`
}
