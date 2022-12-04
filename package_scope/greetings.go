package main // we need to add this since both files are a part of main package. If we don't add this we will not be able to share variables and functions between both the files

import "fmt"

// any of the variables/functions we declare in this file will be automatically accessible in any other file under package main
var points = []int{20, 10, 30, 35, 90}

func sayHello(n string) {
	fmt.Println("hello", n)
}

func showScore() {
	fmt.Println("your scoreOne:", scoreOne)
	// fmt.Println("your scoreTwo:", scoreTwo) // throws an error since any variable declared under main() can't be accesed here it should be under package main to be accesible here
}
