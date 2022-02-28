// main challenges with concurrent programming is: co-ordinating task execution and sharing resources

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	// using "time::Sleep" to co-ordinate is a bad idea, what if one/more goroutines take more time? then sleeping fails to ensure that all goroutines fully executed
	// what if they all finished early? then you're simply wasting time
	wg := &sync.WaitGroup{} // if you plan to pass the wait group to functions, then you don't want to copy it, hence you use a pointer
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(1)                             // doing a "sync::WaitGroup::Add(1)" before the goroutine to tell the main thread that it needs to wait on one goroutine to finish the execution
		go func(id int, wg *sync.WaitGroup) { // pass the wait group pointer into goroutine
			if b, ok := queryCache(id); ok {
				fmt.Println("from cache")
				fmt.Println(b)
			}
			wg.Done() // to indicate to the "WaitGroup" that current goroutine is finished its execution
		}(id, wg) // passing the arguments to in-place anonymous function
		wg.Add(1) // another "sync::WaitGroup::Add(1)" to tell main to wait on the 2nd goroutine
		// alternatively, the above line can be deleted and the first "sync::WaitGroup::Add" call simply do "sync::WaitGroup::Add(2)" to tell main to wait for 2 goroutines
		go func(id int, wg *sync.WaitGroup) {
			if b, ok := queryDatabase(id); ok {
				fmt.Println("from database")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg)
		time.Sleep(150 * time.Millisecond) // now that we have a "WaitGroup" we might think of removing this but removing this will lead to a fatal error
		// coz in "queryCache" we are reading "cache" and in "queryDatabase" we are writing "cache", both these are happening without any co-ordination
		// which means the read and write operation can take place at the same time on a shared memory which leads to runtime error
		// a simple solution to solve the problem for the time being would be to comment "cache" writing code in the "queryDatabase"
		// that would avoid the runtime error, however it would lead to a new error where two goroutines are printing on the console at the sametime which can cause their outputs to be interleaved
		// hence, "time::Sleep" cannot be removed for now though co-ordinating the tasks using delays is a very bad idea
	}
	wg.Wait() // waits until are goroutines have executed "wg.Done()"
}

func queryCache(id int) (Book, bool) {
	b, ok := cache[id]
	return b, ok
}

func queryDatabase(id int) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			cache[id] = b // commenting this can help avoid for the time being, the runtime fatal error caused by data race due to "cache" read in "queryCache"
			return b, true
		}
	}
	return Book{}, false
}
