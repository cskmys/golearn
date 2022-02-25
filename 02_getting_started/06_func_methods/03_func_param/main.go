// in Go,
// unused function parameter doesn't cause compilation error like unused variable does

// argument and parameter are used interchangeably
// parameter is what is declared in the function signature and argument is passed during the function invocation

package main

import "fmt"

func main() {
	port := 3000
	startWebserver(port, 2) // you can pass a variable or a constant directly
}

func startWebserver(port int, nbRetries int) { // signature is reverse of C/C++
	fmt.Println("Starting server...")
	// do stuff
	fmt.Println("Server started on port", port)
	fmt.Println("Nb retries", nbRetries)
}

func wackySig(port, nbRetries int) { // if 2 params have same data type, you can combine and declare them together like tthis

}
