package main

import "fmt"

func main() {
	// Basic pointer example
	var a int = 10
	fmt.Println("Basic Pointer Example:")
	fmt.Println("Value of a:", a)                   // Print the value of a
	fmt.Println("Address of a:", &a)                // Print the address of a
	var p *int = &a                                 // Declare a pointer variable p and assign it the address of a
	fmt.Println("Value stored in pointer p:", *p)   // Dereference p to get the value at the address stored in p (which is the value of a)
	fmt.Println("Address stored in pointer p:", p)  // Print the address stored in p (which is the address of a)
	fmt.Println("Address stored in pointer p:", &p) // Print the address of p

	// updating value of a using p
	*p = *p + 2
	fmt.Println("Updated value stored in pointer p:", *p) // changes value stored in a
	fmt.Println("Updated value stored in pointer p:", a)

	// Pointer to a function example
	fmt.Println("\nPointer to a Function Example:")
	result := add(5, 7)                         // Call the add function
	fmt.Println("Result of add(5, 7):", result) // Print the result of the add function

	addFunctionPointer := add                                     // Create a function pointer and assign it the address of the add function
	result = addFunctionPointer(10, 20)                           // Call the function through the pointer
	fmt.Println("Result of add(10, 20) through pointer:", result) // Print the result

	// Pointer to a slice example
	fmt.Println("\nPointer to a Slice Example:")
	slice := []int{1, 2, 3}
	fmt.Println("Original slice:", slice)
	modifySlice(&slice)                   // Pass the pointer to the slice
	fmt.Println("Modified slice:", slice) // Print the modified slice

	// Pointer to a struct example
	fmt.Println("\nPointer to a Struct Example:")
	person := Person{Name: "John", Age: 30}
	fmt.Println("Original Person:", person)
	updatePerson(&person)                  // Pass the pointer to the struct
	fmt.Println("Updated Person:", person) // Print the updated person

	// Pointer arithmetic example (not allowed in Go)
	// Uncommenting the following line will cause a compile-time error
	// fmt.Println("Pointer arithmetic:", p + 1) // Go does not support pointer arithmetic
}

// Function to add two integers
func add(a int, b int) int {
	return a + b
}

// Function that modifies a slice using a pointer
func modifySlice(slice *[]int) {
	*slice = append(*slice, 4, 5) // Modify the slice by dereferencing the pointer
}

// Struct definition
type Person struct {
	Name string
	Age  int
}

// Function that updates a struct using a pointer
func updatePerson(p *Person) {
	p.Name = "Jane" // Modify the struct by dereferencing the pointer
	p.Age = 25
}
