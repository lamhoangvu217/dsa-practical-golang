package Pointers

import "fmt"

func PointerToSlices() {
	numbers := []int{1, -2, 3, -4, 5}
	fmt.Printf("Original slice: %v\n", numbers)

	// Call your functions here
	doubleElements(&numbers)

	fmt.Println("new numbers after double: ", numbers)

	removeNegatives(&numbers)
	fmt.Println("new numbers after remove negative: ", numbers)

}

func doubleElements(numbers *[]int) {
	for i := range *numbers {
		(*numbers)[i] *= 2
	}
}

func removeNegatives(numbers *[]int) {
	// Remove all negative numbers from a slice
	// Modify the original slice (not create a new one)
	// Change the length of the slice (since we're removing elements)
	original := *numbers
	var tempArr []int
	for _, item := range original {
		if item >= 0 {
			tempArr = append(tempArr, item)
		}
	}
	*numbers = tempArr
}
