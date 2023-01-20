package main

import "fmt"

func updateName(x string) string {
	x = "Plucky"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {
	// group A types -> strings, ints, bools, floats, arrays, structs
	name := "Brainy"

	name = updateName(name)
	fmt.Println(name)

	// group B types -> slices, maps, functions
	menu := map[string]float64{
		"pie":       5.95,
		"Ice cream": 3.99,
	}
	updateMenu(menu)
	fmt.Println(menu)
}
