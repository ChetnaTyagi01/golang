package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value) // calling a function for each element inside the array
	}
}

// return value from a function
func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	sayGreeting("julia")
	sayBye("tia")

	cycleNames([]string{"cloud", "tifa", "barret"}, sayGreeting)
	cycleNames([]string{"cloud", "tifa", "barret"}, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(15)
	fmt.Println(a1, a2)
	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f", a1, a2)
}
