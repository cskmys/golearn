package main

import (
	"fmt"
	"sync"
)

func main() {
	// chComm()
	// panicErr()
	// rxOnClosedUnbuffCh()
	rxOnClosedBuffCh()
}

func chComm() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 42
		close(ch) // closing channel after send has causes no problem cause the receiver can still listen from a closed channel
		// however sending messages into a closed channel is going to cause panic
		// you cannot reopen a closed channel
		wg.Done()
	}(ch, wg)
	wg.Wait()
}

func panicErr() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 42
		close(ch)
		ch <- 27 // causes panic coz it is trying to send value on closed channel
		wg.Done()
	}(ch, wg)
	wg.Wait()
}

func rxOnClosedUnbuffCh() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) { // made it a bidirectional channel argument coz we cannot close a channel on the receiving end(i.e. if its passed as receive-only)
		fmt.Println(<-ch)
		close(ch)
		fmt.Println(<-ch) // reading from a closed unbuffered channel will give 0
		// there's no api to check if a channel is closed or not,
		// however, we can use comma ok syntax to check if the received value is valid or not
		wg.Done()
	}(ch, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 42
		wg.Done()
	}(ch, wg)
	wg.Wait()
}

func rxOnClosedBuffCh() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) { // made it a bidirectional channel argument coz we cannot close a channel on the receiving end(i.e. if its passed as receive-only)
		fmt.Println(<-ch)
		close(ch)
		fmt.Println(<-ch) // reading from a closed buffered channel will give the unread value until they exist
		fmt.Println(<-ch) // once all the buffered values have been read out, you'll get 0 on closed buffered channel
		wg.Done()
	}(ch, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 42
		ch <- 27
		wg.Done()
	}(ch, wg)
	wg.Wait()
}
