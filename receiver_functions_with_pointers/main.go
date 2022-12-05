package main

import "fmt"

func main() {
	// In go, we don't use classes. Instead, we create custom types using struct keyword
	// struct - basically a blueprint which describes a type of data
	mybill := newBill("Mario's Bill")
	mybill.addItem("tomato soup", 4.50)
	mybill.addItem("cake", 5.55)
	mybill.addItem("veg pie", 6.50)
	mybill.addItem("coffee", 3.50)
	mybill.updateTip(10)
	fmt.Println(mybill.format())
}
