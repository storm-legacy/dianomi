// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package sqlc

import (
	"context"
	"time"
)

const addRevokedToken = `-- name: AddRevokedToken :exec
INSERT INTO revoked_tokens (token, user_id, valid_until) VALUES ($1, $2, $3)
`

type AddRevokedTokenParams struct {
	Token      string
	UserID     int64
	ValidUntil time.Time
}

func (q *Queries) AddRevokedToken(ctx context.Context, arg AddRevokedTokenParams) error {
	_, err := q.db.ExecContext(ctx, addRevokedToken, arg.Token, arg.UserID, arg.ValidUntil)
	return err
}

const checkToken = `-- name: CheckToken :one
SELECT id, token, user_id, valid_until FROM revoked_tokens WHERE token = $1 LIMIT 1
`

func (q *Queries) CheckToken(ctx context.Context, token string) (RevokedToken, error) {
	row := q.db.QueryRowContext(ctx, checkToken, token)
	var i RevokedToken
	err := row.Scan(
		&i.ID,
		&i.Token,
		&i.UserID,
		&i.ValidUntil,
	)
	return i, err
}

const createUser = `-- name: CreateUser :exec
INSERT INTO users (email, password) VALUES ($1, $2)
`

type CreateUserParams struct {
	Email    string
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser, arg.Email, arg.Password)
	return err
}

const getPackagesByUserID = `-- name: GetPackagesByUserID :many
SELECT id, user_id, tier, created_at, valid_from, valid_until FROM
  users_packages
WHERE
  user_id = $1
  AND
  (now()::timestamp with TIME ZONE) BETWEEN valid_from AND valid_until
ORDER BY
  created_at DESC
LIMIT 10
`

func (q *Queries) GetPackagesByUserID(ctx context.Context, userID int64) ([]UsersPackage, error) {
	rows, err := q.db.QueryContext(ctx, getPackagesByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UsersPackage
	for rows.Next() {
		var i UsersPackage
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Tier,
			&i.CreatedAt,
			&i.ValidFrom,
			&i.ValidUntil,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, password, verified_at, created_at, updated_at FROM users WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.VerifiedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, email, password, verified_at, created_at, updated_at FROM users WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.VerifiedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
