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
	Gender      int32
	StatusID    int32
}

type Gender struct {
	ID          int32
	Description string
}

type ClientStatus struct {
	ID   int32
	Name string
}

type Membership struct {
	ID           int32
	Name         string
	Price        int32
	DurationDays int32
}

type ClientMembership struct {
	ID           int32
	ClientID     int32
	MembershipID int32
	StartDate    time.Time
	PaidAmount   int32
	CreatedAT    time.Time
	StatusID     int32
}

type MembershipStatus struct {
	ID   int32
	Name string
}

type Payment struct {
	ID           int32
	ClientID     int32
	MembershipID int32
	Amount       int32
	PaymentDate  time.Time
	Method       string
}
