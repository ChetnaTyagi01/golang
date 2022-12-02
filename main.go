package main

import "fmt"

func main() {
	// maps
	// made up key-value pairs. keys can be of multiple different types - strings, integers, float etc
	// all of the keys inside a map must have same type
	// all of the values inside a map must have same type too
	// syntax - map[key-type]value-type{}

	menu := map[string]float64{
		"soup":    4.99,
		"pie":     7.99,
		"salad":   6.99,
		"pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping through a map
	for key, value := range menu {
		fmt.Println(key, "-", value)
	}

	// ints as key type
	phonebook := map[int]string{
		76543567: "marie",
		87654343: "john",
		32456435: "tia",
		99876543: "robert",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[87654343])

	// updating an item inside the map
	phonebook[87654343] = "julia"
	fmt.Println("updated phonebook:", phonebook)
}
