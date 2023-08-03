//go:generate mockery --name UserStore --output ./mock --outpkg mock --case underscore --with-expecter
//go:generate mockery --name UserProvider --output ./mock --outpkg mock --case underscore --with-expecter

package service

import (
	"context"
	"forbole_code_test/model"
)

// UserStore defines the data store access methods used in UserService
type UserStore interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
}

// UserProvider defines the external API access methods used in UserService
type UserProvider interface {
	GetRandomUser(ctx context.Context) (*model.User, error)
}

type UserService struct {
	store    UserStore
	provider UserProvider
}

func NewUserService(store UserStore, provider UserProvider) *UserService {
	return &UserService{store: store, provider: provider}
}

// FetchAndCreateUser fetches a random user from UserProvider and inserts it to UserStore
func (u *UserService) FetchAndCreateUser(ctx context.Context) error {
	ru, err := u.provider.GetRandomUser(ctx)
	if err != nil {
		return err
	}

	_, err = u.store.CreateUser(ctx, ru)
	return err
}
