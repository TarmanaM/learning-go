package clean

type UserData []map[string]Biodata

type Biodata struct {
	ID        int8
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}
