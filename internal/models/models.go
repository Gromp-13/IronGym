package models

import "time"

type Client struct {
	ID          int32
	LastName    string
	FirstName   string
	MiddleName  string
	PhoneNumber string
	BirthDate   time.Time
	CardBarcode string
	GenderID    int32
}

type Subscriptions struct {
	ID           int32
	ClientID     int32
	StartDate    time.Time
	DurationDays int32
	EndDate      time.Time
}

type Gender struct {
	ID          int32
	Description string
}
