package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":          4.99,
		"pie":           7.99,
		"salad":         6.99,
		"toffe pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// Looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// int as key type
	phoneBook := map[int]string{
		2468: "userOne",
		1234: "userTwo",
		3689: "userThree",
	}

	fmt.Println(phoneBook)
	fmt.Println(phoneBook[2468])

	phoneBook[1234] = "newUser"
	fmt.Println(phoneBook)

	phoneBook[3689] = "testUser"
	fmt.Println(phoneBook)
}
