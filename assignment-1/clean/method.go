package clean

import (
	"fmt"
	"reflect"
	"strings"
)

func InspectStruct(data any) {
	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)

	// Pastikan yang dimasukkan adalah struct
	if typ.Kind() != reflect.Struct {
		fmt.Println("Error: Expected a struct")
		return
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i) // Field metadata
		value := val.Field(i) // Field value
		fmt.Printf("%s : %v\n", strings.ToLower(field.Name), value.Interface())
	}
}
