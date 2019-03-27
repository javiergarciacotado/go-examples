package main

import "fmt"

func thirdMain() {
	var i, j int = 1, 2
	k := 3 //inside a function, the := can be used in place of a var declaration with implicit type
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
