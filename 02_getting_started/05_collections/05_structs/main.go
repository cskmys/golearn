// in Go,
// struct fields are fixed at a compile time

package main

import "fmt"

func main() {
	// struct definition is a 2-step process
	// first you'll need to declare the struct type
	// then create an instance of that type
	{
		type user struct {
			ID        int
			FirstName string
			LastName  string
		}
		var u user // uninitialized struct will be initialized with default values at all its fields
		fmt.Println(u)
		u.ID = 1
		u.FirstName = "Sid"
		u.LastName = "Stark"
		fmt.Println(u)
		fmt.Println(u.LastName)

		u2 := user{
			ID:        2,
			FirstName: "Cho",
			LastName:  "Chaud", // if you forget to add "," at the end, the auto semicolon insertion will add a semicolon and causes a compilation error
		}
		fmt.Println(u2)
		fmt.Println(u2.LastName)
	}
	/*
		var u3 user // throws compilation error
		fmt.Println(u3)
		// struct "user" is accessible only within the scope that it was defined in
		// hence "user" is unknown in this scope
	*/
}
