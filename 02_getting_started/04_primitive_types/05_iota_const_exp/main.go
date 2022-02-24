// in Go,
// unlike unused variables, unused constant will not cause compilation errors

package main

import "fmt"

const pi = 3.1415 // constant at package level i.e. global scope

const (
	first  = 1
	second = "second" // "second" is unused, and it doesn't cause compilation error
)

const (
	premiere  = iota
	deuxieme  = iota // every time "iota" is used, its value is incremented by 1
	troisieme = iota // this is helpful in creating a long chain of consecutive constant values
)

const (
	// "iota" in this "const" block is independent of last "iota" value in previous "const" block
	ondhu = 2 + iota // you can do arithmetic with "iota"
	eradu = iota * 6
	muru  = 1 << iota // 2^iota
)

const (
	okkati = 3 * iota
	rendu  // automatically first expression "3 * iota" is inserted for all subsequent constants
	mudu
)

func main() {
	fmt.Println(pi)
	fmt.Println(first)
	fmt.Println(premiere, deuxieme, troisieme)
	fmt.Println(ondhu, eradu, muru)
	fmt.Println(okkati, rendu, mudu)
}
