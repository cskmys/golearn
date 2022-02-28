// to look for race conditions, do
// go run --race .

// here we remove the delay used for co-ordinating memory access and resolve the data race condition using mutex

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
	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.Mutex) {
			if b, ok := queryCache(id, m); ok {
				fmt.Println("from cache")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m)
		go func(id int, wg *sync.WaitGroup, m *sync.Mutex) {
			if b, ok := queryDatabase(id, m); ok {
				fmt.Println("from database")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m)
		// here the delay used previously for co-ordination is deleted and data race is solved by using mutex
		// however, interleaving of output is still a problem
	}
	wg.Wait()
}

func queryCache(id int, m *sync.Mutex) (Book, bool) {
	m.Lock() // the goroutine that calls the "sync::Mutex::Lock" owns the mutex and any other goroutine that tries to acquire the lock will be paused until the goroutine that owns it unlocks it
	b, ok := cache[id]
	m.Unlock()
	return b, ok
}

func queryDatabase(id int, m *sync.Mutex) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			m.Lock() // after the "cache" is filled, very less goroutines come until here, so there are more readers of "cache" than writers
			// when at least one goroutine tries to write to "cache" when one or more goroutines are reading it causes data race
			// hence, we need to block readers only when write operation is taking place, otherwise we can let multiple readers at same time, and
			// we need to block writers when another read/write operation is taking place at same time
			// therefore, simply blocking any type of access is inefficient and to improve efficiency this type of conditional locking needs to be done
			cache[id] = b
			m.Unlock()
			return b, true
		}
	}
	return Book{}, false
}
