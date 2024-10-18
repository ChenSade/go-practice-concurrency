package main

import "fmt"

// e.g. for multiple concurrent jobs pulling items out a queue for processing.
// can check CPU utilization for making sure several cores are participating in the work

func main() {
	// create buffered channels
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// we add several workers to take advantage of the multiple cores CPU.
	// the order of the fibonacci numbers is not guaranteed.
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	// now worker will pull numbers one-by-one, calculate fibonacci and send back to the results channel
	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	// now we expect 100 results to eventually appear on the results channel (the first 100 fibonacci numbers)
	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

// worker calculate fibonacci sequence. jobs: only
// instead of using bidirectional channels,
// we set the Jobs channel to only *Receive*, and Results channel to only *Send*
// that makes it more clear and explicit
func worker(jobs <-chan int, results chan<- int) {
	// get numbers from the jobs, and send the result on the results channel
	for n := range jobs {
		results <- fibonacci(n)
	}
}

// calculate fibonacci sequence
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
