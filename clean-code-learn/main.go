package main

import "clean-code-learn/helpers"

func main() {
	// Buat Mobil 1
	myCar := helpers.Car{
		Name:           "Toyota Supra",
		YearProduction: 2022,
		Kilometer:      5000,
	}

	// Buat Mobil 2
	mySecondCar := helpers.Car{
		Name:           "Honda Civic",
		YearProduction: 2020,
		Kilometer:      10000,
	}

	// Buat drivernya
	myDriver := helpers.Driver{
		Name: "Isa Tarmana",
		Age:  25,
	}

	myCar.Drive()
	myDriver.KeluarMobil()

	// 1. Uji Coba Mengisi (GunakanMobil)
	// Ingat untuk menggunakan `&` (address of) karena parameternya adalah Pointer (*Car)
	myDriver.GunakanMobil(&myCar)
	myCar.Drive() // Menjalankan mobil

	// 2. Uji Coba Mengubah (GantiMobil)
	myDriver.GantiMobil(&mySecondCar)

	// 3. Uji Coba Menghapus (KeluarMobil)
	myDriver.KeluarMobil()
}
