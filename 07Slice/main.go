package main

import (
	"fmt"
)

func main() {
	// Example 1: Creating a slice using a slice literal
	fmt.Println("Example 1: Slice Literal")
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", slice1)

	// Example 2: Slicing an existing slice
	fmt.Println("\nExample 2: Slicing an Existing Slice")
	slice2 := slice1[1:4] // Slice from index 1 to 3 (4 is excluded)
	fmt.Println("Sliced slice (slice1[1:4]):", slice2)

	// Example 3: Modifying a slice
	fmt.Println("\nExample 3: Modifying a Slice")
	slice2[0] = 10 // Modify the first element of slice2
	fmt.Println("Modified slice2:", slice2)
	// When you create slice2 with slice2 := slice1[1:4],
	// slice2 starts from index 1 of slice1 and goes up to,
	// but does not include, index 4. Thus, slice2 is a slice
	// of the same array, starting at index 1 and ending at index 3.
	fmt.Println("Original slice1 after modification:", slice1)

	// Example 4: Creating a slice with make
	fmt.Println("\nExample 4: Creating a Slice with make")
	slice3 := make([]int, 5) // Create a slice with length 5, all elements are zero by default
	fmt.Println("Slice created with make:", slice3)

	// Example 5: Appending to a slice
	fmt.Println("\nExample 5: Appending to a Slice")
	slice3 = append(slice3, 6, 7, 8) // Append elements to the end of slice3
	fmt.Println("Slice3 after appending:", slice3)

	slice3 = append(slice3[1:]) // removes all ele apart from [1,end)
	fmt.Println("Slice3 after 2nd appending:", slice3)

	// Example 6: Copying a slice
	fmt.Println("\nExample 6: Copying a Slice")
	slice4 := make([]int, len(slice3)) // Create a new slice with the same length as slice3
	copy(slice4, slice3)               // Copy elements from slice3 to slice4
	fmt.Println("Copied slice4:", slice4)

	// Example 7: Using slices with multidimensional slices
	fmt.Println("\nExample 7: Multidimensional Slices")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("2D slice (matrix):")
	for _, row := range matrix {
		fmt.Println(row)
	}

	// Example 8: Slicing a slice with a default capacity
	fmt.Println("\nExample 8: Slicing a Slice with Default Capacity")
	slice5 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Original slice5:", slice5)
	slice6 := slice5[:5] // Slice up to index 4 (5 is excluded)
	fmt.Println("Sliced slice5[:5]:", slice6)

	// Example 9: Understanding slice capacity
	fmt.Println("\nExample 9: Slice Capacity")
	slice7 := []int{1, 2, 3}
	fmt.Println("Original slice7:", slice7)
	fmt.Println("Capacity of slice7:", cap(slice7))

	slice8 := slice7[:cap(slice7)] // Extend the slice to its full capacity
	fmt.Println("Extended slice8:", slice8)
	fmt.Println("Capacity of slice8:", cap(slice8))

	// Example 10 removing some elements
	testSlice := [7]int{1, 2, 3, 4, 5, 6, 7}
	index := 3
	slice9 := append(testSlice[:index], testSlice[index+1:]...)
	fmt.Println("New slice is ", slice9) // o/p=1,2,3,5,6,7
}
