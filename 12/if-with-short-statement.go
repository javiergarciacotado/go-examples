package main

import (
	"fmt"
	"math"
)

func pow(base, power, limit float64) float64 {
	//to execute before condition and its scope occurs until the end of the if (including else)
	if v := math.Pow(base, power); v < limit {
		return v
	}
	return limit
}

func sixthMain() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 2, 8),
	)
}
