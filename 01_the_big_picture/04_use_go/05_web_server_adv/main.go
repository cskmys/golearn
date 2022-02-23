// here we update our simple web server to receive information on the url via a query parameter and use that to change the response
// also now we return the response in the json format which is used by a lot of web services instead of a string format

// now while in web browser to provide query parameter by name "name" type "localhost:3000/?name=<name>"

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) { // queries are read from "http.Request" object
		names := request.URL.Query()["name"] // "http.Request::URL::Query" gives a map(a collection of key-value pairs) of queries,
		// "["name"]" is used to read all the query parameters called "name"
		// hence, "names" is potentially a collection having more than one name
		var name string
		if len(names) == 1 {
			name = names[0]
		}
		m := map[string]string{"name": name} // a string-string map which can easily be serialized into json data string
		enc := json.NewEncoder(writer)       // "json::NewEncoder" to encode the map into a json string, the output stream of encoder set to "writer"
		enc.Encode(m)                        // once you call this the map "m" is serialized into a json string by "enc" and written into "writer"
	})

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
