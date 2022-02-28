package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// deadlock()
	// noDeadlock()
	// chComm()
}

func deadlock() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1) // bidirectional channel

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		time.Sleep(5 * time.Millisecond) // simulating delay in scheduling this goroutine
		fmt.Println(<-ch)
		ch <- 27
		wg.Done()
	}(ch, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 42
		fmt.Println(<-ch)
		// if "ch" is buffered and for whatever the reason, the other goroutine gets scheduled late,
		// it is possible that this goroutine sends 42 on the channel and itself receives and prints 42 and goes to sleep
		// in this case, the other goroutine will not get scheduled at all coz there is now nobody to put something on the channel for it to receive
		// hence this gives raise to a deadlock
		wg.Done()
	}(ch, wg)

	wg.Wait()
}

func noDeadlock() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		time.Sleep(5 * time.Millisecond)
		fmt.Println(<-ch)
		ch <- 27
		wg.Done()
	}(ch, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 42
		fmt.Println(<-ch)

		// if "ch" is unbuffered then no problem coz even if there's a delay in scheduling other goroutine,
		// this goroutine also gets blocked until there's a receiver on the other end of the channel
		// hence each send and receive on this goroutine will have a corresponding match on the other goroutine and there's no deadlock
		wg.Done()
	}(ch, wg)

	wg.Wait()
}

func chComm() {
	wg := &sync.WaitGroup{}
	ch := make(chan int) // bidirectional channel

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) { // bidirectional channel passed as receive only channel using "<-chan" in the function signature
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) { // bidirectional channel passed as send only channel using "chan<-" in the function signature
		ch <- 42
		wg.Done()
	}(ch, wg)
	// arrow "<-" is always pointing to the left
	wg.Wait()
}
