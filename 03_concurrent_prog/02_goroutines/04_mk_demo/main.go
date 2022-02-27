// here we build  a simulation of a database query system that has an in-memory cache in front of it

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var cache = map[int]Book{}                                // the in-memory cache at the front that will refer to "Book" object by an integer key
var rnd = rand.New(rand.NewSource(time.Now().UnixNano())) // random nb generator for building queries

func main() {
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1           // "+ 1" coz "Book::ID" is 1-10
		if b, ok := queryCache(id); ok { // combining line "b, ok := queryCache(id)" and "if ok == true {"
			fmt.Println("from cache")
			fmt.Println(b)
			continue
		}
		if b, ok := queryDatabase(id); ok {
			fmt.Println("from database")
			fmt.Println(b)
			continue
		}
		fmt.Printf("book not found with id: '%v'\n", id) // similar to "fmt.Print" but with the ability to format strings
		time.Sleep(150 * time.Millisecond)
	}
}

func queryCache(id int) (Book, bool) {
	b, ok := cache[id] // if there is a hit in "cache", "ok" will be "true", otherwise "false"
	return b, ok
}

func queryDatabase(id int) (Book, bool) {
	time.Sleep(100 * time.Millisecond) // to simulate the delay in accessing database data
	for _, b := range books {
		if b.ID == id {
			cache[id] = b
			return b, true
		}
	}
	return Book{}, false
}
