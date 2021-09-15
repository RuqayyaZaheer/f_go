package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello World")

	// how to declear variable
	var x int = 10
	var y int = 10
	var sum int = x + y

	fmt.Println(sum)
	// short hand syntex
	a := 5
	b := 5
	sumAb := a + b

	fmt.Println(sumAb)

	i := 3
	if i > 2 {
		fmt.Println("i is grater then 3")
	} else if i < 2 {
		fmt.Println("i is less then 2")
	} else {
		fmt.Println("its an else statement")
	}

	// array in go []
	// go initialize element in 0

	// var simpleArray [5]int
	// short hand for array
	simpleArray := []int{1, 2, 3, 4, 5}

	// add value in array

	// simpleArray[2] = 19

	// use build in append function to add a element in array

	simpleArray = append(simpleArray, 13)

	fmt.Println(simpleArray)

	/* map in GO
	what is map in go?
	ans : its a key value pair

	["name"] = "your name"

	how to use first call map then [] inside [] use key type then outside of [] value type
	all inside make function

	exm :
		slist := make (map[string]string)


	*/

	list := make(map[string]string)
	list["name"] = "Student 1"
	list["address"] = "Saitama"
	list["fev_food"] = "Vegetable"
	list["delete_value"] = "This value add for delete"
	// Using Delete function
	fmt.Println("Before use delete Function", list)
	delete(list, "delete_value")

	fmt.Println("after delete", list)

	// There is only one loop in Go for loop
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// while loop in go
	n := 9
	for n < 10 {
		n -= 1
		fmt.Println(n)
		if n == 0 {
			break
		}
	}

	// what is sum of 1+2+3+..... +10

	f := 0

	fmt.Println("sum print start")
	for z := 1; z <= 10; z++ {
		f += z
	}
	fmt.Println(f)

}
