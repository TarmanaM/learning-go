package helpers

type Car struct {
	Name           string
	YearProduction uint16
	Kilometer      uint16
	TheDriver      *Driver
}

type Driver struct {
	Name    string
	Age     uint8
	UsedCar *Car
}
