package model

type JadwalMatkul struct {
	Kode       string 	`json:"kode"`
	NamaMatkul string `json:"nama_matkul"`
	Hari		string `json:"hari"`
	WaktuMulai	string `json:"waktu_mulai"`
	WaktuSelesai string	`json:"waktu_selesai"`
	Dosen      string `json:"dosen"`
	Nidn		string	`json:"nidn"`
	Ruangan    string `json:"ruangan"`
}