package LinkedList

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/merge-two-sorted-lists

type Number struct {
	value    uint
	next     *Number // Pointer to next number
	previous *Number // Pointer to previous number
}

type NumberList struct {
	currentNumber *Number
}

// NewNumberList creates a new number list
func NewNumberList() *NumberList {
	return &NumberList{
		currentNumber: nil,
	}
}

func (n *NumberList) AddNumberToList(number uint) {
	newNumber := &Number{
		value: number,
	}
	// If number is empty, make this the first number
	if n.currentNumber == nil {
		n.currentNumber = newNumber
		return
	}

	// Find the last number in the list
	current := n.currentNumber
	for current.next != nil {
		current = current.next
	}

	// link the new number
	current.next = newNumber
	newNumber.previous = current
}

func (n *NumberList) Length() int {
	count := 0
	current := n.currentNumber
	for current != nil {
		count++
		current = current.next
	}
	return count
}
func MergeList(list1 *NumberList, list2 *NumberList) *NumberList {
	// Create a new list for the merged result
	mergedList := NewNumberList()

	// Get the first nodes of both lists
	current1 := list1.currentNumber
	current2 := list2.currentNumber

	// While both lists have nodes to process
	for current1 != nil && current2 != nil {
		if current1.value <= current2.value {
			// add the smaller value from list1
			mergedList.AddNumberToList(current1.value)
			current1 = current1.next
		} else {
			mergedList.AddNumberToList(current2.value)
			current2 = current2.next
		}
	}
	// Add remaining elements from list1, if any
	for current1 != nil {
		mergedList.AddNumberToList(current1.value)
		current1 = current1.next
	}
	for current2 != nil {
		mergedList.AddNumberToList(current2.value)
		current2 = current2.next
	}
	return mergedList

}

func (n *NumberList) ViewNumberList() string {
	// Handle empty playlist case
	if n.currentNumber == nil {
		return "NumberList is empty"
	}

	// Create a string builder for efficient string concatenation
	var result strings.Builder
	result.WriteString("=== Your Number List ===\n")

	// First, find the first number in the number list
	firstNumber := n.currentNumber
	for firstNumber.previous != nil {
		firstNumber = firstNumber.previous
	}

	// Traverse through all songs starting from the first one
	current := firstNumber
	position := 1
	for current != nil {
		// Add currently playing indicator if this is the current song
		nowPlaying := " "
		if current == n.currentNumber {
			nowPlaying = "▶ " // Unicode play symbol
		}

		// Format and write this song's information
		numberInfo := fmt.Sprintf("%s%d. %d \n",
			nowPlaying,
			position,
			current.value)
		result.WriteString(numberInfo)

		current = current.next
		position++
	}

	result.WriteString("==================")
	return result.String()
}
func MergeTwoSortedLists() {
	// create 2 lists
	numberList1 := NewNumberList()
	numberList2 := NewNumberList()
	// assign value to each list
	numberList1.AddNumberToList(1)
	numberList1.AddNumberToList(2)
	numberList1.AddNumberToList(4)

	numberList2.AddNumberToList(1)
	numberList2.AddNumberToList(3)
	numberList2.AddNumberToList(4)

	// View 2 list after added
	fmt.Println(numberList1.ViewNumberList())
	fmt.Println(numberList2.ViewNumberList())
	// merge 2 lists

	mergedList := MergeList(numberList1, numberList2)

	fmt.Println("Merged List:")
	fmt.Println(mergedList.ViewNumberList())
	// sort list of node

}
