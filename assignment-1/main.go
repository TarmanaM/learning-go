package main

import (
	"assignment-1/clean"
	"fmt"
	"os"
)

func main() {
	data := clean.Biodata{
		ID:        0,
		Nama:      "Thomas",
		Pekerjaan: "Kota A",
		Alamat:    "programmer",
		Alasan:    "Alasan Thomas",
	}

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: go run main.go <Name> ")
		os.Exit(1)
	}

	if args[0] == data.Nama {
		clean.InspectStruct(data)
	}

	// _ = data
	fmt.Print("end")
}
