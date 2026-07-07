# Membuat Fungsi Generic di Go (Menggunakan `any` / `interface{}`)

Ketika menggunakan package `reflect` untuk melakukan inspeksi data (seperti membaca nama atau nilai field pada struct), seringkali kita ingin agar fungsi tersebut bisa menerima **tipe struct apa saja**, bukan hanya spesifik satu struct (seperti `Biodata`).

Untuk membuatnya generic (bisa menerima parameter tipe apa saja), kita bisa merubah *method* menjadi *function biasa* dan menggunakan tipe data `any` (atau `interface{}` di Go versi lama).

### Sebelum (Hanya untuk tipe Biodata)
```go
// Ini adalah Method, terikat HANYA pada struct Biodata
func (bio Biodata) InspectStruct() {
	val := reflect.ValueOf(bio)
	typ := reflect.TypeOf(bio)
    // ...
}
```
Jika ditulis seperti di atas, kita hanya bisa memanggilnya dengan cara: `dataBiodata.InspectStruct()`. Kita tidak bisa menggunakannya untuk struct lain.

### Sesudah (Generic untuk semua tipe)
```go
// Ini adalah Function biasa, menerima parameter 'data' bertipe 'any'
func InspectStruct(data any) {
	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)
	
	// Praktik yang baik: Pastikan data yang dimasukkan benar-benar sebuah Struct
	if typ.Kind() != reflect.Struct {
		fmt.Println("Error: Expected a struct")
		return
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)
		fmt.Printf("%s : %v\n", strings.ToLower(field.Name), value.Interface())
	}
}
```

### Cara Memanggilnya
Sekarang fungsi ini bisa digunakan untuk struct apa saja:

```go
type Mobil struct {
    Merk  string
    Tahun int
}

type User struct {
    Nama string
    Umur int
}

func main() {
    m := Mobil{Merk: "Toyota", Tahun: 2020}
    u := User{Nama: "Budi", Umur: 25}

    // Panggil fungsinya
    InspectStruct(m) 
    InspectStruct(u)
}
```

> **Catatan Tambahan**: Di Go versi 1.18+, ada fitur yang disebut **Generics** secara native (menggunakan bracket `[T any]`), namun jika tujuannya adalah menggunakan `reflect`, cara menggunakan parameter `data any` seperti di atas adalah pendekatan yang paling umum dan standar di Go.
