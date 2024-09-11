package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// creates a new Reader instance that reads from the provided io.Reader (in this case, os.Stdin).
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome user")
	fmt.Println("Enter your name: ")
	name, err := reader.ReadString('\n')

	fmt.Printf("Type of this input is %T\n", name)

	if err != nil {
		fmt.Println("Error reading name:", err)
		return
	}

	fmt.Printf("Please rate our tutorials out of 5 Stars %s", name)

	// Read the input from the user
	input, err := reader.ReadString('\n')

	// print err if any
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Trim the newline character from the input
	// reader.ReadString('\n') includes the newline character (\n) at the end of the input.
	//  to remove the trailing newline, you can use the strings.TrimSpace function:
	input = strings.TrimSpace(input)

	// Convert input to an integer
	rating, err := strconv.Atoi(input)

	// if not valid rating
	if err != nil {
		fmt.Println("Invalid rating. Please enter a number between 1 and 5.")
		return
	}

	fmt.Printf("Type of this input is %T\n", rating)

	// Check the rating and provide feedback
	if rating > 3 {
		fmt.Println("We are happy you liked it!")
	} else if rating == 3 {
		fmt.Println("We will work to make it better next time.")
	} else if rating < 3 {
		fmt.Println("We regret that you didn't enjoy it. We will do our best next time.")
	} else {
		fmt.Println("Invalid rating. Please enter a number between 1 and 5.")
	}

	// Print the input value directly
	fmt.Printf("Thanks for rating us %s starts\n", input)
}
