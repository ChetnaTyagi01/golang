package main

import "fmt"

func main() {
	// In go, we don't use classes. Instead, we create custom types using struct keyword
	// struct - basically a blueprint which describes a type of data
	mybill := newBill("Mario's Bill")
	fmt.Println(mybill.format())
}
