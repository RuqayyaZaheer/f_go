package main

import (
	"fmt"
)

func main() {
	a := 5
	b := &a

	fmt.Println(a, "\n", b)
	fmt.Printf("%T\n", b)

	// use * to read val from address

	fmt.Println(*b)
	fmt.Println(*&a)

	// give a value on pointer b

	*b = 10

	fmt.Println("printing ", a, "value")

}
