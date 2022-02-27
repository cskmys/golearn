package main

import (
	"05_methods/models"
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
