package main

import "fmt"

func main() {
	// var age [3]int = [3]int{5,10,15}
	var ages = [3]int{10, 20, 30}

	names := [4]string{"Nsongurua", "Akpan", "Jimmy", "Johnson"}
	// fmt.Println(names, len(names))s

	names[1] = "Favour"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// Slices(uses arrays under the hood)
	var scores = []int{100, 40, 30}
	scores[2] = 400
	scores = append(scores, 22)

	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "kahoots")
	fmt.Println(rangeOne)

}
