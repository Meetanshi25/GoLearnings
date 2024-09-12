package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Initialize and assign values to an integer array
	intArray := [5]int{1, 2, 3, 4, 5}
	// to directly print complete arr
	fmt.Println("Integer Array:", intArray)

	fmt.Println("Integer Array:")
	for i := 0; i < len(intArray); i++ {
		fmt.Println(intArray[i])
	}

	// Initialize and assign values to a string array with one value missing
	var stringArray [5]string
	stringArray[0] = "Hello"
	stringArray[1] = "World"
	stringArray[2] = "Go"
	stringArray[4] = "Programming"
	fmt.Println("\nString Array (with one value missing):")
	for i := 0; i < len(stringArray); i++ {
		if stringArray[i] == "" {
			fmt.Println("Missing value")
		} else {
			fmt.Println(stringArray[i])
		}
	}

	// Declare another array and take input from the user
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\nEnter the size of the new integer array: ")
	sizeStr, _ := reader.ReadString('\n')
	// trim /n from above input
	sizeStr = strings.TrimSpace(sizeStr)
	// convert string format input to int
	size, err := strconv.Atoi(sizeStr)
	if err != nil || size <= 0 {
		fmt.Println("Invalid size. Please enter a positive integer.")
		return
	}

	// Declare a new integer array with user-defined size
	newArray := make([]int, size)
	fmt.Println("Enter the elements of the array:")

	for i := 0; i < size; i++ {
		fmt.Printf("Element %d: ", i+1)
		elementStr, _ := reader.ReadString('\n')
		elementStr = strings.TrimSpace(elementStr)
		element, err := strconv.Atoi(elementStr)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer.")
			i-- // Retry the current iteration
			continue
		}
		newArray[i] = element
	}

	// Print the user-defined array
	fmt.Println("\nUser-defined Array:")
	for i, value := range newArray {
		fmt.Printf("Element %d: %d\n", i+1, value)
	}
}
