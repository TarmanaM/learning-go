package main

import "fmt"

func main2() {

	var data = []User{
		{
			Name: "Budi",
			Age:  25,
		},
		{
			Name: "Andi",
			Age:  30,
		},
	}
	fmt.Print(data)

	nestedData := map[string]any{
		"number": 2,
		"string": "tulisan",
		"array":  [3]int8{1, 2, 3},
		"slice":  []int8{1},
		"make":   make([]int8, 2, 10),
		"data":   []map[string]any{},
		"nested": map[string]any{"nested2": map[string]any{}},
	}

	_ = nestedData

}

type User2 struct {
	Name string
	Age  int
}
