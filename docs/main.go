package main

import "fmt"

func main() {
	/*Komentar
	 */

	const x uint8 = 4

	switch {
	case x == 5:
		fmt.Printf("get %d", x)
	case (x < 10) && (x > 3):
		fmt.Printf("get %d. the value is under ten up to 3", x)
		fallthrough
	case x > 5:
		fmt.Printf("get %d. more than five", x)
	case x >= 4:
		fmt.Printf("get %d. more than 4", x)
	default:
		{
			fmt.Println("unknown")
		}
	}

	// var name string // isa is default value
	// var age int

	// u := User{Name: "Alice", Age: 30}

	// // Default formatting
	// fmt.Printf("Using %%v:  %v\n", u)
	// // Output: {Alice 30}

	// // Plus flag (adds struct field names)
	// fmt.Printf("Using %%+v: %+v\n", u)
	// // Output: {Name:Alice Age:30}

	// // Go-syntax format
	// fmt.Printf("Using %%#v: %#v\n", u)
	// // Output: main.User{Name:"Alice", Age:30}

	// fmt.Println("Halo, selamat datang di dunia Go!")
	// fmt.Println("Ini pesan Baris kedua")
	// fmt.Println(name, age)

	// varAny := 62
	// fmt.Printf("%v", varAny)

}

type User struct {
	Name string
	Age  int
}
