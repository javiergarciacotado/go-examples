package main

import "fmt"

func thirdMain() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Printf("%v\n", sum)
}
