package Pointers

import "fmt"

func BasicDeclaration() {
	x := 42
	p := &x

	fmt.Println("original value of x: ", x)
	fmt.Println("value through x", p, &p, &x)

	*p = 80
	fmt.Println("value x after change p", *p)
}
