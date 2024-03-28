// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: assignments.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createAssignment = `-- name: CreateAssignment :one
INSERT INTO assignments (
    type,
    title,
    description,
    submission_date
) VALUES (
    $1, $2, $3, $4
)  RETURNING assignment_id, course_id, type, title, description, submission_date
`

type CreateAssignmentParams struct {
	Type           string    `json:"type"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	SubmissionDate time.Time `json:"submission_date"`
}

func (q *Queries) CreateAssignment(ctx context.Context, arg CreateAssignmentParams) (Assignment, error) {
	row := q.db.QueryRowContext(ctx, createAssignment,
		arg.Type,
		arg.Title,
		arg.Description,
		arg.SubmissionDate,
	)
	var i Assignment
	err := row.Scan(
		&i.AssignmentID,
		&i.CourseID,
		&i.Type,
		&i.Title,
		&i.Description,
		&i.SubmissionDate,
	)
	return i, err
}

const deleteAssignment = `-- name: DeleteAssignment :exec
DELETE FROM assignments
WHERE assignment_id = $1
`

func (q *Queries) DeleteAssignment(ctx context.Context, assignmentID int64) error {
	_, err := q.db.ExecContext(ctx, deleteAssignment, assignmentID)
	return err
}

const getAssignment = `-- name: GetAssignment :one
SELECT assignment_id, course_id, type, title, description, submission_date FROM assignments
WHERE course_id = $1
`

func (q *Queries) GetAssignment(ctx context.Context, courseID sql.NullInt64) (Assignment, error) {
	row := q.db.QueryRowContext(ctx, getAssignment, courseID)
	var i Assignment
	err := row.Scan(
		&i.AssignmentID,
		&i.CourseID,
		&i.Type,
		&i.Title,
		&i.Description,
		&i.SubmissionDate,
	)
	return i, err
}

const updateAssignment = `-- name: UpdateAssignment :one
UPDATE assignments
SET type = $2, title = $3, description = $4, submission_date = $5
WHERE course_id = $1
RETURNING assignment_id, course_id, type, title, description, submission_date
`

type UpdateAssignmentParams struct {
	CourseID       sql.NullInt64 `json:"course_id"`
	Type           string        `json:"type"`
	Title          string        `json:"title"`
	Description    string        `json:"description"`
	SubmissionDate time.Time     `json:"submission_date"`
}

func (q *Queries) UpdateAssignment(ctx context.Context, arg UpdateAssignmentParams) (Assignment, error) {
	row := q.db.QueryRowContext(ctx, updateAssignment,
		arg.CourseID,
		arg.Type,
		arg.Title,
		arg.Description,
		arg.SubmissionDate,
	)
	var i Assignment
	err := row.Scan(
		&i.AssignmentID,
		&i.CourseID,
		&i.Type,
		&i.Title,
		&i.Description,
		&i.SubmissionDate,
	)
	return i, err
}