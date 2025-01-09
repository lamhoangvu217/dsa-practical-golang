package Pointers

import "fmt"

func ModifyValue() {
	name := "Go"

	fmt.Printf("Before modification: %s\n", name)

	// Your task: Write a function that takes a pointer to a string
	// and adds "Programming" after it
	// So "Go" becomes "Go Programming"
	addProgramming(&name)

	fmt.Printf("After modification: %s\n", name)
}

func addProgramming(text *string) {
	// Write your code here
	// Hint: To modify what a pointer points to, use *text
	*text = *text + " " + "Programming"
}
