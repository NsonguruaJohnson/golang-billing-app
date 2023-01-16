package main

import "fmt"

func main() {

	age := 24
	name := "Plucky"

	// Print
	fmt.Print("Hello,")
	fmt.Print("World! \n")
	fmt.Print("New line \n")

	// Println
	fmt.Println("Hello Ebonyi")
	fmt.Println("Goodbye Ebonyi")
	fmt.Println("My age is", age, "and my name is", name)

	// Formatted string - Printf %_ = format specifier
	fmt.Printf("My age is %v and my name is %v \n", age, name)
	fmt.Printf("My age is %q and my name is %q \n", age, name)
	fmt.Printf("Age is of type %T \n", age)
	fmt.Printf("You score %0.1f points! \n", 20.57)

	// Sprintf(Save formatted strings)
	str := fmt.Sprintf("My age is %v and my name is %v \n", age, name)
	fmt.Println("The saved sring is", str)

}
