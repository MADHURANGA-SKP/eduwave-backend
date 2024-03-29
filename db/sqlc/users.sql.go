// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    user_name,
    role,
    full_name,
    hashed_password,
    email
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING user_name, role, hashed_password, full_name, email, is_email_verified, password_changed_at, created_at
`

type CreateUserParams struct {
	UserName       string   `json:"user_name"`
	Role           UserRole `json:"role"`
	FullName       string   `json:"full_name"`
	HashedPassword string   `json:"hashed_password"`
	Email          string   `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.UserName,
		arg.Role,
		arg.FullName,
		arg.HashedPassword,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.UserName,
		&i.Role,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.IsEmailVerified,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT user_name, role, hashed_password, full_name, email, is_email_verified, password_changed_at, created_at FROM users
WHERE user_name = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, userName string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, userName)
	var i User
	err := row.Scan(
		&i.UserName,
		&i.Role,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.IsEmailVerified,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET 
    hashed_password = COALESCE($1, hashed_password),
    password_changed_at = COALESCE($2, password_changed_at),
    full_name = COALESCE($3, full_name),
    email = COALESCE($4, email),
    is_email_verified = COALESCE($5, is_email_verified)
WHERE
    user_name = COALESCE($6, user_name)
RETURNING user_name, role, hashed_password, full_name, email, is_email_verified, password_changed_at, created_at
`

type UpdateUserParams struct {
	HashedPassword    sql.NullString `json:"hashed_password"`
	PasswordChangedAt sql.NullTime   `json:"password_changed_at"`
	FullName          sql.NullString `json:"full_name"`
	Email             sql.NullString `json:"email"`
	IsEmailVerified   sql.NullBool   `json:"is_email_verified"`
	UserName          sql.NullString `json:"user_name"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.HashedPassword,
		arg.PasswordChangedAt,
		arg.FullName,
		arg.Email,
		arg.IsEmailVerified,
		arg.UserName,
	)
	var i User
	err := row.Scan(
		&i.UserName,
		&i.Role,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.IsEmailVerified,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}
