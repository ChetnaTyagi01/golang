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
		items: map[string]float64{},
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
	// add tip
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "tip:", b.tip)
	// total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tip)
	return fs
}

// receiver functions with pointers as below
// update tip
// when we call a function, when we take an argument into a function it creates copy
// but also on receiver functions we also take a copy of the bill rather than the actual bill - mybill := newBill("Mario's Bill")
// and we copy the bill received in main.go and pass that bill in the receiver fuction - (b *bill)

func (b *bill) updateTip(tip float64) { // passing a pointer to the bill
	b.tip = tip // and here we are updating the original value of the bill
}

// add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// whenever we are calling a method where we are updating the value we pass in the pointer
