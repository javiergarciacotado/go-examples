package main

import "fmt"

type I2 interface {
	M()
}

func thirdMain() { //rename to main for its execution
	var i2 I2
	Describe(i2)
	i2.M() //is a run-time error because there is no "concreate" type inside
}


