package helpers

import "fmt"

func (car *Car) Drive() {
	if car.TheDriver == nil {
		fmt.Println("Mobil tidak ada driver")
		return
	}
	fmt.Printf("%s %d Running... to %d by %s\n", car.Name, car.YearProduction, car.Kilometer+10, car.TheDriver.Name)
}

// 1. Mengisi Relasi (GunakanMobil)
// Perhatikan kita menggunakan `*Driver` (Pointer Receiver) agar data aslinya berubah, bukan copy-nya.
// Parameternya juga menerima `*Car` (Alamat memori mobil)
func (driver *Driver) GunakanMobil(targetCar *Car) {
	// Isi pointer mobil ke driver
	driver.UsedCar = targetCar
	// Isi pointer driver ke mobil
	targetCar.TheDriver = driver

	fmt.Printf("%s mulai menggunakan mobil %s\n", driver.Name, targetCar.Name)
}

// 2. Mengubah Relasi (GantiMobil)
func (driver *Driver) GantiMobil(newCar *Car) {
	fmt.Printf("%s mengganti mobil dari %s ke %s\n", driver.Name, driver.UsedCar.Name, newCar.Name)

	// Putuskan relasi dengan mobil lama
	if driver.UsedCar != nil {
		driver.UsedCar.TheDriver = nil
	}

	// Gunakan mobil baru
	driver.GunakanMobil(newCar)
}

// 3. Menghapus Relasi (KeluarMobil)
func (driver *Driver) KeluarMobil() {
	if driver.UsedCar != nil {
		fmt.Printf("%s keluar dari mobil %s\n", driver.Name, driver.UsedCar.Name)

		// Putuskan relasi dari sisi mobil
		driver.UsedCar.TheDriver = nil
		// Putuskan relasi dari sisi driver
		driver.UsedCar = nil
	} else {
		fmt.Printf("%s memang sedang tidak menggunakan mobil apapun\n", driver.Name)
	}
}
