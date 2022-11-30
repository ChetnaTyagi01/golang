package main

import "fmt"

func main() {
	fmt.Println("Hello world!")

	// strings
	var nameOne string = "chetna"            // explicitly giving the type
	var nameTwo = "tyagi"                    // Go is automaticalling inferring the type
	var nameThree string                     // declaring for future use
	fmt.Println(nameOne, nameTwo, nameThree) // empty string value for nameThree

	// shorthand version
	nameFour := "chetnatyagi.ct" // this version can't be used outside of a function
	fmt.Println(nameFour)

	// ints
	var ageOne int = 40
	var ageTwo = 55
	ageThree := 60
	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	var numOne int8 = 25
	var numTwo int16 = 256
	var numThree uint = 300 // unsigned int
	fmt.Println(numOne, numTwo, numThree)

	// floats - mandatory to specify bits
	var scoreOne float32 = 25.8
	var scoreTwo float64 = 888888777876543245789.9
	scoreThree := 9876546789.9
	fmt.Println(scoreOne, scoreTwo, scoreThree)

	// https://go.dev/ref/spec#Numeric_types
}
