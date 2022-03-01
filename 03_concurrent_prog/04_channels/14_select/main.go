package main

import (
	"fmt"
	"sync"
)

func main() {
	selectBlocking()
	// selectNonblocking()
}

func selectBlocking() {
	wg := &sync.WaitGroup{}
	ch1 := make(chan int)
	ch2 := make(chan string)
	wg.Add(3)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}(ch1, wg)
	go func(ch1 <-chan int, ch2 chan<- string, wg *sync.WaitGroup) {
		one := false
		two := false
		for {
			select { // blocking "select" as there is no default statement
			case i, temp := <-ch1: // first this will be run
				if temp == true {
					fmt.Println(i)
				}
				one = !temp
			case ch2 <- "hello": // second this will be run
				two = true
			}
			if one == true && two == true {
				break
			}
		}
		wg.Done()
	}(ch1, ch2, wg)
	go func(ch <-chan string, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		wg.Done()
	}(ch2, wg)
	wg.Wait()
}

func selectNonBlocking() {
	wg := &sync.WaitGroup{}
	ch1 := make(chan int)
	ch2 := make(chan string)
	wg.Add(3)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}(ch1, wg)
	go func(ch1 <-chan int, ch2 chan<- string, wg *sync.WaitGroup) {
		one := false
		two := false
		for {
			select { // not blocking "select" as there is a default statement
			case i, temp := <-ch1: // first this will be checked
				if temp == true {
					fmt.Println(i)
				}
				one = !temp
			case ch2 <- "hello": // second this will be checked
				two = true
			default:
			}
			if one == true && two == true {
				break
			}
		}
		wg.Done()
	}(ch1, ch2, wg)
	go func(ch <-chan string, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		wg.Done()
	}(ch2, wg)
	wg.Wait()
}
