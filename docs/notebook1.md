## Clean Code Type

### Rule

Pada Go setiap folder adalah package. jadi file-file yang ada di sebuah folder harus diawali dengan `package <nama_folder>`.

### Method

merupakan fungsi yang menempel dengan struct

```go
type Car struct {
	name           string
	yearProduction uint16
	kilometer      uint16
}

func (c Car) Drive() {
	fmt.Printf("%s %d Running... to %d", c.name, c.yearProduction, c.kilometer+10)
}

func main() {

	car1 := Car{
		name: "BMW Mk4", yearProduction: 1999, kilometer: 2000,
	}

	car1.Drive()

}
```

## Array

- fixed length = jadi harus di difine panjangnya serta tipe datanya
- tidak bisa di tambah atau append

```go
var data [4]int{1,2,3,4}
```

- jika ingin di duplikasi maka akan membuat memori baru conoth

```go
var data [4]int{1,2,3,4}
data2 := data
```

maka data2 tidak sama dengan data secara memori atau berdasarkan komputasinya. tapi isi dari array tetap sama.

### function Make

fungsi yang membuat sebuah array yang memiliki nilai default jika kosong

```
data := make([]int8,2,10) // make(type, length, capacity)
data[0] = 2
fmt.Print(data[1:8]) // maka outputnya [0 0 0 0 0 0 0]
fmt.Print(data[1:22]) // maka outputnya error karena melebihi capacity
```

### Slice

- Mirip Array tapi lebih flexible.
- tidak fixed length

- view terhadap array
- memiliki 3 komponen yaitu pointer, length, dan capacity.
  Pointer adalah alamat memori awal dimana slice tersebut menunjuk.
  Length adalah panjang slice yang bisa diakses saat ini.
  Capacity adalah panjang memori penuh yang bisa di akses

```
var data = []string{"ddd","sss"}

// dibelakang bekerja seperti
data Pointer : Array {"ddd","sss"}
data Length : 2
data Capacity : 2
```

#### Append

- jika slice melakukan append maka

```
data := []int{1,2,3}
// dibelakang
data Pointer : Array {1,2,3}
data Length : 3
data Capacity : 3

data = append(data, 23, 33)
// maka bekerja dibelakang
Array {1,2,3} -> array lama di copy menjadi array baru. dan yang lama di hapus
Array {1,2,3,23,33} -> array baru

data Pointer : Array {1,2,3,23,33}
data Length : 5
data Capacity : 6
```

- Append slice dengan slice

gunakan `...` pada slice kedua agar isinya bisa masuk. jikatidak mengunakan `...` maka slice kedua akan dianggap sebagai satu kesatuan.

```
data :=[]int{1,2,3}

data2 :=[]int{4,5,6}

data = append(data, data2...)
// maka bekerja dibelakang
Array {1,2,3} -> array lama di copy menjadi array baru. dan yang lama di hapus
Array {1,2,3,4,5,6} -> array baru

data Pointer : Array {1,2,3,4,5,6}
data Length : 6
data Capacity : 7
```

## Looping

### first way for

```go
for i:=0; i<3;i++{
   fmt.Println("Angka",i)
}
```

### second way for

```go
var i = 0

for i <3{
   fmt.Println("Angka",i)
   i++
}
```

### third way for

```go
for{
   fmt.Println("Angka",i)
   i++
   if i ==3 {
      break
   }
}
```

## Typing

### Tipe Data

1. Basic Type
   - number
   - string
   - boolean
2. Aggregate Type
   - array dan struct
3. Reference Type
   - slice, pointer, map, function, channel
4. Interface Type: interface`

#### Integer

1. int
   berisikan any format integer. semua angka bisa masuk baik negatif atau positif bilangan bulat. namun boros memori jika penggnaan terlalu redundant

sebaiknya gunakan sesuai kebutuhan :

2. uint8 (0~255)
3. uint16 (0~65535)
4. uint32 (0 to 4294967295)
5. uint64 (0 to 18446744073709551615)
6. int8 (-128 to 127)
7. int16 (-32768 to 32767)
8. int32 (-2.147.483.648 sampai 2.147.483.647)
9. int64 (-9.223.372.036.854.775.808 ~ 9.223.372.036.854.775.807)

#### Float

1. float32

2. float64

### camelCase

camelCase — camelCase adalah konvensi penamaan di mana huruf pertama dari setiap kata dalam kata majemuk akan dikapitalisasi KECUALI untuk huruf pertama di awal kata

Kapitalisasi akronim — misal : HTTP, DNS, URL
Perhatikan huruf “f” diawal ditulis
dengan huruf kecil, sedangkan
untuk huruf “C” dan “S” ditulis
huruf besar

#### GO gin gonic

merupakan framework Websitenya GO

#### Short Declaration

Penggunaan saat variable tidak menentu tipe datanya

```
var age = 62

//menjadi
age := "62" //or 62
```

#### penggumanaan Print format

1. `%s` digunakan untuk print variable
2. `%T` digunakan tipe data
3. `%d` digunakan untuk print integer
4. `%v` digunakan untuk print isi semua tipe data
5. `%t` digunakan untuk print isi Boolean
6. `%+v` digunakan untuk print isi semua tipe data dan nama field
7. `%#v` digunakan untuk print isi semua tipe data dan nama field dalam format Go-syntax

#### Underscore Variable

Bahasa Golang memiliki satu aturan yang cukup
strict/ketat, yang dimana tidak ada variable yang boleh
menganggur ketika sudah kita buat.

## Inisitation

1. Go mod Init
   CLI membuat project atau module go. menghasilkan file go.mod. file tersebut seperti package.json atau file dari identitas project tersebut

## Recursive Struct (Struct Saling Berelasi)

Jika kita membuat dua `struct` yang saling memiliki field satu sama lain (seperti `Car` memiliki field tipe `Driver`, dan `Driver` memiliki field tipe `Car`), maka akan terjadi error: `invalid recursive type`.

**Penyebab:**
Go compiler perlu mengetahui seberapa besar alokasi memori yang dibutuhkan oleh setiap struct saat kompilasi. Jika `Car` mengandung `Driver` secara langsung dan sebaliknya, maka ukuran datanya tidak terhingga (infinite loop size).

**Solusi:**
Gunakan **Pointer** (`*`) pada salah satu atau kedua field yang saling berelasi. Ukuran memori dari sebuah pointer sudah pasti, sehingga compiler bisa memprosesnya.

```go
type Car struct {
	name           string
	yearProduction uint16
	kilometer      uint16
	theDriver      *Driver // <- Menggunakan pointer (alamat memori)
}

type Driver struct {
	name    string
	age     uint8
	usedCar *Car    // <- Menggunakan pointer
}
```

### Cara Mengisi, Mengubah, dan Menghapus Relasi Pointer Struct

Karena field yang kita gunakan adalah sebuah **Pointer**, maka untuk berinteraksi dengan field tersebut (mengisi/mengubah), kita harus memberikan **alamat memori** menggunakan simbol ampersand (`&`). Untuk menghapusnya, kita berikan nilai kosong untuk pointer, yaitu `nil`.

#### 1. Mengisi (Assigning) Ikatan
Gunakan simbol `&` di depan variabel struct untuk mengambil alamat memorinya, lalu masukkan ke dalam field pointer.

```go
// 1. Buat object masing-masing terlebih dahulu
myCar := Car{
    name: "Toyota Supra",
    yearProduction: 2022,
    kilometer: 5000,
}

myDriver := Driver{
    name: "Isa Tarmana",
    age: 25,
}

// 2. Isi (hubungkan) ikatannya
myCar.theDriver = &myDriver // Menunjuk ke alamat memori myDriver
myDriver.usedCar = &myCar   // Menunjuk ke alamat memori myCar
```

#### 2. Mengubah (Updating) Ikatan
Sama seperti mengisi, kamu cukup menimpa field tersebut dengan alamat memori dari struct yang baru.

```go
newDriver := Driver{
    name: "Budi",
    age: 30,
}

// Mengubah driver dari mobil tersebut ke Budi
myCar.theDriver = &newDriver 

// Jangan lupa update mobil si Budi juga jika diperlukan
newDriver.usedCar = &myCar
```

#### 3. Menghapus (Removing) Ikatan
Dalam Go, "kosong" untuk sebuah pointer adalah `nil`. Jadi, untuk menghapus atau memutuskan ikatan, cukup beri nilai `nil`.

```go
// Menghapus ikatan driver pada mobil (mobil jadi tidak punya driver)
myCar.theDriver = nil

// Menghapus ikatan mobil pada driver (driver tidak sedang menyetir mobil)
myDriver.usedCar = nil
```
