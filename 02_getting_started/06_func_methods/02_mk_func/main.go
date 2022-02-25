// in Go,
// unlike unused variable, unused function does not throw compilation error(obvious coz think of the library code, not all functions are used)

package main

import "fmt"

func main() { // "func" keyword to declare a function
	startWebserver() // if you don't include parenthesis post function name, compiler will think you are passing a function pointer
	// parenthesis make it clear that you are invoking the function

} // curly braces in KnR style to create body of the function

func startWebserver() {
	fmt.Println("Starting server...")
	// do stuff
	fmt.Println("Server started")
}
