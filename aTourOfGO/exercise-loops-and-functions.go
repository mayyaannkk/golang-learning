package tourOfGo

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := x / 2
	var prev float64
	i := 0
	for ; prev != z; i++ {
		prev = z
		z -= (z*z - x) / (2 * z)
	}
	fmt.Println(i)
	return z
}

func FindSqrt() {
	fmt.Println(Sqrt(25))
}
