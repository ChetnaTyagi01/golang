package main

import "fmt"

func main() {
	// Print
	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n")

	// Println
	fmt.Println("hello chetna!")
	fmt.Println("bye chetna!")

	age := 35
	name := "shaun"
	fmt.Println("my name is", name, "and my age is", age)

	// Printf - Formatted strings -  a way to create strings with variables inside
	// %_ - format specifier, here _ can be many different letters
	// %v - default format for variables
	fmt.Printf("my name is %v and my age is %v \n", name, age) // order matters here

	// %q - places quotes around variables when we output them, shoud be used on strings
	fmt.Printf("my name is %q and my age is %q \n", name, age) // will not work on age as it is an integer

	// %T - gets the type of the variable
	fmt.Printf("age is of type %T \n", age)

	// %f - for float variables
	fmt.Printf("you scored %f points \n", 225.45)    // by default %f takes upto 6 decimal places
	fmt.Printf("you scored %0.1f points \n", 225.45) // 0.1f rounds up to 1 decimal place

	// Sprintf - save formatted strings
	// returns a formatted string and we can store that in a variable
	var str = fmt.Sprintf("my name is %v and my age is %v \n", name, age)
	fmt.Println("the saved string is:", str)

	// https://golang.org/pkg/fmt
}
