package models

type (
	DaftarHp struct {
		Id                 int    `json:"id"`
		Imei               string `json:"imei"`
		OS                 string `json:"os"`
		Brand              string `json:"brand"`
		Model              string `json:"model"`
		Tanggal_Daftar     string `json:"tanggal_daftar"`
		Status_Aktif       string `json:"status_aktif"`
		Nama               string `json:"nama"`
		Tanggal_Verifikasi string `json:"tanggal_verifikasi"`
		User_Verifikasi    string `json:"user_verifikasi"`
	}
)
