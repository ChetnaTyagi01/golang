package main

import "fmt"

var scoreOne = 99.2

// main() function only appears once in the program and is an entry point to the application
func main() {
	// var scoreTwo = 97.5 // can't be accessed in any other file
	sayHello("chetna") // using other file's function and variables here
	for _, v := range points {
		fmt.Println(v)
	}
	showScore()
}
