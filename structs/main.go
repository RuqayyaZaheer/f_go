package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// methord in go(value receiver)

func (p Person) greet() string {

	return "Hello my name is " + p.firstName + " " + p.lastName + ", and i am " + strconv.Itoa(p.age)
}

// methord in go (point receiver)

func (p *Person) hasBirthday() {

	p.age++

}

// func which can change value (Point receiver)

func (p *Person) getMarrid(spouseLastName string) {

	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {

	person1 := Person{firstName: "Angelina", lastName: "Bur", city: "Bostin", age: 19, gender: "f"}

	// fmt.Println(person1)
	// fmt.Println(person1.gender)

	fmt.Println(person1.greet())

	person1.hasBirthday()

	person1.getMarrid("zz")
	fmt.Println(person1.greet())

}
