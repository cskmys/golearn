// in Go,
// a module is a directory that contains the "go.mod" file
// a module can contain several packages
// a package is a discrete unit of code which is related to an aspect of the module

// here we start to build a web service

package main

import (
	"06_demo/models" // to import package is located from the top level module
	"fmt"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Sid",
		LastName:  "Stark",
	}
	fmt.Println(u)
}
