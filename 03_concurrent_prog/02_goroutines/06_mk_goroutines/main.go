// In Go,
// the runtime code provides goroutines that uses OS threads to implement and run user threads

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	//sequential()
	parallel()
}

func queryCache(id int) (Book, bool) {
	b, ok := cache[id]
	return b, ok
}

func queryDatabase(id int) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			cache[id] = b
			return b, true
		}
	}
	return Book{}, false
}

func sequential() {
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		if b, ok := queryCache(id); ok {
			fmt.Println("from cache")
			fmt.Println(b)
			continue
		}
		if b, ok := queryDatabase(id); ok {
			fmt.Println("from database")
			fmt.Println(b)
			continue
		}
		fmt.Printf("book not found with id: '%v'\n", id)
		time.Sleep(150 * time.Millisecond)
	}
}

func parallel() {
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		go func(id int) { // keyword "go" converts this function into a goroutine which runs concurrently with other goroutines
			if b, ok := queryCache(id); ok {
				fmt.Println("from cache")
				fmt.Println(b)
			}
		}(id) // in-place anonymous function
		go func(id int) { // we have "queryCache" and "queryDatabase" executing in parallel, rather than having "queryDatabase" execute only when "queryCache" is unsuccessful
			if b, ok := queryDatabase(id); ok {
				fmt.Println("from database")
				fmt.Println(b)
			}
		}(id)
		time.Sleep(150 * time.Millisecond) // if you comment this you see zero output
		// reason is main routine just loops 10 times creating 2 goroutines each time and finishes execution before any of the goroutines finish their execution
		// once main routine is done, all goroutines are cleaned up as well
		// having that "time::Sleep" pauses main routine and gives goroutines some time to execute
	}
	time.Sleep(2 * time.Second) // if you comment this, main routine can still finish execution before the last few goroutines execute
	// hence additional sleep of 2 seconds is added to make sure that main routine waits and gives time for the last few goroutines to finish their execution
}
