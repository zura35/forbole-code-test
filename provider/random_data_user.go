package provider

import (
	"context"
	"forbole_code_test/model"
	"io"
	"net/http"
)

type UserProvider struct {
	dsu DataSourceUser
}

// DataSourceUser defines the source of data and the data transformation methods
type DataSourceUser interface {
	// Source returns the data source URL
	Source() string

	// FromJSON transforms JSON response to a corresponding struct
	FromJSON(data []byte) error

	// ToUserModel transforms the struct to the User model
	ToUserModel() (*model.User, error)
}

func NewUserProvider(dsu DataSourceUser) *UserProvider {
	return &UserProvider{dsu: dsu}
}

func (r *UserProvider) GetRandomUser(ctx context.Context) (*model.User, error) {
	res, err := http.Get(r.dsu.Source())
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = r.dsu.FromJSON(body)
	if err != nil {
		return nil, err
	}

	return r.dsu.ToUserModel()
}
