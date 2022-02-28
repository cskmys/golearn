package main

import (
	"fmt"
	"sync"
)

func main() {
	// deadlock1()
	// deadlock2()
	chComm()
}

func chComm() {
	wg := &sync.WaitGroup{}
	ch := make(chan int) // creating an unbuffered channel
	// for unbuffered channel, every goroutine that wants to send a message on the channel is blocked until it is received by another goroutine on the channel

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch) // "<-ch" indicates want to print anything that comes out of the channel, but this goroutine doesn't know who the sender is
		wg.Done()
	}(ch, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 42 // sending 42 to "ch", but this goroutine doesn't know who'll receive it
		wg.Done()
	}(ch, wg)

	wg.Wait()
}

func deadlock1() {
	ch := make(chan int)
	fmt.Println(<-ch) // goroutine expects to receive a msg from the channel, and until it gets one it is blocked
	ch <- 42
}

func deadlock2() {
	ch := make(chan int)
	ch <- 42 // goroutine expects a receiver on the other end, and until it finds one on the channel it is blocked
	// this is done because if there is no one to receive, the goroutine can sleep and let others run
	fmt.Println(<-ch)
}
