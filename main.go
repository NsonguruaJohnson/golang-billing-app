package main

import "fmt"

func main() {
	age := 25

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	// if age < 30 {
	// 	fmt.Println("Age is less than 30")
	// } else if age < 40 {
	// 	fmt.Println("Age is less than 40")
	// } else {
	// 	fmt.Println("Age is not less than 45")
	// }

	names := []string{"Nsongurua", "Akpan", "Johnson", "Excellent", "Loyal"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("Continuing at position \n", index)
			continue //break from that particular iteration to the top of the loop
		}

		if index > 2 {
			fmt.Println("Breaking at pos", index)
			break // break out of a loop completely
		}

		fmt.Printf("The value at pos %v is %v \n", index, value)
	}

}
