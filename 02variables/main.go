package main

import "fmt"

// variable def type
// var variable_list optional_data_type;

func main() {
	// int Variable
	var userID int = 10
	fmt.Println(userID)
	fmt.Printf("userID type is: %T \n", userID)

	// String Variable
	var useremail string = "abc@ibm.com"
	fmt.Println(useremail)
	fmt.Printf("useremail type is: %T \n", useremail)

	// String Variable
	var username string = "Meetanshi"
	fmt.Println(username)
	fmt.Printf("Username type is: %T \n", username)

	// float Variable
	var salary float32 = 2457000.56
	fmt.Println(salary)
	fmt.Printf("salary type is: %T \n", salary)

	// bool Variable
	var isIBMUser bool = true
	fmt.Println(isIBMUser)
	fmt.Printf("Username type is: %T \n", isIBMUser)

	// initialisation after dec
	var x float64
	x = 20.0
	fmt.Println(x)
	fmt.Printf("x is of type %T\n", x)

	// Declare and initialize variable d
	d := 42

	// Print the variable d and its type
	fmt.Printf("d: %T\n", d) // Use fmt.Printf for formatting

	var i, j, k int8
	i = 1
	j = 2
	k = -3
	fmt.Println(i, j, k)

	// Correct usage of fmt.Printf with format string
	n, err := fmt.Printf("i: %d, j: %d, k: %d\n", i, j, k)

	// Printing the return values of fmt.Printf
	fmt.Println("Bytes written:", n)
	fmt.Println("Error:", err)

}
