package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x
	previous_z := x
	for i := 0; i < 100; i++ {
		z -= (z*z - x) / (2*z)
		if math.Abs(previous_z - z) <= 0.0000001 {
			fmt.Println(i)
			return z
		}
		previous_z = z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
