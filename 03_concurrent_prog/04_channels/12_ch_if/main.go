package main

import (
	"fmt"
	"sync"
)

func main() {
	// rxOnClosedUnbuffCh()
	// rxOnClosedBuffCh()
	rxOnClosedBuffChIfOk()
}

func rxOnClosedUnbuffCh() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		msg, ok := <-ch
		fmt.Println(msg, ok) // true in the output here indicates that the 0 read from channel is valid
		close(ch)
		msg, ok = <-ch
		fmt.Println(msg, ok) // false in the output here indicates that the 0 read from channel is invalid i.e. channel is closed
		wg.Done()
	}(ch, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 0
		wg.Done()
	}(ch, wg)
	wg.Wait()
}

func rxOnClosedBuffCh() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		msg, ok := <-ch
		fmt.Println(msg, ok)
		close(ch)
		msg, ok = <-ch
		fmt.Println(msg, ok) // true in the output here indicates that the 0 read from channel is valid
		msg, ok = <-ch
		fmt.Println(msg, ok) // false in the output here indicates that the 0 read from channel is invalid i.e. channel is closed
		wg.Done()
	}(ch, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 42
		ch <- 0
		wg.Done()
	}(ch, wg)
	wg.Wait()
}

func rxOnClosedBuffChIfOk() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		if msg, ok := <-ch; ok { // now we print msg only if it is valid
			fmt.Println(msg)
		}
		close(ch)
		if msg, ok := <-ch; ok {
			fmt.Println(msg)
		}
		if msg, ok := <-ch; ok { // this one is invalid we never print the value
			fmt.Println(msg)
		}
		wg.Done()
	}(ch, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 42
		ch <- 0
		wg.Done()
	}(ch, wg)
	wg.Wait()
}
