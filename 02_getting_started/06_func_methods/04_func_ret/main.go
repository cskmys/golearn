// in Go,
// you get a compilation error if the return value is missing
// it doesn't matter if you catch the return value or not, at the place of invocation
// you can return multiple values
// there is no exception throwing mechanism, you just return something and let the caller call panic if it wants to crash the code

package main

import (
	"errors"
	"fmt"
)

func main() {
	port := 3000
	isStarted := startWebserver(port, 2) // you can use implicit initialization syntax to catch the return type
	// as there is return datatype in function signature, compiler can easily infer it
	fmt.Println(isStarted)
	nilErr := startWebserverWErr(port)
	fmt.Println(nilErr)
	_, err := startWebserverWErrRetErr(port) // to receive multiple return values, you can use a comma delimited list
	// whichever value you want to ignore you can use a write-only variable "_" instead of a regular variable
	fmt.Println(err)
}

func startWebserver(port, nbRetries int) bool { // the return type is declared after the arguments, reverse of C/C++
	fmt.Println("Starting server...")
	// do stuff
	fmt.Println("Server started on port", port)
	fmt.Println("Nb retries", nbRetries)
	return true
}

func startWebserverWErr(port int) error { // returning true and false is not the Go way of doing things
	// when everything is fine return "nil" and when something goes wrong, return what went wrong
	fmt.Println("Starting server...")
	// do stuff
	fmt.Println("Server started on port", port)
	return nil
}

func startWebserverWErrRetErr(port int) (int, error) { // to return multiple values you put all their types in a tuple
	fmt.Println("Starting server...")
	// do stuff
	fmt.Println("Server started on port", port)
	// "errors" package can be used to create errors
	return port, errors.New("something went wrong") // to return multiple values just express them as a comma delimited list
}
