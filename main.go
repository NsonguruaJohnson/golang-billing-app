package main

import (
	"fmt"
	"strings"
)

func main() {
	greeting := "Hello world changers"

	// fmt.Println(strings.Contains(greeting, "Hello"))
	// fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "w"))
	fmt.Println(strings.Split(greeting, " "))

	fmt.Println("Original string is", greeting)
}
