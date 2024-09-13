package main

import (
	"fmt"
)

// 1. Basic Function
// A simple function that prints a welcome message.
func greet() {
	fmt.Println("Hello, World!")
}

// 2. Function with Parameters
// A function that adds two integers and returns the result.
func add(a int, b int) int {
	return a + b
}

// 3. Function with Multiple Return Values
// A function that swaps two integers and returns them in reversed order.
func swap(a, b int) (int, int) {
	return b, a
}

// 4. Named Return Values
// A function that divides two integers and returns the result and an error if division by zero is attempted.
func divide(a, b int) (result int, err error) {
	if b == 0 {
		err = fmt.Errorf("cannot divide by zero")
		return
	}
	result = a / b
	return
}

// 5. Variadic Function
// A function that sums a variable number of integers.
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// 6. Anonymous Function
// An anonymous function is defined and executed immediately.
func main() {
	// 1. Basic Function
	greet() // Call the basic function

	// 2. Function with Parameters
	result := add(10, 20)
	fmt.Println("Sum of 10 and 20:", result)

	// 3. Function with Multiple Return Values
	x, y := swap(1, 2)
	fmt.Println("Swapped values:", x, y)

	// 4. Named Return Values
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division result:", result)
	}

	// 5. Variadic Function
	total := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum of multiple values:", total)

	// 6. Anonymous Function
	func(message string) {
		fmt.Println("Anonymous function says:", message)
	}("Hello, Go!") // Pass argument immediately

	// Assign an anonymous function to a variable and call it
	greetFunc := func(name string) string {
		return "Hello, " + name
	}
	fmt.Println(greetFunc("Alice"))

	// 7. Function as a Return Value
	// A function that returns another function to multiply by a given factor.
	multiplier := func(factor int) func(int) int {
		return func(x int) int {
			return x * factor
		}
	}

	timesTwo := multiplier(2)
	fmt.Println("5 multiplied by 2 is:", timesTwo(5))

	timesThree := multiplier(3)
	fmt.Println("5 multiplied by 3 is:", timesThree(5))

	// 8. Deferred Function Calls
	// Demonstrate deferred execution of a function.
	defer fmt.Println("Deferred: This will be executed last")
	fmt.Println("Processing...")
}
