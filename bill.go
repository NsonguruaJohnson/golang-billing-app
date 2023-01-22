package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// Make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// format the bill (any character can be used in place of c)
func (c *bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// list items
	for k, v := range c.items {
		fs += fmt.Sprintf("%-25v ... $%v \n", k+":", v)
		total += v
	}

	// add tip
	fs += fmt.Sprintf("%-25v ... $%v \n", "tip:", c.tip)

	// total
	fs += fmt.Sprintf("%-25v ... $%0.2f", "total:", total+c.tip)

	return fs
}

// update tip
func (c *bill) updateTip(tip float64) {
	(*c).tip = tip
	// Telling Go to dereference the line above
	// Go can also auto matically dereference it
}

// add an item to the bill
func (c *bill) addItem(name string, price float64) {
	c.items[name] = price
	// Go automatically deferences the line above
}
