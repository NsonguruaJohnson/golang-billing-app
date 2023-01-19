package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	// sayGreeting("Nsongurua")
	// sayGreeting("Johnson")
	// sayBye("Boss")

	cycleNames([]string{"Tom", "Jerry", "Chase", "Ryder"}, sayGreeting)
	cycleNames([]string{"Tom", "Jerry", "Chase", "Ryder"}, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(20)

	fmt.Println(a1, a2)
	fmt.Printf("%0.3f, %0.2f \n", a1, a2)

}
