package provider

import (
	"context"
	"encoding/json"
	"forbole_code_test/model"
	"io"
	"net/http"
	"time"
)

type RandomDataUser struct{}

type userJSON struct {
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

const (
	url = "https://random-data-api.com/api/users/random_user"
)

func NewRandomDataUser() *RandomDataUser {
	return &RandomDataUser{}
}

func (r *RandomDataUser) GetRandomUser(ctx context.Context) (*model.User, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var uJSON userJSON
	err = json.Unmarshal(body, &uJSON)
	if err != nil {
		return nil, err
	}

	dob, err := time.Parse(time.DateOnly, uJSON.DOB)
	if err != nil {
		return nil, err
	}

	return &model.User{
		FirstName: uJSON.FirstName,
		LastName:  uJSON.LastName,
		DOB:       dob,
		Address: &model.Address{
			City:          uJSON.Address.City,
			StreetName:    uJSON.Address.StreetName,
			StreetAddress: uJSON.Address.StreetAddress,
			ZipCode:       uJSON.Address.ZipCode,
			State:         uJSON.Address.State,
			Country:       uJSON.Address.Country,
			Coordinates: &model.Coordinates{
				Latitude:  uJSON.Address.Coordinates.Latitude,
				Longitude: uJSON.Address.Coordinates.Longitude,
			},
		},
	}, nil

}
