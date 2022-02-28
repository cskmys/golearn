package main

import (
	"fmt"
	"sync"
)

func main() {
	// chComm()
	// chRxLoopDeadlock()
	chRxLoop()
}

func chComm() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		wg.Done()
	}(ch, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		wg.Done()
	}(ch, wg)
	// looping works well when you exactly know how many msgs will be sent and received
	// but in many scenarios the sender will not know how many it will be sending and the receiver similarly won't be knowing how many it has to receive
	wg.Wait()
}

func chRxLoopDeadlock() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		for msg := range ch { // after receiving 10 times and going to receive for 11th time is going to block the goroutine
			// but the other goroutine is sleeping after sending 10 msgs, hence there's a deadlock which cause go runtime code to throw panic
			fmt.Println(msg)
		}
		wg.Done()
	}(ch, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		wg.Done()
	}(ch, wg)
	wg.Wait()
}

func chRxLoop() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		for msg := range ch { // for loop will terminate once it reads from a closed channel
			fmt.Println(msg)
		}
		wg.Done()
	}(ch, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch) // now after all the sending, we close the channel which will cause the "for" loop in the other goroutine to terminate
		wg.Done()
	}(ch, wg)
	wg.Wait()
}
