package main

import "fmt"

type I2 interface {
	M()
}

func main() {
	var i2 I2
	describe(i2)
	i2.M() //is a run-time error because there is no "concreate" type inside
}

func describe(i2 I2) {
	fmt.Printf("(%v %T)\n", i2, i2)
}
