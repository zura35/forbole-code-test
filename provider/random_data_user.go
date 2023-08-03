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

type DataSourceUser interface {
	Source() string
	FromJSON(data []byte) error
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
