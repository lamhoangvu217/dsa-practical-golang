package Pointers

import "fmt"

func PointerArrays() {
	// Task: Create an array of 3 integers
	// Create an array of 3 pointers to integers
	// Make each pointer in the second array point to
	// the corresponding element in the first array
	// Modify values through pointers and verify changes

	// Your code here
	// Creating the base integer array
	arr1 := [3]int{1, 2, 3}

	// Declaring an array of pointers to integers
	var pointers [3]*int

	// Establishing pointer connections
	for i := range arr1 {
		pointers[i] = &arr1[i]
	}
	fmt.Printf("Original array: %v\n", arr1)

	// Modifying values through pointers
	for i := range pointers {
		*pointers[i] += 3
	}
	fmt.Printf("Modified through pointers: %v\n", arr1)
}
