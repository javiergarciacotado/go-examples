package main

import (
	"fmt"
)

func do(i interface{}) {
	switch v := i.(type) { //switch v := i.(T) ==> is a construct that permits several type assertions in series
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about the type %T\n", v)
	}
}

func sixthMain() { //rename to main for its execution
	do(21)
	do("hello")
	do(true)
}
