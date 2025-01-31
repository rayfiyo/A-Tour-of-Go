package main

import (
	"fmt"
	"math"
)

func compute34(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute34(hypot))
	fmt.Println(compute34(math.Pow))
}
