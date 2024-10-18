package main

import (
	"fmt"
	"go-practice-concurrency/concurrency"
	"sync"
)

func main() {

	// 1. Concurrency With WaitGroup

	fmt.Println("--- Concurrency With WaitGroup ---")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		concurrency.Count1("sheep")
		wg.Done()
	}()
	wg.Wait()

	// 2. Concurrency With Channels 1
	// this prints "fish" only once... not good enough

	fmt.Println("--- Concurrency With Channels 1 ---")
	c := make(chan string)
	go concurrency.CountChannelNotClosed("fish", c)
	msg := <-c
	fmt.Println(msg)

	// 3. Concurrency With Channels 2
	// this prints "fish" 5 times as expected,
	// but then crashes with a deadlock, because "main" func is waiting for more messages on the channel

	//fmt.Println("--- Concurrency With Channels 2 (crashes the program, thus commented out) ---")
	//c2 := make(chan string)
	//go concurrency.CountChannelNotClosed("fish", c2)
	//for {
	//	msg := <-c2
	//	fmt.Println(msg)
	//}

	// 4. Concurrency With Channels 3

	fmt.Println("--- Concurrency With Channels 3 ---")
	c3 := make(chan string)
	go concurrency.CountChannelIsClosed("chick", c3)
	for {
		msg, open := <-c3
		if !open {
			break
		}
		fmt.Println(msg)
	}

	// 5. Concurrency With Channels 4: using "range" to auto-close channel without checking explicitly

	fmt.Println("--- Concurrency With Channels 4 ---")
	c4 := make(chan string)
	go concurrency.CountChannelIsClosed("pig", c4)
	for msg := range c4 {
		fmt.Println(msg)
	}

	// 6. Concurrency which blocks with deadlock (thus commented-out)

	//fmt.Println("--- Concurrency With Channels 5 ---")
	//c5 := make(chan string)
	//// this line blocks main until something receives from the channel, but can't get to the "receive" line
	//// since this "send" line is blocking in the same routine
	//c5 <- "YO YO!"

	// never gets here to the "receive" line.
	// to make that work, need to call this in a separate goroutine,
	// or work with buffered channel as done in the next e.g.
	//msg2 := <-c5
	//fmt.Println(msg2)

	// 7. Concurrency with buffered channels. does not block

	fmt.Println("--- Concurrency With Buffered Channels 6 ---")
	c6 := make(chan string, 1)
	c6 <- "YO YO!"
	msg2 := <-c6
	fmt.Println(msg2)

	// 8. Concurrency with buffered channels.
	// this crashes since trying to send 3 messages in a channel with a capacity of 2

	//fmt.Println("--- Concurrency With Buffered Channels 7 ---")
	//c7 := make(chan string, 2)
	//c7 <- "YO YO!"
	//c7 <- "Yada Yada!"
	//c7 <- "Bla Bla!"
}
