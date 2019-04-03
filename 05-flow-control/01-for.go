package main

import "fmt"

func firstMain() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("%v\n", sum)
}
