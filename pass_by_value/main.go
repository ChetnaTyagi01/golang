package main

import "fmt"

// in the below function x is the copy of the name variable and not the original name variable
func updateName(x string) {
	x = "ria" // so here we are just updating the copy and not the original value
	// return x // if we want to change the original value we can return the value from here
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99 // updates the original map
}
func main() {
	// group A types -  strings, ints, floats, boolean, arrays, structs
	name := "john"
	updateName(name)
	fmt.Println(name) // output - john
	// since go is a pass by value language
	// everytime we pass a value/variable into a function, go creates a copy of the variable

	// group B types - slices, maps, functions
	menu := map[string]float64{
		"pie":     106.09,
		"pudding": 215.02,
	}
	updateMenu(menu)
	fmt.Println(menu)
	// for group B types, go creates a copy of the pointer which points to the same value hence, the original value gets changed
}
