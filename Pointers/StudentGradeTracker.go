package Pointers

import (
	"fmt"
	"strings"
)

// We want to track student grades in different ways:
// - Track all grades
// - Track only passing grades (>= 60)
// - Track only failing grades (< 60)
// The key is that when we modify a grade in one list,
// it should update everywhere that grade appears.

func StudentGradeTracker() {

	// Task 1: Create three slices of grade pointers:
	// - allGrades: contains pointers to all grades
	// - passingGrades: contains pointers to grades >= 60
	// - failingGrades: contains pointers to grades < 60

	// Initial grades
	grades := []int{45, 68, 82, 55, 75, 90, 50}
	// Your code here to create and populate the pointer slices
	var allGrades []*int
	var passingGrades []*int
	var failingGrades []*int
	// Task 2: Create a function that:
	// - Takes a grade value and its index
	// - Updates the grade at that index
	// - The change should be reflected in all relevant slices
	// Populate the pointer slices
	for i := range grades {
		allGrades = append(allGrades, &grades[i])
		if grades[i] >= 60 {
			passingGrades = append(passingGrades, &grades[i])
		} else {
			failingGrades = append(failingGrades, &grades[i])
		}
	}
	fmt.Println("Initial State:")
	printGradeLists(allGrades, passingGrades, failingGrades)

	// Task 3: Create a function that prints all three lists
	// showing which grades are tracked in each category

	// Test your implementation:
	// 1. Print initial state
	// 2. Update some grades
	// 3. Print final state to verify changes are reflected everywhere

	// Update a failing grade to passing
	updateGrade(allGrades, 0, 75) // Update first grade from 45 to 75
	fmt.Println("\nAfter updating first grade to 75:")
	printGradeLists(allGrades, passingGrades, failingGrades)
}

func updateGrade(grades []*int, updateIndex int, updateValue int) {
	// Your code here
	*grades[updateIndex] = updateValue
}

func printGradeLists(allGrades []*int, passingGrades []*int, failingGrades []*int) {
	// Your code here
	fmt.Println("allGrades", allGrades)
	// Print each list with clear labels
	// Show both the index and value for each grade
	// Consider adding visual indicators to make the output easy to read
	formatList := func(grades []*int) string {
		var sb strings.Builder
		for i, grade := range grades {
			sb.WriteString(fmt.Sprintf("%d: %d", i, *grade)) // Include index and value
			if i < len(grades)-1 {
				sb.WriteString(", ") // Add comma except for the last item
			}
		}
		return sb.String()
	}
	// Print each list with labels
	fmt.Println("allGrades:")
	fmt.Println(formatList(allGrades))

	fmt.Println("passingGrades:")
	fmt.Println(formatList(passingGrades))

	fmt.Println("failingGrades:")
	fmt.Println(formatList(failingGrades))
}
