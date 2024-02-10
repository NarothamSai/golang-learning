package main

import (
	"fmt"
	"math"
)

func main() {
	val := 9.0
	fmt.Println(sqrt(val))
	fmt.Println(math.Sqrt(val))
}

func sqrt(x float64) float64 {
	z := 1.0

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}

	return z
}
