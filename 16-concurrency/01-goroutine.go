package main

import (
	"fmt"
	"time"
)

func say(message string) {
	fmt.Printf("that is '%s' from function invocation\n", message)
	for index := 0; index < 5; index++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(message)
	}
}

func firstMain() { //rename to main for its execution
	go say("world") //Execution happens in a lightweight thread managed by the Go runtime
	say("Hello")
}
