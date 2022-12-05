package main

import "fmt"

// custom type - blueprint for any bill object
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:   0,
	}
	return b
}

// creating a receiver function associated with bill object - to format the bill
// by adding the first parenthesis we receive a bill into the below function OR we are limiting this function only to bill object
func (b bill) format() string {
	fs := "Bill Breakdown: \n"
	var total float64 = 0
	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v) // %-25v makes this character value to be output 25 characters long
		total += v
	}
	// total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)
	return fs
}
