package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "hello there friends!"

	// strings package
	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " ")) // returns a slice of 3 elements

	// original value is unchanged
	fmt.Println("original string value:", greeting)

	// sort package
	ages := []int{45, 20, 35, 30, 42, 65, 55, 70}
	sort.Ints(ages)
	fmt.Println(ages) // original value of ages variable gets changed

	index := sort.SearchInts(ages, 42) // will search the sorted ages
	fmt.Println(index)

	names := []string{"jack", "tia", "john", "kate", "marie"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "tia"))

	// https://pkg.go.dev/std#stdlib
}
