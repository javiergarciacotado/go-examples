package main

import "fmt"

type myFloat float64

func (myFloat myFloat) abs() float64 { //type needs to be defined in the same package as the method
	if myFloat < 0 {
		return float64(-myFloat)
	}
	return float64(myFloat)
}

func main() {
	myFloat := myFloat(-2.0)
	fmt.Println(myFloat.abs())
}
