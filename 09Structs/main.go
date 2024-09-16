package main

import (
	"fmt"
)

// 1. Basic Struct Declaration
// Define a struct named Person with fields Name and Age.
type Person struct {
	Name string
	Age  int
}

// 2. Struct Initialization
// Initialize a struct using named fields.
func main() {
	// Initializing a Person struct with named fields
	person1 := Person{
		Name: "Alice",
		Age:  30,
	}
	fmt.Println("Person 1:", person1)

	// Initializing a Person struct with zero values and setting fields afterward
	var person2 Person
	person2.Name = "Bob"
	person2.Age = 25
	fmt.Println("Person 2:", person2)

	// 3. Struct with Embedded Fields
	// Define a struct with embedded fields (inheritance-like behavior).
	type Address struct {
		City    string
		Country string
	}

	type Contact struct {
		Person
		Address
		Phone string
	}

	// Initialize Contact struct
	contact := Contact{
		Person: Person{
			Name: "Charlie",
			Age:  35,
		},
		Address: Address{
			City:    "New York",
			Country: "USA",
		},
		Phone: "123-456-7890",
	}
	fmt.Println("Contact:", contact)

	// 4. Struct Methods
	// Define methods on the Person struct.
	func (p Person) Greet() string {
		return fmt.Sprintf("Hello, my name is %s and I am %d years old.", p.Name, p.Age)
	}

	func (p *Person) HaveBirthday() {
		p.Age++
	}

	// Using methods on the Person struct
	person3 := Person{Name: "Diana", Age: 40}
	fmt.Println(person3.Greet()) // Call method to greet
	person3.HaveBirthday()      // Call method to have a birthday
	fmt.Println("After birthday:", person3.Greet())

	// 5. Structs with Constructors
	// Define a constructor function to create a new Person struct.
	func NewPerson(name string, age int) Person {
		return Person{Name: name, Age: age}
	}

	// Create a Person instance using the constructor
	person4 := NewPerson("Eve", 28)
	fmt.Println("Person 4:", person4)

	// 6. Anonymous Structs
	// Define and use an anonymous struct.
	anonPerson := struct {
		Name string
		Age  int
	}{
		Name: "Frank",
		Age:  45,
	}
	fmt.Println("Anonymous Person:", anonPerson)

	// 7. Comparing Structs
	// Compare two structs for equality.
	person5 := Person{Name: "Grace", Age: 50}
	person6 := Person{Name: "Grace", Age: 50}

	if person5 == person6 {
		fmt.Println("Person 5 and Person 6 are equal.")
	} else {
		fmt.Println("Person 5 and Person 6 are not equal.")
	}

	// 8. Struct with JSON Tags
	// Define a struct with JSON tags for serialization.
	type Product struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Price float64 `json:"price"`
	}

	// Initialize and print Product struct
	product := Product{ID: 1, Name: "Laptop", Price: 999.99}
	fmt.Println("Product:", product)
}
