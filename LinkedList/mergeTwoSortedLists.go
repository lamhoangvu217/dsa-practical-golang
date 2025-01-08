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
	temp := &NumberList{
		currentNumber: &Number{}, // Create a dummy Number node
	}
	prev := temp.currentNumber
	current1, current2 := list1.currentNumber, list2.currentNumber
	for current1 != nil && current2 != nil {
		if current1.value < current2.value {
			prev.next = current1
			current1 = current1.next
		} else {
			prev.next = current2
			current2 = current2.next
		}
		prev = prev.next
	}
	if current1 != nil {
		prev.next = current1
	} else {
		prev.next = current2
	}
	temp.currentNumber = temp.currentNumber.next
	return temp
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
	numberList1.AddNumberToList(4)
	numberList1.AddNumberToList(5)
	numberList1.AddNumberToList(8)

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
