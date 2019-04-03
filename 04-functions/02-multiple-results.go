package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func secondMain() { //rename to main for its execution
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
