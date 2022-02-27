// this will handle all the rough routing
// when a network request is received it will route it to proper controller to process it

package controllers

import "net/http"

func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc) // process all the requests on "/users" route and new user controller is passed as
	// it is the one which has the implementation of the Handle's interface function "ServeHTTP"
	http.Handle("/users/", *uc) // in Go, "/users" and "/users/" are different routes, hence you need to add this as well
	// "/users/" means anything that comes after "/users/" meaning "/users/5" will be routed to this handler
}
