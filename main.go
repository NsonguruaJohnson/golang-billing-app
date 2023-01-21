package main

import "fmt"

func main() {
	myBill := newBill("Nsongurua's bill")

	myBill.addItem("Sharwama", 2.55)
	myBill.addItem("Pizza", 6.42)
	myBill.addItem("Icecream", 5.55)

	myBill.updateTip(10)

	fmt.Println(myBill.format())

}
