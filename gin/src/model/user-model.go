package model

type User struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	NIK             string `json:"nik"`
	Alamat          string `json:"alamat"`
	Telepon         string `json:"telepon"`
	Jenis_kelamin   string `json:"jenis_kelamin"`
	Kepala_keluarga string `json:"kepala_keluarga"`
	Tempat_lahir    string `json:"tempat_lahir"`
	Tanggal_lahir   string `json:"tanggal_lahir"`
	Jenis_usaha     string `json:"jenis_usaha"`
}

type UserSavedNews struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	News  []News `json:"news"`
}

type UserSavedNewsComment struct {
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	NewsComments []NewsComments `json:"news"`
}

type UserFacilities struct {
	Name       string       `json:"nama"`
	Email      string       `json:"email"`
	Sertifikat []Sertifikat `json:"sertifikat"`
	Pelatihan  []Pelatihan  `json:"pelatihan"`
	Bantuan    []Bantuan    `json:"bantuan"`
}
