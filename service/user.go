//go:generate mockery --name UserStore --output ./mock --outpkg mock --with-expecter
//go:generate mockery --name RandomUserProvider --output ./mock --outpkg mock --with-expecter

package service

type UserStore interface {
	CreateUser(user *User) error
}

type RandomUserProvider interface {
	GetRandomUser() (*User, error)
}

type User struct {
	store    UserStore
	provider RandomUserProvider
}

func NewUser(store UserStore, provider RandomUserProvider) *User {
	return &User{store: store, provider: provider}
}

func (u *User) FetchAndCreateUser() error {
	ru, err := u.provider.GetRandomUser()
	if err != nil {
		return err
	}

	return u.store.CreateUser(ru)
}
