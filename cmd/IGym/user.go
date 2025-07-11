package main

import "time"

type User struct {
	User_id     int32
	LastName    string
	FirstName   string
	MiddleName  string
	PhoneNumber string
	BirthDate   time.Time
	CardBarcode string
}
