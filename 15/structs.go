package main

import "fmt"

type vertex struct { //collection of fields
	x int
	y int
}

func main() {
	v1 := vertex{1, 2}
	v1.x = 4 //struct fields are accessed by .

	v2 := vertex{}     //x: 0, y: 0
	v3 := vertex{x: 1} //y:0
	v4 := &vertex{1, 2}

	fmt.Println(v1, v2, v3, v4)
}
