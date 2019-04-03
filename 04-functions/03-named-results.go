package main

import "fmt"

func split(number int) (x, y int) {
	x = number * 4 / 9
	y = number - x
	return //naked return as they are defined at the top of the function. As a recommemndation, they should be used only in short functions
}

func thirdMain() { //rename to main for its execution
	fmt.Println(split(21))
}
