package main

import (
	"fmt"
	"time"
)

func main() {

	// Concurrency with channels "select"

	// this code triggers 2 go routines. shorter with 500ms interval, and longer with 2sec interval.

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "every two seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	// this infinite for-loop is working, but the shorter 500ms operation needs to wait 2 seconds
	// after each run to let the longer one finish, instead of running nicely in 500ms intervals.

	//for {
	//	fmt.Println(<-c1)
	//	fmt.Println(<-c2)
	//}

	// using "select" to print from whatever channel is ready. that way, the shorter 500ms routine does not
	// have to wait on the longer 2-sec one in every run. working nicely!

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
