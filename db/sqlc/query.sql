-- name: CreateUser :one
INSERT INTO users (
    first_name,
    last_name,
    dob,
    address
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;