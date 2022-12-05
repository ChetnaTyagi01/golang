package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64) // parsing string to float64
		if err != nil {
			fmt.Println("The price must be a number!")
			promptOptions(b) // to back to the start so that user can chose the option again and provide a valid input
		}
		b.addItem(name, p) // add item to the bill
		fmt.Println("item added - ", name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number!")
			promptOptions(b)
		}
		b.updateTip(t) // add tip to the bill
		fmt.Println("tip added - ", tip)
		promptOptions(b)
	case "s":
		b.saveBill()
		fmt.Println("You saved the file", b.name)
	default:
		fmt.Println("Wrong choice!")
		promptOptions(b) // to restart the function - as we want the user to choose only a valid option
	}
}

func main() {
	mybill := createBill()
	promptOptions(mybill)
	// fmt.Println(mybill)
}
