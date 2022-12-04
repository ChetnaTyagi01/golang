package main

import "fmt"

func main() {
	x := 0
	// for works as while loop here
	for x < 5 {
		fmt.Println("value of x:", x)
		x++
	}

	// traditional for loop
	for i := 0; i < 5; i++ {
		fmt.Println("value of i:", i)
	}

	// looping through a slice of strings/numbers
	names := []string{"maya", "riya", "tia", "ria"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// alternate way - using for-in loop
	// it cycles through slice - names, for each going round this slice we get the index and value at that index
	for index, value := range names {
		fmt.Printf("the value at index %v is %v \n", index, value)
	}

	// in the above scenario - if we don't want to use the index and use just the value, we can replace it with an underscore
	// likewise is the case with value
	for _, value := range names {
		fmt.Printf("the value is %v \n", value)
		value = "new string" // adding a new value here doesn't modify the actual names slice because it is just a copy of original slice
	}

	fmt.Println("original slice:", names)
}
