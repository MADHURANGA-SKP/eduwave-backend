// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: matirials.sql

package db

import (
	"context"
	"database/sql"
)

const createMatirials = `-- name: CreateMatirials :one
INSERT INTO matirials (
    title,
    description
) VALUES (
    $1, $2
) RETURNING matirial_id, course_id, title, description, created_at
`

type CreateMatirialsParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (q *Queries) CreateMatirials(ctx context.Context, arg CreateMatirialsParams) (Matirial, error) {
	row := q.db.QueryRowContext(ctx, createMatirials, arg.Title, arg.Description)
	var i Matirial
	err := row.Scan(
		&i.MatirialID,
		&i.CourseID,
		&i.Title,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const deleteMatirials = `-- name: DeleteMatirials :exec
DELETE FROM matirials
WHERE matirial_id = $1 AND course_id = $2
`

type DeleteMatirialsParams struct {
	MatirialID int64         `json:"matirial_id"`
	CourseID   sql.NullInt64 `json:"course_id"`
}

func (q *Queries) DeleteMatirials(ctx context.Context, arg DeleteMatirialsParams) error {
	_, err := q.db.ExecContext(ctx, deleteMatirials, arg.MatirialID, arg.CourseID)
	return err
}

const getMatirials = `-- name: GetMatirials :one
SELECT matirial_id, course_id, title, description, created_at FROM matirials
WHERE course_id = $1
`

func (q *Queries) GetMatirials(ctx context.Context, courseID sql.NullInt64) (Matirial, error) {
	row := q.db.QueryRowContext(ctx, getMatirials, courseID)
	var i Matirial
	err := row.Scan(
		&i.MatirialID,
		&i.CourseID,
		&i.Title,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const listMatirials = `-- name: ListMatirials :many
SELECT matirial_id, course_id, title, description, created_at FROM matirials
WHERE course_id = $1
ORDER BY matirial_id
LIMIT $2
OFFSET $3
`

type ListMatirialsParams struct {
	CourseID sql.NullInt64 `json:"course_id"`
	Limit    int32         `json:"limit"`
	Offset   int32         `json:"offset"`
}

func (q *Queries) ListMatirials(ctx context.Context, arg ListMatirialsParams) ([]Matirial, error) {
	rows, err := q.db.QueryContext(ctx, listMatirials, arg.CourseID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Matirial{}
	for rows.Next() {
		var i Matirial
		if err := rows.Scan(
			&i.MatirialID,
			&i.CourseID,
			&i.Title,
			&i.Description,
			&i.CreatedAt,
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

const updateMatirials = `-- name: UpdateMatirials :one
UPDATE matirials
SET title = $3, description = $4
WHERE matirial_id = $1 AND course_id = $2
RETURNING matirial_id, course_id, title, description, created_at
`

type UpdateMatirialsParams struct {
	MatirialID  int64         `json:"matirial_id"`
	CourseID    sql.NullInt64 `json:"course_id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
}

func (q *Queries) UpdateMatirials(ctx context.Context, arg UpdateMatirialsParams) (Matirial, error) {
	row := q.db.QueryRowContext(ctx, updateMatirials,
		arg.MatirialID,
		arg.CourseID,
		arg.Title,
		arg.Description,
	)
	var i Matirial
	err := row.Scan(
		&i.MatirialID,
		&i.CourseID,
		&i.Title,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}