package repository

import (
	"context"
	"forbole_code_test/model"

	sqlc "forbole_code_test/sqlc_generated"

	"database/sql"
)

type PostgresUserStore struct {
	db *sqlc.Queries
}

func NewPostgresUserStore(db *sql.DB) *PostgresUserStore {
	return &PostgresUserStore{
		db: sqlc.New(db),
	}
}

func (s *PostgresUserStore) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	params := sqlc.CreateUserParams{
		FirstName: sql.NullString{user.FirstName, user.FirstName != ""},
		LastName:  sql.NullString{user.LastName, user.LastName != ""},
		Dob:       sql.NullTime{user.DOB, !user.DOB.IsZero()},
		Address:   sql.NullString{user.Address, user.Address != ""},
	}

	dbUser, err := s.db.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	return &model.User{
		FirstName: dbUser.FirstName.String,
		LastName:  dbUser.LastName.String,
		DOB:       dbUser.Dob.Time,
		Address:   dbUser.Address.String,
	}, nil
}
