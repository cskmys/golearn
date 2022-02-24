package main

import "fmt"

func main() {
	// short declaration syntax used in almost all the situations
	m := map[string]int{"foo": 42} // map which maps string(key) to integer(value)
	fmt.Println(m)
	fmt.Println(m["foo"]) // indexing map using string key to get integer value

	m["foo"] = 27
	fmt.Println(m["foo"])

	delete(m, "foo") // remove elements from map
	fmt.Println(m)
}
