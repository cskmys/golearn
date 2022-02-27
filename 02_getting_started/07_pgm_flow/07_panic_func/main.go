// in Go,
// there are no exceptions, when you hit a condition from which you cannot recover, you cause "panic"
// note there is a way to recover from "panic" but currently it is beyond the scope of our discussion here

package main

func main() {
	panic("shit")
}
