package main

import "fmt"

// using pointers as arguments to access the data within the underying functions
func updateName(x *string) {
	*x = "wedge"
}

func main() {
	name := "tifa"

	fmt.Println("memory address of name:", &name) // &name gets a pointer to the memory location of name variable
	m := &name                                    // m itself has its own memory address but is currently storing the memory address of name variable
	fmt.Println("memory address:", m)
	fmt.Println("value at memory address:", *m) // using * infront of a pointer gets the value that it points to

	fmt.Println("before:", name) // tifa
	updateName(m)
	fmt.Println("after:", name) // wedge
}
