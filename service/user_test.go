package service_test

import (
	"context"
	"forbole_code_test/service"
	"forbole_code_test/service/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	ctx = context.Background()
)

func TestCreateUser(t *testing.T) {
	s := newUserStore(t)

	err := s.FetchAndCreateUser(ctx)
	assert.NoError(t, err)
}

func newUserStore(t *testing.T) *service.User {
	mockProvider := &mock.RandomUserProvider{}
	mockStore := &mock.UserStore{}

	mockProvider.EXPECT().GetRandomUser(ctx).Return(&service.User{}, nil)
	mockStore.EXPECT().CreateUser(ctx, &service.User{}).Return(nil)

	return service.NewUser(mockStore, mockProvider)
}
