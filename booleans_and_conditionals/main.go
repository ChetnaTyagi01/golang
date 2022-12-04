package main

import "fmt"

func main() {
	age := 45
	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less than 45")
	}

	// nesting if statement inside a loop
	// using continue and break statements
	names := []string{"zayn", "jack", "lisa", "rachel", "joey"}
	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("breaking at pos", index)
			break
		}

		fmt.Printf("value at pos %v is %v \n", index, value)
	}
}
