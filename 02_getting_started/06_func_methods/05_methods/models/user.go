package models

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

func AddUser(u User) (User, error) { // in Go, you always return error if something can go wrong and let the caller decide how to proceed
	u.ID = nextID
	nextID++
	users = append(users, &u) // "users" take pointer, hence we pass "&u" i.e. address of "u"
	return u, nil
}
