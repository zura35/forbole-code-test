package repository

import (
	"context"
	"encoding/json"
	"forbole_code_test/model"

	sqlc "forbole_code_test/sqlc_generated"

	pgtype "github.com/sqlc-dev/pqtype"

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
	addrJSON, err := json.Marshal(user.Address)
	if err != nil {
		return nil, err
	}

	params := sqlc.CreateUserParams{
		FirstName: sql.NullString{user.FirstName, user.FirstName != ""},
		LastName:  sql.NullString{user.LastName, user.LastName != ""},
		Dob:       sql.NullTime{user.DOB, !user.DOB.IsZero()},
		Address:   pgtype.NullRawMessage{addrJSON, &user.Address != nil},
	}

	dbUser, err := s.db.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	var address *model.Address
	if err := json.Unmarshal(dbUser.Address.RawMessage, &address); err != nil {
		return nil, err
	}

	return &model.User{
		FirstName: dbUser.FirstName.String,
		LastName:  dbUser.LastName.String,
		DOB:       dbUser.Dob.Time,
		Address:   address,
	}, nil
}
