package request

// Request For UsersAuth
type CreateUsersAuth struct {
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required, min=8, max=100"`
	Role     string `json:"role" validate:"omitempty,oneof=admin users"`
}

type UpdateUsersAuth struct {
	Email    string `json:"email" validate:"omitempty, email"`
	Password string `json:"password" validate:"omitempty, min=8, max=100"`
	Role     string `json:"role" validate:"omitempty,oneof=admin users"`
}

type UpdateFotoProfile struct {
	FotoProfile string `json:"foto_profile" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

type ForgetPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetPasswordRequest struct {
	Token    string `json:"token" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}
