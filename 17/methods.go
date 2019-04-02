package main

import (
	"fmt"
	"math"
)

type vertex struct {
	x, y float64
}

func (v vertex) abs() float64 { //a method is just a function with a receiver argument
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func firstMain() {
	v := vertex{3, 4}
	fmt.Println(v.abs())
}
