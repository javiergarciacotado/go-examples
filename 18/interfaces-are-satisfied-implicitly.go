package main

import "fmt"

type I interface {
	m()
}

type T struct {
	S string
}

func (t T) m() { //no "implements" keyword
	fmt.Println(t)
}

func main() {
	var i I = T{"hello"}
	i.m()
}
