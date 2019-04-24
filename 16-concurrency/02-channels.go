package main

import "fmt"

func sum(integers []int, channel chan int) {
	sum := 0
	for _, v := range integers {
		sum += v
	}
	channel <- sum
}

func secondMain() { //rename to main for its execution

	integers := []int{7, 2, 8, -9, 4, 0}
	channel := make(chan int)

	go sum(integers[:len(integers)/2], channel)
	go sum(integers[len(integers)/2:], channel)
	x, y := <-channel, <-channel //sends and receives block until the other side (ending of goroutine) is ready

	fmt.Println(x, y, x+y)
}
