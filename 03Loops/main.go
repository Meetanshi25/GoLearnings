package main

import "fmt"

func main() {

	// for loop
	var b int = 15
	var a int
	numbers := [6]int{1, 2, 3, 5}

	// For loop execution
	for a := 0; a < 10; a++ {
		fmt.Printf("value of a: %d\n", a)
	}
	for a < b {
		a++
		fmt.Printf("value of a: %d\n", a)
	}
	for i, x := range numbers {
		fmt.Printf("value of x = %d at %d\n", x, i)
	}
	// Expected output
	// value of x = 1 at 0
	// value of x = 2 at 1
	// value of x = 3 at 2
	// value of x = 5 at 3
	// value of x = 0 at 4
	// value of x = 0 at 5

	// Nested loops
	var itr, j int

	for itr = 2; itr < 100; itr++ {
		for j = 2; j <= (itr / j); j++ {
			if itr%j == 0 {
				break // if factor found, not prime
			}
		}
		if j > (itr / j) {
			fmt.Printf("%d is prime\n", itr)
		}
	}

	// goto statement
	// A goto statement in Go programming language provides
	// an unconditional jump from the goto to a labeled statement in the same function.

	// syntax
	// goto label;
	// ..
	// .
	// label: statement;

	/* local variable definition */
	var atr int = 10

	/* do loop execution */
LOOP:
	for atr < 20 {
		if atr == 15 {
			/* skip the iteration */
			atr = atr + 1
			goto LOOP
		}
		fmt.Printf("value of atr: %d\n", atr)
		atr++
	}

}
