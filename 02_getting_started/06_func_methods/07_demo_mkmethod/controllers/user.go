// In Go,
// there is no special constructor function

package controllers

import (
	"net/http"
	"regexp"
)

// eventually will handle routing responsibilities
// 2 types of resource requests:
// on the entire collection of users
// on an individual user
type userController struct {
	userIDPattern *regexp.Regexp // regular expression is used to discern which type of resource request is on the incoming HTTP requests path

}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the user controller!"))
}

// for the definition of the regex of the "userController", you need a constructor
func newUserController() *userController { // Go convention is to create a function by name "new<struct_name>" to create a struct constructor
	// returning a pointer is another convention to avoid doing a copy coz some of them can be quite heavy
	return &userController{ // taking address of directly without having a variable is allowed only in case of "struct", for example you cannot do "&42"
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`), // " ` " not " ' " is used to define regex
		// "^/users/(\d+)/?" means look for names such as "/users/<number>"
	} // here we are creating a struct in a local scope and passing a return for that
	// in C/C++, this struct would be cleaned up and hence the address returned would be an invalid one
	// in Go, it automatically detects this kind of scenario, and it will automatically promote the variable to the level that it needs to be in the call stack
	// hence, there is no problem with returning address of local variable in Go
}
