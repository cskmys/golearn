// In Go,
// to provide an implementation of an interface in a package, all we need to do is make sure that the package is imported and write a function with same name and signature as the interface in the package that we want to implement

package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) { // in "net/http" package "Handler" is an interface type that expects implementation of "ServeHTTP(http.ResponseWriter, *http.Request)"
	// because we have imported the package "net/http", the compiler automatically figures out that this function we have whose name and signature is the same as the interface, is an implementation of the interface in "Handler" type
	w.Write([]byte("Hello from the user controller!"))
}

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
