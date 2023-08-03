package provider_test

import (
	"context"
	"forbole_code_test/provider"
	"forbole_code_test/provider/source"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetRandomUser__Integration(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userProvider := provider.NewUserProvider(source.NewRandomDataAPIUser())
	user, err := userProvider.GetRandomUser(ctx)
	assert.NoError(t, err)

	assert.NotEmpty(t, user.FirstName)
	assert.NotEmpty(t, user.LastName)
	assert.Equal(t, user.DOB.IsZero(), false)
	assert.NotEmpty(t, user.Address)
}
