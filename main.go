package main

import "fmt"

func main() {
	// Arrays
	// var ages [3]int = [3]int{20, 23, 35}
	var ages = [3]int{20, 23, 35}

	names := [4]string{"Nsongurua", "Akpan", "Jimmy", "Johnson"}
	names[1] = "Testing"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// Slices(use arrays under the hood)
	var scores = []int{100, 50, 80}
	scores[2] = 25
	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

	// Slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "koppa")
	fmt.Println(rangeOne)

}
