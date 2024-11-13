package model

type News struct {
	NewsId    string     `json:"news_id"`
	Gambar    NullString `json:"gambar"`
	Judul     string     `json:"judul"`
	Subjudul  string     `json:"sub_judul"`
	Isi       string     `json:"isi"`
	CreatedAt string     `json:"created_at"`
}

type NewsComments struct {
	NewsId    string     `json:"news_id"`
	Gambar    NullString `json:"gambar"`
	Judul     string     `json:"judul"`
	Subjudul  string     `json:"sub_judul"`
	Isi       string     `json:"isi"`
	CreatedAt string     `json:"created_at"`
	Comments  []Comments `json:"comments"`
}

type Comments struct {
	Comment   string `json:"comment"`
	User      string `json:"user"`
	CreatedAt string `json:"created_at"`
}
