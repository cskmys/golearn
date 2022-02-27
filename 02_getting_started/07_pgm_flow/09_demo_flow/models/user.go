package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("new user must not include id or it must be set to zero")
		// when there is an error, and you are expected to return an object convention is to return an empty object
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

func GetUserById(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil // "nil" was returned then "u" is valid, and shouldn't be ignored
		}
	}
	return User{}, fmt.Errorf("user with id %v not found", id) // "fmt.Errorf" is same as "errors.New" but with an ability to produce a formatted string
}

func UpdateUser(u User) (User, error) {
	for i, user := range users {
		if user.ID == u.ID {
			users[i] = &u // replacing by just making the pointer to point to a new address
			return u, nil
		}
	}
	return User{}, fmt.Errorf("user with id %v not found", u.ID)
}

func RemoveUserById(id int) error {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...) // splice operation on the slice
			// "[:i]" is all users upto but including the one at index "i", and "[i+1:]" means from user at index "i + 1" all the way to the end
			// no idea what "..." here means
			return nil
		}
	}
	return fmt.Errorf("user with id %v not found", id)
}
