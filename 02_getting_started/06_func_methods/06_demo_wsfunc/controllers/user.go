package controllers

import "net/http"

type userController struct { // empty struct to hang a behavior
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) { // to bind a function to a struct as a method, you need to declare it in parentheses after "func" keyword in function signature
	// in Go, we avoid using "this" as a name and instead give it a proper name
	w.Write([]byte("Hello from the user controller!")) // "http.ResponseWriter::Write" takes in a slice of bytes, and we can directly convert string into byte slice
	// in Go, a string is an alias for byte slice
}
