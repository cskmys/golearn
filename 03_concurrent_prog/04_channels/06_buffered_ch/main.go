package main

import (
	"fmt"
	"sync"
)

func main() {
	chComm()
}

func chComm() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1) // indicating the size makes it buffered channel
	// the internal capacity of an unbuffered channel is 0, hence the sender is blocked until the receiver receives the msg
	// for the buffer capacity "n", sender can send "n" msgs(that are not yet received) before it gets blocked

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch) // only 42 is printed, not 27 coz 27 is still on the buffer and is not yet received, hence 27 is lost
		wg.Done()
	}(ch, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 42
		ch <- 27
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
