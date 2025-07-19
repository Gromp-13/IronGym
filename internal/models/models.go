package models

import "time"

type Clients struct {
	ID          int32
	LastName    string
	FirstName   string
	MiddleName  string
	PhoneNumber string
	BirthDate   time.Time
	CardBarcode string
}

type Subscriptions struct {
	ID           int32
	ClientID     int32
	StartDate    time.Time
	DurationDays int32
	EndDate      time.Time
}
