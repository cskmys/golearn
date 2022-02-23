// to just build a go application, you can do:
// go build .

// in Go,
// you can write anonymous functions on the spot while passing it as an argument, without any pre-definition
// in functions we first give the variable name and then the type(opposite of C/C++)

// here we create a simple web server that returns a static text

// here once you run the web server from either "go run ." or "./04_web_server" in the temp directory "/tmp/go-build<.....>" where you built the application using "go build ."
// go to web browser and type "localhost:3000" to see the static text
// use Ctrl-C in the terminal to close the web server

package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// "http.ResponseWriter" is a response stream, and writers are output stream
		// "http.Request" is the request object from which you can get the url of the request, header information and request body
		writer.Write([]byte("Hello World")) // writers in Go work with collection of bytes and "[]byte(...)" converts the string inside into a collection of bytes
	}) // "HandleFunc" responds to HTTP requests with a function call
	// here it is going to listen to all the requests that comes into the application on "/" i.e. root path
	// whenever a request arrives it will use the function parameter to hande it

	err := http.ListenAndServe(":3000", nil) // starts the webserver that listens on port 3000 with a default handler
	// here passing "nil" makes the api use the default handler
	if err != nil {
		log.Fatal(err)
	}
}
