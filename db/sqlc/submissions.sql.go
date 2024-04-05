// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: submissions.sql

package db

import (
	"context"
)

const getsubmissions = `-- name: Getsubmissions :one
SELECT submission_id, assignment_id, user_id FROM submissions
WHERE assignment_id = $1 AND user_id = $2 
LIMIT 1
`

type GetsubmissionsParams struct {
	AssignmentID int64 `json:"assignment_id"`
	UserID       int64 `json:"user_id"`
}

func (q *Queries) Getsubmissions(ctx context.Context, arg GetsubmissionsParams) (Submission, error) {
	row := q.db.QueryRowContext(ctx, getsubmissions, arg.AssignmentID, arg.UserID)
	var i Submission
	err := row.Scan(&i.SubmissionID, &i.AssignmentID, &i.UserID)
	return i, err
}

const listsubmissions = `-- name: Listsubmissions :many
SELECT submission_id, assignment_id, user_id FROM submissions
WHERE assignment_id = $1 AND user_id = $2
ORDER BY submission_id
LIMIT $2
OFFSET $3
`

type ListsubmissionsParams struct {
	AssignmentID int64 `json:"assignment_id"`
	Limit        int32 `json:"limit"`
	Offset       int32 `json:"offset"`
}

func (q *Queries) Listsubmissions(ctx context.Context, arg ListsubmissionsParams) ([]Submission, error) {
	rows, err := q.db.QueryContext(ctx, listsubmissions, arg.AssignmentID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Submission{}
	for rows.Next() {
		var i Submission
		if err := rows.Scan(&i.SubmissionID, &i.AssignmentID, &i.UserID); err != nil {
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
