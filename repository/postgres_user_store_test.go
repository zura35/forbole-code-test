package repository_test

import (
	"context"
	"database/sql"
	"fmt"
	"forbole_code_test/model"
	"forbole_code_test/repository"
	"testing"
	"time"

	_ "github.com/lib/pq"

	"github.com/stretchr/testify/assert"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

var (
	store *repository.PostgresUserStore
)

func TestCreateUser(t *testing.T) {
	s := newPostgresUserStore(t)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	now := time.Now()
	u, err := s.CreateUser(ctx, &model.User{
		FirstName: "John",
		LastName:  "Doe",
		DOB:       now,
		Address: &model.Address{
			StreetName: "123 Main St",
			City:       "New York",
			State:      "NY",
		},
	})

	dateNow, _ := time.Parse("2006-01-02", now.Format("2006-01-02"))

	assert.NoError(t, err)
	assert.Equal(t, "John", u.FirstName)
	assert.Equal(t, "Doe", u.LastName)
	assert.Equal(t, dateNow.Compare(u.DOB), 0)

	assert.Equal(t, "123 Main St", u.Address.StreetName)
	assert.Equal(t, "New York", u.Address.City)
	assert.Equal(t, "NY", u.Address.State)
}

func newPostgresUserStore(t *testing.T) *repository.PostgresUserStore {
	if store != nil {
		return store
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatal(err)
	}

	store = repository.NewPostgresUserStore(db)
	return store
}
