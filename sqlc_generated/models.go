// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package sqlc_generated

import (
	"database/sql"
)

type SchemaMigration struct {
	Version int64
	Dirty   bool
}

type User struct {
	ID        int32
	FirstName sql.NullString
	LastName  sql.NullString
	Dob       sql.NullTime
	Address   sql.NullString
}