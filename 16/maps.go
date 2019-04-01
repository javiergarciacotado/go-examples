package main

import "fmt"

type vertex struct {
	lat, long float64
}

func thirdMain() {
	var aMap = map[string]vertex{ //or use make
		"test": vertex{1, 2},
	}

	fmt.Println(aMap["test"])

	anotherMap := make(map[string]int)
	anotherMap["answer"] = 42
	fmt.Println("Value:", anotherMap["answer"])
	anotherMap["answer"] = 48
	fmt.Println("Value:", anotherMap["answer"])
	delete(anotherMap, "answer")
	fmt.Println("Value:", anotherMap["answer"])

	v, ok := anotherMap["answer"] //is the key present?
	fmt.Println("Value:", v, "Present?", ok)
}
