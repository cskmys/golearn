package main

import (
	"06_demo_mkmethod/models"
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
