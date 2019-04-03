package main //Every Go program must be a part of some package

import (
	"errors"
	"fmt"
	"go-examples/02-packages/greeting"
	"math/rand"
)

func firstMain() { //rename to main for its execution
	fmt.Println("My favourite number is", rand.Intn(10)) //from go distribution
	error := errors.New("Some error")
	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(greeting.Morning) //from local package
	version = "1.0.1"             //variable re-declaration from version.go
	fmt.Println("version ==> ", version)

}
