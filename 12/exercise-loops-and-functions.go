package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	temp := 0.0
	delta := 1e-8

	for math.Abs(z-temp) > delta {
		z, temp = z-(z*z-x)/(2*z), z
	}
	return z
}

func exerciseMain() {
	fmt.Println(Sqrt(16))
}
