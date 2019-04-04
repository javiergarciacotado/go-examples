package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (person Person) String() string {
	return fmt.Sprintf("%v (%d years)", person.Name, person.Age)
}

func seventhMain() { //rename to main for its execution
	a := Person{Name: "Javier Garcia Cotado"}
	b := Person{"Matusalen", 12000000}
	fmt.Println(a.String(), b.String())
}
