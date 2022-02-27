package main

import (
	"08_demo_intf/models"
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
