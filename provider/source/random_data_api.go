package source

import (
	"encoding/json"
	"forbole_code_test/model"
	"time"
)

type RandomDataAPIUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	DOB       string `json:"date_of_birth"`
	Address   struct {
		City          string `json:"city"`
		StreetName    string `json:"street_name"`
		StreetAddress string `json:"street_address"`
		ZipCode       string `json:"zip_code"`
		State         string `json:"state"`
		Country       string `json:"country"`
		Coordinates   struct {
			Latitude  float32 `json:"lat"`
			Longitude float32 `json:"lng"`
		}
	} `json:"address"`
}

func NewRandomDataAPIUser() *RandomDataAPIUser {
	return &RandomDataAPIUser{}
}

func (u *RandomDataAPIUser) Source() string {
	return "https://random-data-api.com/api/users/random_user"
}

func (u *RandomDataAPIUser) FromJSON(data []byte) error {
	return json.Unmarshal(data, u)
}

func (u *RandomDataAPIUser) ToUserModel() (*model.User, error) {
	dob, err := time.Parse(time.DateOnly, u.DOB)
	if err != nil {
		return nil, err
	}

	return &model.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		DOB:       dob,
		Address: &model.Address{
			City:          u.Address.City,
			StreetName:    u.Address.StreetName,
			StreetAddress: u.Address.StreetAddress,
			ZipCode:       u.Address.ZipCode,
			State:         u.Address.State,
			Country:       u.Address.Country,
			Coordinates: &model.Coordinates{
				Latitude:  u.Address.Coordinates.Latitude,
				Longitude: u.Address.Coordinates.Longitude,
			},
		},
	}, nil
}
