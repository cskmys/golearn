package main

type HTTPRequest struct {
	Method string
}

func main() {
	r := HTTPRequest{Method: "POST"}

	switch r.Method { // much cleaner than an "if-(else if)-else" ladder
	case "GET":
		println("Get request") // unlike C/C++ there is no need for a "break" statement
	case "DELETE":
		println("Delete request")
		break // "break" won't throw an error, but it is not needed as it is implicitly added
	case "POST":
		println("Post request")
		fallthrough // "fallthrough" overrides implicit "break" to cause a fallthrough until a "break"(either implicit or explicit) is encountered
	case "PUT":
		println("Put request")
	default: // optional
		println("Unknown request")
	}

}
