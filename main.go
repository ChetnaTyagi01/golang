package main

import "fmt"

func main() {
	// declaring array - fixed length - length of arrays cannot change
	// var ages [4]int = [4]int{20, 25, 30, 35}
	var ages = [4]int{20, 25, 30, 35}
	fmt.Println("array:", ages)
	fmt.Println("array length:", len(ages))
	var names = [5]string{"max", "sam", "ron", "mario", "ash"}
	fmt.Println("array before:", names)
	names[1] = "tia"
	fmt.Println("array after:", names)
	fmt.Println("array length:", len(names))

	// slices - arrays where we can change the length, we can add/delete items
	// slices use arrays under the hood
	var scores = []int{100, 50, 60} // not specified the length
	scores[2] = 25
	// we can also append items to the above slice
	fmt.Println("After append:", append(scores, 85)) // append returns a new slice, doesnt autoamtically update the existing slice
	scores = append(scores, 85)                      // so we need to assign to the older slice
	fmt.Println(scores, len(scores))

	// slice ranges
	// range is a way to get a range of values from an existing array or slice and store them in a new slice
	rangeOne := names[1:3]             // inclusive of the first number but not the second
	fmt.Println("rangeOne:", rangeOne) // doesnt include value at index 3

	rangeTwo := names[2:] // this goes to the end of the array or slice
	fmt.Println("rangeTwo", rangeTwo)

	rangeThree := names[:3] // goes from the start uptill index 3 but doesnt include the value at index 3
	fmt.Println("rangeThree:", rangeThree)

	// we can append the ranges as well
	rangeOne = append(rangeOne, "david")
	fmt.Print("rangeOne After append:", rangeOne)

}
