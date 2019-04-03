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

func secondMain() { //rename to main for its execution
	var i I = T{"hello"}
	i.m()
}
