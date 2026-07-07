package main

import (
	"fmt"
	"reflect"
)

func main() {

	var number int8 = 2
	var reflect = reflect.ValueOf(number)

	fmt.Print(" tipe var ", reflect.Type())

}
