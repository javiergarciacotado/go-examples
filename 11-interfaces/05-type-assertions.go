package main

import "fmt"

func fifthMain() { //rename to main for its execution

	var i interface{} = "hello"

	s := i.(string) //Does interface i hold the concrete type T? aka i.(T)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s)

	f, ok := i.(float64) //ok ==> false, and f will be the zero value of type T
	fmt.Println(f, ok)
}
