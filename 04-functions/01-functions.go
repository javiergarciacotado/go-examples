package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func firstMain() { //rename to main for its execution
	fmt.Println(add(42, 13))
}
