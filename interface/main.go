package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Square struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius

}

func (c Square) area() float64 {

	return c.width * c.height
}

func getArea(s Shape) float64 {

	return s.area()
}

func main() {
	ci := Circle{0, 0, 5}

	re := Square{10, 5}

	fmt.Printf("Circle are: %f\n", getArea(ci))
	fmt.Printf("Square are: %f\n", getArea(re))

	fmt.Println("Hello World")

}
