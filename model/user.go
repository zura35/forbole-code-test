package model

import "time"

type User struct {
	FirstName string
	LastName  string
	DOB       time.Time
	Address   *Address
}

type Address struct {
	City          string       `json:"city"`
	StreetName    string       `json:"street_name"`
	StreetAddress string       `json:"street_address"`
	ZipCode       string       `json:"zip_code"`
	State         string       `json:"state"`
	Country       string       `json:"country"`
	Coordinates   *Coordinates `json:"coordinates"`
}

type Coordinates struct {
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lng"`
}
