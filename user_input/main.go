package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// prompt - contains the string i.e the question we ask the user
// r - pointer to the reader object so that we dont need to create the reader object again n again
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err

}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Create a new bill name: ")
	// name, _ := reader.ReadString('\n') // reader will read the input when the user presses enter and control transfers to a new line
	// name = strings.TrimSpace(name)

	name, _ := getInput("Create a new bill name: ", reader)
	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
	fmt.Println(opt)
}

func main() {
	mybill := createBill()
	promptOptions(mybill)
	// fmt.Println(mybill)
}
