// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: assignments.sql

package db

import (
	"context"
	"time"
)

const createAssignment = `-- name: CreateAssignment :one
INSERT INTO assignments (
    resource_id,
    type,
    title,
    description,
    submission_date
) VALUES (
    $1, $2, $3, $4, $5
)  RETURNING assignment_id, resource_id, type, title, description, submission_date, created_at
`

type CreateAssignmentParams struct {
	ResourceID     int64     `json:"resource_id"`
	Type           string    `json:"type"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	SubmissionDate time.Time `json:"submission_date"`
}

func (q *Queries) CreateAssignment(ctx context.Context, arg CreateAssignmentParams) (Assignment, error) {
	row := q.queryRow(ctx, q.createAssignmentStmt, createAssignment,
		arg.ResourceID,
		arg.Type,
		arg.Title,
		arg.Description,
		arg.SubmissionDate,
	)
	var i Assignment
	err := row.Scan(
		&i.AssignmentID,
		&i.ResourceID,
		&i.Type,
		&i.Title,
		&i.Description,
		&i.SubmissionDate,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAssignment = `-- name: DeleteAssignment :exec
DELETE FROM assignments
WHERE assignment_id = $1
`

func (q *Queries) DeleteAssignment(ctx context.Context, assignmentID int64) error {
	_, err := q.exec(ctx, q.deleteAssignmentStmt, deleteAssignment, assignmentID)
	return err
}

const getAssignment = `-- name: GetAssignment :one
SELECT assignment_id, resource_id, type, title, description, submission_date, created_at FROM assignments
WHERE assignment_id = $1
`

func (q *Queries) GetAssignment(ctx context.Context, assignmentID int64) (Assignment, error) {
	row := q.queryRow(ctx, q.getAssignmentStmt, getAssignment, assignmentID)
	var i Assignment
	err := row.Scan(
		&i.AssignmentID,
		&i.ResourceID,
		&i.Type,
		&i.Title,
		&i.Description,
		&i.SubmissionDate,
		&i.CreatedAt,
	)
	return i, err
}

const getAssignmentByResource = `-- name: GetAssignmentByResource :one
SELECT assignment_id, resource_id, type, title, description, submission_date, created_at FROM assignments
WHERE resource_id = $1
`

func (q *Queries) GetAssignmentByResource(ctx context.Context, resourceID int64) (Assignment, error) {
	row := q.queryRow(ctx, q.getAssignmentByResourceStmt, getAssignmentByResource, resourceID)
	var i Assignment
	err := row.Scan(
		&i.AssignmentID,
		&i.ResourceID,
		&i.Type,
		&i.Title,
		&i.Description,
		&i.SubmissionDate,
		&i.CreatedAt,
	)
	return i, err
}

const updateAssignment = `-- name: UpdateAssignment :one
UPDATE assignments
SET type = $2, title = $3, description = $4, submission_date = $5
WHERE assignment_id = $1
RETURNING assignment_id, resource_id, type, title, description, submission_date, created_at
`

type UpdateAssignmentParams struct {
	AssignmentID   int64     `json:"assignment_id"`
	Type           string    `json:"type"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	SubmissionDate time.Time `json:"submission_date"`
}

func (q *Queries) UpdateAssignment(ctx context.Context, arg UpdateAssignmentParams) (Assignment, error) {
	row := q.queryRow(ctx, q.updateAssignmentStmt, updateAssignment,
		arg.AssignmentID,
		arg.Type,
		arg.Title,
		arg.Description,
		arg.SubmissionDate,
	)
	var i Assignment
	err := row.Scan(
		&i.AssignmentID,
		&i.ResourceID,
		&i.Type,
		&i.Title,
		&i.Description,
		&i.SubmissionDate,
		&i.CreatedAt,
	)
	return i, err
}
