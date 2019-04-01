package main

import "fmt"

func main() {

	i, j := 42, 2701

	p := &i //A pointer holds the memory address of a value
	*p = 21
	fmt.Println(*p)

	p = &j
	fmt.Println(*p)
}
