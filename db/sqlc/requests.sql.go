// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: requests.sql

package db

import (
	"context"
	"database/sql"
)

const createRequest = `-- name: CreateRequest :one
INSERT INTO requests (
    user_id,
    course_id,
    is_active,
    is_pending,
    is_accepted,
    is_declined
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING request_id, user_id, course_id, is_active, is_pending, is_accepted, is_declined, created_at
`

type CreateRequestParams struct {
	UserID     int64        `json:"user_id"`
	CourseID   int64        `json:"course_id"`
	IsActive   sql.NullBool `json:"is_active"`
	IsPending  sql.NullBool `json:"is_pending"`
	IsAccepted sql.NullBool `json:"is_accepted"`
	IsDeclined sql.NullBool `json:"is_declined"`
}

func (q *Queries) CreateRequest(ctx context.Context, arg CreateRequestParams) (Request, error) {
	row := q.queryRow(ctx, q.createRequestStmt, createRequest,
		arg.UserID,
		arg.CourseID,
		arg.IsActive,
		arg.IsPending,
		arg.IsAccepted,
		arg.IsDeclined,
	)
	var i Request
	err := row.Scan(
		&i.RequestID,
		&i.UserID,
		&i.CourseID,
		&i.IsActive,
		&i.IsPending,
		&i.IsAccepted,
		&i.IsDeclined,
		&i.CreatedAt,
	)
	return i, err
}

const deleteRequest = `-- name: DeleteRequest :exec
DELETE FROM requests
WHERE request_id = $1
`

func (q *Queries) DeleteRequest(ctx context.Context, requestID int64) error {
	_, err := q.exec(ctx, q.deleteRequestStmt, deleteRequest, requestID)
	return err
}

const getRequest = `-- name: GetRequest :one
SELECT request_id, user_id, course_id, is_active, is_pending, is_accepted, is_declined, created_at FROM requests
WHERE user_id = $1 AND course_id = $2
`

type GetRequestParams struct {
	UserID   int64 `json:"user_id"`
	CourseID int64 `json:"course_id"`
}

func (q *Queries) GetRequest(ctx context.Context, arg GetRequestParams) (Request, error) {
	row := q.queryRow(ctx, q.getRequestStmt, getRequest, arg.UserID, arg.CourseID)
	var i Request
	err := row.Scan(
		&i.RequestID,
		&i.UserID,
		&i.CourseID,
		&i.IsActive,
		&i.IsPending,
		&i.IsAccepted,
		&i.IsDeclined,
		&i.CreatedAt,
	)
	return i, err
}

const listRequest = `-- name: ListRequest :many
SELECT request_id, user_id, course_id, is_active, is_pending, is_accepted, is_declined, created_at FROM requests
ORDER BY request_id
LIMIT $1
OFFSET $2
`

type ListRequestParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListRequest(ctx context.Context, arg ListRequestParams) ([]Request, error) {
	rows, err := q.query(ctx, q.listRequestStmt, listRequest, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Request{}
	for rows.Next() {
		var i Request
		if err := rows.Scan(
			&i.RequestID,
			&i.UserID,
			&i.CourseID,
			&i.IsActive,
			&i.IsPending,
			&i.IsAccepted,
			&i.IsDeclined,
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

const listRequestByCourse = `-- name: ListRequestByCourse :many
SELECT request_id, user_id, course_id, is_active, is_pending, is_accepted, is_declined, created_at FROM requests
WHERE course_id = $1
ORDER BY request_id
LIMIT $2
OFFSET $3
`

type ListRequestByCourseParams struct {
	CourseID int64 `json:"course_id"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

func (q *Queries) ListRequestByCourse(ctx context.Context, arg ListRequestByCourseParams) ([]Request, error) {
	rows, err := q.query(ctx, q.listRequestByCourseStmt, listRequestByCourse, arg.CourseID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Request{}
	for rows.Next() {
		var i Request
		if err := rows.Scan(
			&i.RequestID,
			&i.UserID,
			&i.CourseID,
			&i.IsActive,
			&i.IsPending,
			&i.IsAccepted,
			&i.IsDeclined,
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

const listRequestByUser = `-- name: ListRequestByUser :many
SELECT request_id, user_id, course_id, is_active, is_pending, is_accepted, is_declined, created_at FROM requests
WHERE user_id = $1
ORDER BY request_id
LIMIT $2
OFFSET $3
`

type ListRequestByUserParams struct {
	UserID int64 `json:"user_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListRequestByUser(ctx context.Context, arg ListRequestByUserParams) ([]Request, error) {
	rows, err := q.query(ctx, q.listRequestByUserStmt, listRequestByUser, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Request{}
	for rows.Next() {
		var i Request
		if err := rows.Scan(
			&i.RequestID,
			&i.UserID,
			&i.CourseID,
			&i.IsActive,
			&i.IsPending,
			&i.IsAccepted,
			&i.IsDeclined,
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

const updateRequests = `-- name: UpdateRequests :one
UPDATE requests
SET is_active = $3, is_pending = $4, is_accepted = $5, is_declined = $6 
WHERE user_id = $1 AND course_id = $2
RETURNING request_id, user_id, course_id, is_active, is_pending, is_accepted, is_declined, created_at
`

type UpdateRequestsParams struct {
	UserID     int64        `json:"user_id"`
	CourseID   int64        `json:"course_id"`
	IsActive   bool `json:"is_active"`
	IsPending  bool `json:"is_pending"`
	IsAccepted bool `json:"is_accepted"`
	IsDeclined bool `json:"is_declined"`
}

func (q *Queries) UpdateRequests(ctx context.Context, arg UpdateRequestsParams) (Request, error) {
	row := q.queryRow(ctx, q.updateRequestsStmt, updateRequests,
		arg.UserID,
		arg.CourseID,
		arg.IsActive,
		arg.IsPending,
		arg.IsAccepted,
		arg.IsDeclined,
	)
	var i Request
	err := row.Scan(
		&i.RequestID,
		&i.UserID,
		&i.CourseID,
		&i.IsActive,
		&i.IsPending,
		&i.IsAccepted,
		&i.IsDeclined,
		&i.CreatedAt,
	)
	return i, err
}
