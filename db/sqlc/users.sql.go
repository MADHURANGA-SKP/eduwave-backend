// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package db

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    user_name,
    full_name,
    hashed_password,
    email,
    role,
    qualification
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING user_id, user_name, role, hashed_password, full_name, email, is_email_verified, password_changed_at, created_at, qualification
`

type CreateUserParams struct {
	UserName       string `json:"user_name"`
	FullName       string `json:"full_name"`
	HashedPassword string `json:"hashed_password"`
	Email          string `json:"email"`
	Role           string `json:"role"`
	Qualification  string `json:"qualification"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.UserName,
		arg.FullName,
		arg.HashedPassword,
		arg.Email,
		arg.Role,
		arg.Qualification,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.UserName,
		&i.Role,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.IsEmailVerified,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.Qualification,
	)
	return i, err
}

const deleteUsers = `-- name: DeleteUsers :exec
DELETE FROM users
WHERE user_id = $1
`

func (q *Queries) DeleteUsers(ctx context.Context, userID int64) error {
	_, err := q.db.ExecContext(ctx, deleteUsers, userID)
	return err
}

const getUser = `-- name: GetUser :one


SELECT user_id, user_name, role, hashed_password, full_name, email, is_email_verified, password_changed_at, created_at, qualification FROM users
WHERE user_name = $1 LIMIT 1
`

// -- name: GetUser :one
// SELECT users.user_name AS user_username, teachers.user_name AS teacher_username, admins.user_name AS admin_username
// FROM users
// LEFT JOIN teachers ON users.user_id = teachers.user_id
// LEFT JOIN admins ON users.user_id = admins.user_id
// WHERE
//
//	users.user_name = $1 OR teachers.user_name = $1 OR admins.user_name = $1;
func (q *Queries) GetUser(ctx context.Context, userName string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, userName)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.UserName,
		&i.Role,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.IsEmailVerified,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.Qualification,
	)
	return i, err
}

const listUser = `-- name: ListUser :many
SELECT user_id, user_name, role, hashed_password, full_name, email, is_email_verified, password_changed_at, created_at, qualification FROM users
WHERE role = $1
ORDER BY user_id
LIMIT $2
OFFSET $3
`

type ListUserParams struct {
	Role   string `json:"role"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (q *Queries) ListUser(ctx context.Context, arg ListUserParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUser, arg.Role, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.UserID,
			&i.UserName,
			&i.Role,
			&i.HashedPassword,
			&i.FullName,
			&i.Email,
			&i.IsEmailVerified,
			&i.PasswordChangedAt,
			&i.CreatedAt,
			&i.Qualification,
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

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET 
    hashed_password = COALESCE($1, hashed_password),
    password_changed_at = COALESCE($2, password_changed_at),
    full_name = COALESCE($3, full_name),
    email = COALESCE($4, email),
    is_email_verified = COALESCE($5, is_email_verified)
WHERE
    user_name = $6
RETURNING user_id, user_name, role, hashed_password, full_name, email, is_email_verified, password_changed_at, created_at, qualification
`

type UpdateUserParams struct {
	HashedPassword    string `json:"hashed_password"`
	PasswordChangedAt  time.Time  `json:"password_changed_at"`
	FullName          string `json:"full_name"`
	Email             string `json:"email"`
	IsEmailVerified   bool   `json:"is_email_verified"`
	UserName          string         `json:"user_name"`
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
		&i.UserID,
		&i.UserName,
		&i.Role,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.IsEmailVerified,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.Qualification,
	)
	return i, err
}
