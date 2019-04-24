package main

import "fmt"

func thirdMain() { //rename to main for its execution
	integers := make(chan int, 2)
	integers <- 1
	integers <- 2
	fmt.Println(<-integers)
	fmt.Println(<-integers)
	fmt.Println(len(integers))
	integers <- 3
	fmt.Println(<-integers)

}
