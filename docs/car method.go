package main

import "fmt"

type Car struct {
	name           string
	yearProduction uint16
	kilometer      uint16
}

func (c Car) Drive() {
	fmt.Printf("%s %d Running... to %d", c.name, c.yearProduction, c.kilometer+10)
}

func carMethod() {

	car1 := Car{
		name: "BMW Mk4", yearProduction: 1999, kilometer: 2000,
	}

	car1.Drive()

}
