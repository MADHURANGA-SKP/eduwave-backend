// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: verify_email.sql

package db

import (
	"context"
)

const createVerifyEmail = `-- name: CreateVerifyEmail :one
INSERT INTO verify_emails (
    user_name,
    email,
    secret_code
) VALUES (
    $1, $2, $3
) RETURNING email_id, user_name, email, secret_code, is_used, created_at, expired_at
`

type CreateVerifyEmailParams struct {
	UserName   string `json:"user_name"`
	Email      string `json:"email"`
	SecretCode string `json:"secret_code"`
}

func (q *Queries) CreateVerifyEmail(ctx context.Context, arg CreateVerifyEmailParams) (VerifyEmail, error) {
	row := q.db.QueryRowContext(ctx, createVerifyEmail, arg.UserName, arg.Email, arg.SecretCode)
	var i VerifyEmail
	err := row.Scan(
		&i.EmailID,
		&i.UserName,
		&i.Email,
		&i.SecretCode,
		&i.IsUsed,
		&i.CreatedAt,
		&i.ExpiredAt,
	)
	return i, err
}

const getVerifyEmail = `-- name: GetVerifyEmail :one
SELECT email_id, user_name, email, secret_code, is_used, created_at, expired_at FROM verify_emails
WHERE secret_code = $1 LIMIT 1
`

func (q *Queries) GetVerifyEmail(ctx context.Context, secretCode string) (VerifyEmail, error) {
	row := q.db.QueryRowContext(ctx, getVerifyEmail, secretCode)
	var i VerifyEmail
	err := row.Scan(
		&i.EmailID,
		&i.UserName,
		&i.Email,
		&i.SecretCode,
		&i.IsUsed,
		&i.CreatedAt,
		&i.ExpiredAt,
	)
	return i, err
}

const updateVerifyEmail = `-- name: UpdateVerifyEmail :one
UPDATE verify_emails
SET
    is_used = $1
WHERE
    email_id = $2
    AND secret_code = $3
    AND is_used = FALSE
    AND expired_at > now()
RETURNING email_id, user_name, email, secret_code, is_used, created_at, expired_at
`

type UpdateVerifyEmailParams struct {
	IsUsed     bool   `json:"is_used"`
	EmailID    int64  `json:"email_id"`
	SecretCode string `json:"secret_code"`
}

func (q *Queries) UpdateVerifyEmail(ctx context.Context, arg UpdateVerifyEmailParams) (VerifyEmail, error) {
	row := q.db.QueryRowContext(ctx, updateVerifyEmail, arg.IsUsed, arg.EmailID, arg.SecretCode)
	var i VerifyEmail
	err := row.Scan(
		&i.EmailID,
		&i.UserName,
		&i.Email,
		&i.SecretCode,
		&i.IsUsed,
		&i.CreatedAt,
		&i.ExpiredAt,
	)
	return i, err
}
