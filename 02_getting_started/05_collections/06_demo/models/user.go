package models // it is not "package main" coz "main" is the entry point of the program and this file is not the entry point

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var ( // just like "import" and "const" can have their own blocks, variables can have their own block too
	users []*User // creating a slice to an underlying array consisting of pointers to data of type "User"
	// pointers are used to avoid copy
	nextID = 1 // by providing initial value, compiler can infer the type

)
