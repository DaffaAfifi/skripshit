package model

type Sertifikat struct {
	Id             NullString `json:"id"`
	Name           NullString `json:"nama"`
	Tanggal_terbit NullTime   `json:"tanggal_terbit"`
	Kadaluarsa     NullTime   `json:"kadaluarsa"`
	Keterangan     NullString `json:"keterangan"`
	No_sertifikat  NullString `json:"no_sertifikat"`
}

type Pelatihan struct {
	Id                  NullString `json:"id"`
	Name                NullString `json:"name"`
	Penyelenggara       NullString `json:"penyelenggara"`
	Tanggal_pelaksanaan NullTime   `json:"tanggal_pelaksanaan"`
	Tempat              NullString `json:"tempat"`
}

type Bantuan struct {
	Id              NullString  `json:"id"`
	Name            NullString  `json:"name"`
	Koordinator     NullString  `json:"koordinator"`
	Sumber_anggaran NullString  `json:"sumber_anggaran"`
	Tahun_pemberian NullTime    `json:"tahun_pemberian"`
	Total_anggaran  NullFloat64 `json:"total_anggaran"`
	Alat            []Alat      `json:"alat"`
}

type Alat struct {
	Id        NullString  `json:"id"`
	Name      NullString  `json:"name"`
	Harga     NullFloat64 `json:"harga"`
	Deskripsi NullString  `json:"deskripsi"`
	Kuantitas NullInt64   `json:"kuantitas"`
}
