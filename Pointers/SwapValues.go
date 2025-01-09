package Pointers

import "fmt"

func SwapValues() {
	// Task: Write a function that swaps two integer values using pointers
	// Bonus: Make it work with any data type using interfaces

	x := 40
	y := 20

	fmt.Printf("Before swap: x = %d, y = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("After swap: x = %d, y = %d\n", x, y)
}

func swap(x *int, y *int) {
	// Your code here
	xPointer := x
	yPointer := y

	tempPointer := *xPointer

	*xPointer = *yPointer
	*yPointer = tempPointer
}
