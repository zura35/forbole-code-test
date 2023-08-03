//go:generate mockery --name UserStore --output ./mock --outpkg mock --with-expecter
//go:generate mockery --name RandomUserProvider --output ./mock --outpkg mock --with-expecter

package service

import (
	"context"
	"forbole_code_test/model"
)

type UserStore interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
}

type RandomUserProvider interface {
	GetRandomUser(ctx context.Context) (*model.User, error)
}

type UserService struct {
	store    UserStore
	provider RandomUserProvider
}

func NewUserService(store UserStore, provider RandomUserProvider) *UserService {
	return &UserService{store: store, provider: provider}
}

func (u *UserService) FetchAndCreateUser(ctx context.Context) error {
	ru, err := u.provider.GetRandomUser(ctx)
	if err != nil {
		return err
	}

	_, err = u.store.CreateUser(ctx, ru)
	return err
}
