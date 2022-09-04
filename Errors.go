package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative numer: %f", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	z := x
	previous_z := x
	for i := 0; i < 100; i++ {
		z -= (z*z - x) / (2*z)
		if math.Abs(previous_z - z) <= 0.00000000000001 {
			// fmt.Println(i)
			return z, nil
		}
		previous_z = z
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
