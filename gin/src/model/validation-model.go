package model

type CreateUserRequest struct {
	Nama            string `validate:"required,max=100"`
	Email           string `validate:"required,email"`
	Password        string `validate:"required,min=6"`
	NIK             string `validate:"required,len=16,numeric"`
	Alamat          string `validate:"required,max=100"`
	Telepon         string `validate:"required,max=15,numeric"`
	Jenis_kelamin   string `validate:"required,oneof=L P"`
	Kepala_keluarga int    `validate:"oneof=0 1" default:"0"`
	Tempat_lahir    string `validate:"required,max=50"`
	Tanggal_lahir   string `validate:"required,datetime=2006-01-02"`
	Jenis_usaha     string `validate:"required,max=50"`
}

type LoginUserRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6"`
}

type UpdateUserRequest struct {
	Nama            string `validate:"omitempty,max=100"`
	Email           string `validate:"omitempty,email"`
	Password        string `validate:"omitempty,min=6"`
	NIK             string `validate:"omitempty,len=16,numeric"`
	Alamat          string `validate:"omitempty,max=100"`
	Telepon         string `validate:"omitempty,max=15,numeric"`
	Jenis_kelamin   string `validate:"omitempty,oneof=L P"`
	Kepala_keluarga int    `validate:"omitempty,oneof=0 1"`
	Tempat_lahir    string `validate:"omitempty,max=50"`
	Tanggal_lahir   string `validate:"omitempty,datetime=2006-01-02"`
	Jenis_usaha     string `validate:"omitempty,max=50"`
}

type CreateAssistanceToolsRequest struct {
	Assistance_id string `validate:"required"`
	Tools_id      string `validate:"required"`
	Kuantitas     int    `validate:"required,number"`
}
