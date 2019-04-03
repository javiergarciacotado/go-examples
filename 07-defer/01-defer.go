package main

import "fmt"

func firstMain() {
	defer fmt.Println("world") //it will be executed when the surrounding function returns

	fmt.Println("hello")
}
