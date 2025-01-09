package Pointers

import (
	"fmt"
	"strings"
)

// Book represents a book in a library
type Book struct {
	Title  string
	Author string
	Pages  int
	// IsCheckedOut tells us if someone has borrowed the book
	IsCheckedOut bool
}

// Library keeps track of books and some statistics
type Library struct {
	Books           []Book
	TotalBooks      int
	CheckedOutBooks int
	// We'll use a pointer to track the most popular book
	MostPopularBook *Book
}

func LibraryExercise() {
	// Create a new library
	library := Library{
		Books: []Book{
			{Title: "The Go Programming Language", Author: "Alan Donovan", Pages: 380, IsCheckedOut: false},
			{Title: "Pride and Prejudice", Author: "Jane Austen", Pages: 432, IsCheckedOut: false},
			{Title: "The Hobbit", Author: "J.R.R. Tolkien", Pages: 310, IsCheckedOut: false},
		},
		TotalBooks:      3,
		CheckedOutBooks: 0,
	}

	fmt.Println("=== Welcome to the Library System ===")

	// Let's perform some operations:

	// 1. Check out a book, borrow a book, remove a book from library
	checkOutBook(&library, 0)
	checkOutBook(&library, 1)
	fmt.Printf("\nAfter checking out first book:\n")
	printLibraryStatus(&library)

	// 2. Return a book, add a book into library
	returnBook(&library, 0)
	fmt.Printf("\nAfter returning first book:\n")
	printLibraryStatus(&library)

	// 3. Update the most popular book
	//updateMostPopular(&library, 0)
	//fmt.Printf("\nMost popular book is now: %s\n", library.MostPopularBook.Title)
}

// checkOutBook should:
// - Set the book's IsCheckedOut status to true
// - Increment the library's CheckedOutBooks counter
// - Return false if the book is already checked out
func checkOutBook(lib *Library, bookIndex int) bool {
	// Your code here
	// Remember to check if the book is already checked out!

	// get access to the specific book
	if bookIndex >= len((*lib).Books) || bookIndex < 0 { // dereference (*lib)
		return false // Book doesn't exist
	}
	// Get a reference to the specific book we want to check out
	bookToCheck := &(*lib).Books[bookIndex]

	if bookToCheck.IsCheckedOut {
		fmt.Println("This book has already checked out")
		return false
	}
	bookToCheck.IsCheckedOut = true
	(*lib).CheckedOutBooks += 1
	return true
}

// returnBook should:
// - Set the book's IsCheckedOut status to false
// - Decrement the library's CheckedOutBooks counter
// - Return false if the book wasn't checked out
func returnBook(lib *Library, bookIndex int) bool {
	// Your code here
	// Remember to check if the book was actually checked out!
	fmt.Println("processing")
	return false
}

// updateMostPopular should:
// - Make the book at bookIndex the most popular book
// - Use proper pointer assignment
func updateMostPopular(lib *Library, bookIndex int) {
	// Your code here
	// Hint: think about how to make MostPopularBook point to a book in the slice
	fmt.Println("processing")
}

// printLibraryStatus prints the current status of the library
func printLibraryStatus(lib *Library) {
	// Your code here
	// Print total books, checked out books, and the most popular book if it exists
	var sb strings.Builder
	for i, lib := range (*lib).Books {
		sb.WriteString(fmt.Sprintf("%d: %s", i, lib.Title))
	}
}
