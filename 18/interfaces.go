package main

import (
	"fmt"
	"math"
)

type myFloat float64

func (f myFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type vertex struct {
	X, Y float64
}

func (v *vertex) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Abser interface {
	abs() float64
}

func firstMain() {
	var a Abser
	f := myFloat(-math.Sqrt2)
	v := vertex{3, 4}

	a = f
	a = &v

	fmt.Println(a.abs())

}
