package model

import "time"

type User struct {
	FirstName string
	LastName  string
	DOB       time.Time
	Address   string
}
