package concurrency

import (
	"fmt"
	"time"
)

func Count1(thing string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 50)
	}
}

func CountChannelNotClosed(thing string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 50)
	}
}

func CountChannelIsClosed(thing string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 50)
	}
	close(c)
}
