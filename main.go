package main

import (
	"fmt"
	"sort"
)

func main() {
	// greeting := "Hello world people"

	// fmt.Println(strings.Contains(greeting, "Hello"))
	// fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "w"))
	// fmt.Println(strings.Split(greeting, "o"))

	// fmt.Println("Original string is", greeting)

	ages := []int{100, 15, 10, 40, 20, 30, 5, 25, 35, 40}
	sort.Ints(ages)

	fmt.Println(ages)

	index := sort.SearchInts(ages, 12)
	fmt.Println(index)

	names := []string{"All", "is", "well", "backend", "developer"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "xyz"))
	fmt.Println(sort.SearchStrings(names, "develop"))

}
