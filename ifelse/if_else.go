package main

import (
	"fmt"
	"math"
)

func pow(x, y, linf float64) float64 {

	return math.Pow(x, y)

}

func main() {
	fmt.Println(math.Pow(3, 3))
	fmt.Println(pow(5, 5, 10))
}
