package main

import (
	"fmt"
)

// naked return exm

func naked(val int) (x, y int) {
	x = val + 3
	y = x - 3

	return
}

// variables

var c, python, java bool

func main() {

	fmt.Println(naked(10))

	// get return in separate variable
	x, z := naked(10)
	fmt.Println(x, z)
	var i int
	fmt.Println(i, c, python, java)
}
