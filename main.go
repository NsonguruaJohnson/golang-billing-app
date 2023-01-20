package main

import "fmt"

func updateName(x *string) {
	*x = "Plucky"
}

func main() {
	name := "Brainy"

	// updateName(name)

	// fmt.Println("Memory address of name is: ", &name)

	m := &name

	// fmt.Println("Memory address:", m)
	// fmt.Println("Value at memory address:", *m)
	fmt.Println(name)
	updateName(m)
	fmt.Println(name)
}
