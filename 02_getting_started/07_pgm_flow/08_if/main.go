package main

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func main() {
	u1 := User{
		ID:        1,
		FirstName: "Arthur",
		LastName:  "Dent",
	}
	u2 := User{
		ID:        2,
		FirstName: "Ford",
		LastName:  "Prefect",
	}
	if u1.ID == u2.ID {
		println("Same user")
	}

	if u1.ID != u2.ID {
		println("Different user")
	}

	if u1.FirstName == u2.FirstName {
		println("Same first name")
	} else if u1.LastName == u2.LastName { // "else if" is optional, you can have just "if" or "if-else"
		println("Same last user")
	} else { // "else" is optional, you can have just "if" or "if-else if"
		println("Very different names")
	}

}
