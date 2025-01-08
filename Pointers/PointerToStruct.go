package Pointers

import "fmt"

type Person struct {
	Name string
	Age  int
}

func PointerToStruct() {
	// Task: Create a Person struct and a pointer to it
	// Initialize the struct through the pointer
	// Create a function modifyAge that takes a pointer to Person
	// and increases the age by 1
	person1 := Person{Name: "Vu Hoang Lam 1", Age: 24}
	person2 := Person{Name: "Vu Hoang Lam 2", Age: 25}
	person1Pointer := &person1
	person2Pointer := &person2

	fmt.Println("person pointer", person1Pointer, person2Pointer)
	// Your code here

	modifiedPerson1 := modifyAge(&person1)
	*person2Pointer = modifyAge(&person2)
	*person2Pointer = modifyName(&person2)

	fmt.Println("Modified Person 1", modifiedPerson1, person1)
	fmt.Println("Modified Person 2", person2Pointer, person2)
}

func modifyAge(newPerson *Person) Person {
	// Your code here
	newPerson.Age = newPerson.Age + 1
	return *newPerson
}

func modifyName(newPerson *Person) Person {
	newPerson.Name = "New name hehe"
	return *newPerson
}
