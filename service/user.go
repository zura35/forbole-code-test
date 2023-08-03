//go:generate mockery --name UserStore --output ./mock --outpkg mock --with-expecter
//go:generate mockery --name RandomUserProvider --output ./mock --outpkg mock --with-expecter

package service

import "context"

type UserStore interface {
	CreateUser(ctx context.Context, user *User) error
}

type RandomUserProvider interface {
	GetRandomUser(ctx context.Context) (*User, error)
}

type User struct {
	store    UserStore
	provider RandomUserProvider
}

func NewUser(store UserStore, provider RandomUserProvider) *User {
	return &User{store: store, provider: provider}
}

func (u *User) FetchAndCreateUser(ctx context.Context) error {
	ru, err := u.provider.GetRandomUser(ctx)
	if err != nil {
		return err
	}

	return u.store.CreateUser(ctx, ru)
}
