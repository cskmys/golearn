// here we are going to start our web server
// run the webserver, go to web browser and enter "http://localhost:3000/users" or "http://localhost:3000/users/<number>"
package main

import (
	"09_demo_startws/controllers"
	"net/http"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil) // ":3000" means listen of port 3000 of localhost and "nil" tells library to use the default/in-built handler to handle all the requests at the front
	// here the default handler is the front controller and the "ServeHTTP" we implemented in "controllers/user.go" is the back controller
	// read up front controller back controller pattern for more information
}
