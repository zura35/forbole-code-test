package service_test

import (
	"forbole_code_test/service"
	"forbole_code_test/service/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	s := newUserStore(t)

	err := s.FetchAndCreateUser()
	assert.NoError(t, err)
}

func newUserStore(t *testing.T) *service.User {
	mockProvider := &mock.RandomUserProvider{}
	mockStore := &mock.UserStore{}

	mockProvider.EXPECT().GetRandomUser().Return(&service.User{}, nil)
	mockStore.EXPECT().CreateUser(&service.User{}).Return(nil)

	return service.NewUser(mockStore, mockProvider)
}
