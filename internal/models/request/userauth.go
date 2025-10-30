package request

// Request For User
type CreateUser struct {
	NIK             uint   `json:"nik" validate:"required"`
	NPWP            string `json:"npwp" validate:"required,min=15,max=25"`
	Nama            string `json:"nama" validate:"required,min=3,max=255"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8,max=100"`
	Alamat          string `json:"alamat" validate:"required"`
	JenisWajibPajak string `json:"jenis_wajib_pajak" validate:"required,oneof=badan individu instansi"`
	RoleID          uint   `json:"role_id" validate:"required"`
}

type SignupRequest struct {
	NIK             uint   `json:"nik" validate:"required"`
	NPWP            string `json:"npwp" validate:"required,min=15,max=25"`
	Nama            string `json:"nama" validate:"required,min=3,max=255"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8,max=100"`
	Alamat          string `json:"alamat" validate:"required"`
	JenisWajibPajak string `json:"jenis_wajib_pajak" validate:"required,oneof=badan individu instansi"`
}

type UpdateUser struct {
	NIK             uint   `json:"nik" validate:"omitempty"`
	NPWP            string `json:"npwp" validate:"omitempty,min=15,max=25"`
	Nama            string `json:"nama" validate:"omitempty,min=3,max=255"`
	Email           string `json:"email" validate:"omitempty,email"`
	Password        string `json:"password" validate:"omitempty,min=8,max=100"`
	Alamat          string `json:"alamat" validate:"omitempty"`
	JenisWajibPajak string `json:"jenis_wajib_pajak" validate:"omitempty,oneof=badan individu instansi"`
	RoleID          uint   `json:"role_id" validate:"omitempty"`
	StatusAktif     *bool  `json:"status_aktif" validate:"omitempty"`
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
