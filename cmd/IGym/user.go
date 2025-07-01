package main

import "time"

type User struct {
	User_id    int32
	LastName   string
	FirstName  string
	Patronymic string
	Phone      string
	BirthDate  time.Time
	Barcode    string
}
