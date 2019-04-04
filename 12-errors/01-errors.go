package main

import (
	"fmt"
	"strconv"
	"time"
)

type MyError struct {
	when time.Time
	what string
}

func (myError *MyError) Error() string {
	return fmt.Sprintf("at %v %s", myError.when, myError.what)
}

func run() error {
	return &MyError{
		when: time.Now(),
		what: "it didn't work"}
}

func firstMain() { //rename to main for its execution

	i, err := strconv.Atoi("42")
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	}
	fmt.Printf("Converted integer: %d\n", i)

	if err := run(); err != nil {
		fmt.Println(err)
	}
}
