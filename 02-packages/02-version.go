package main

import "fmt"

func init() { //called when a package is declared
	fmt.Println("greeting/day.go ==>  init()")
}

var version = "1.0.0"
