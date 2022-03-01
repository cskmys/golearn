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
	m := &sync.RWMutex{}
	cacheCh := make(chan Book)
	dbCh := make(chan Book)
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, cacheCh chan<- Book) {
			if b, ok := queryCache(id, m); ok {
				cacheCh <- b
			}
			wg.Done()
		}(id, wg, m, cacheCh)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, dbCh chan<- Book) {
			if b, ok := queryDatabase(id, m); ok {
				dbCh <- b
			}
			wg.Done()
		}(id, wg, m, dbCh)
		// everytime 2 goroutines are used to simultaneously search cache and database
		// whoever finds it will send the book on the channel
		// please note that database will always succeed and cache will conditionally succeed
		// hence, unlike cache channel there's always a book on the database channel
		go func(wg *sync.WaitGroup, cacheCh, dbCh <-chan Book) {
			select {
			case b := <-cacheCh: // first preference
				fmt.Println("from cache")
				fmt.Println(b)
				<-dbCh // clearing "dbCh" as it will get msg on it even when the data was found on cache
				// by doing this, when this case runs the other case doesn't run which is what we wanted
			case b := <-dbCh: // second preference
				fmt.Println("from database")
				fmt.Println(b)
			}
		}(wg, cacheCh, dbCh)
		time.Sleep(150 * time.Millisecond) // to give all the goroutines some time to execute
		// commenting above line is going to screw up everything, and you see interleaving and not all 10 books are printed
		// meaning, not all goroutines are running to completion
	}
	wg.Wait()
}

func queryCache(id int, m *sync.RWMutex) (Book, bool) {
	m.RLock()
	b, ok := cache[id]
	m.RUnlock()
	return b, ok
}

func queryDatabase(id int, m *sync.RWMutex) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			m.Lock()
			cache[id] = b
			m.Unlock()
			return b, true
		}
	}
	return Book{}, false
}
