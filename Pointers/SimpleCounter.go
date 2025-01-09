package Pointers

import "fmt"

type Counter struct {
	count int
}

func SimpleCounter() {
	counter := Counter{count: 0}

	fmt.Printf("Initial count: %d\n", counter.count)

	// Your task: Write two functions
	// 1. increment: adds 1 to count
	// 2. decrement: subtracts 1 from count
	increment(&counter)
	increment(&counter)
	decrement(&counter)

	fmt.Printf("Final count: %d\n", counter.count)
}
func increment(c *Counter) {
	// Write your code here
	c.count += 1
}

func decrement(c *Counter) {
	// Write your code here
	c.count -= 1
}
