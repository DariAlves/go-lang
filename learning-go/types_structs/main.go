package main

import (
	"log"
	"time"
)

// var FirstName string
// var LastName string
// var PhoneNumber string
// var Age int
// var BirthDate time.Time

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Bruce",
		LastName:    "Wayne",
		PhoneNumber: "0000-0000",
	}

	log.Println(user.FirstName, user.LastName, "BirthDate: ", user.BirthDate)
}
