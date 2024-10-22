// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: course_progress.sql

package db

import (
	"context"
)

const createCourseProgress = `-- name: CreateCourseProgress :one
INSERT INTO course_progress (
    enrolment_id,
    progress
) VALUES (
    $1, $2
) RETURNING courseprogress_id, enrolment_id, progress
`

type CreateCourseProgressParams struct {
	EnrolmentID int64  `json:"enrolment_id"`
	Progress    string `json:"progress"`
}

func (q *Queries) CreateCourseProgress(ctx context.Context, arg CreateCourseProgressParams) (CourseProgress, error) {
	row := q.queryRow(ctx, q.createCourseProgressStmt, createCourseProgress, arg.EnrolmentID, arg.Progress)
	var i CourseProgress
	err := row.Scan(&i.CourseprogressID, &i.EnrolmentID, &i.Progress)
	return i, err
}

const getCourseProgress = `-- name: GetCourseProgress :one
SELECT courseprogress_id, enrolment_id, progress FROM course_progress
WHERE courseprogress_id = $1 AND enrolment_id = $2 
LIMIT 1
`

type GetCourseProgressParams struct {
	CourseprogressID int64 `json:"courseprogress_id"`
	EnrolmentID      int64 `json:"enrolment_id"`
}

func (q *Queries) GetCourseProgress(ctx context.Context, arg GetCourseProgressParams) (CourseProgress, error) {
	row := q.queryRow(ctx, q.getCourseProgressStmt, getCourseProgress, arg.CourseprogressID, arg.EnrolmentID)
	var i CourseProgress
	err := row.Scan(&i.CourseprogressID, &i.EnrolmentID, &i.Progress)
	return i, err
}

const listCourseProgress = `-- name: ListCourseProgress :many
SELECT courseprogress_id, enrolment_id, progress FROM course_progress
WHERE enrolment_id = $1
ORDER BY courseprogress_id
LIMIT $2
OFFSET $3
`

type ListCourseProgressParams struct {
	EnrolmentID int64 `json:"enrolment_id"`
	Limit       int32 `json:"limit"`
	Offset      int32 `json:"offset"`
}

func (q *Queries) ListCourseProgress(ctx context.Context, arg ListCourseProgressParams) ([]CourseProgress, error) {
	rows, err := q.query(ctx, q.listCourseProgressStmt, listCourseProgress, arg.EnrolmentID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CourseProgress{}
	for rows.Next() {
		var i CourseProgress
		if err := rows.Scan(&i.CourseprogressID, &i.EnrolmentID, &i.Progress); err != nil {
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

const updateCourseProgress = `-- name: UpdateCourseProgress :one
UPDATE course_progress
SET progress = $2
WHERE enrolment_id = $1 
RETURNING courseprogress_id, enrolment_id, progress
`

type UpdateCourseProgressParams struct {
	EnrolmentID int64  `json:"enrolment_id"`
	Progress    string `json:"progress"`
}

func (q *Queries) UpdateCourseProgress(ctx context.Context, arg UpdateCourseProgressParams) (CourseProgress, error) {
	row := q.queryRow(ctx, q.updateCourseProgressStmt, updateCourseProgress, arg.EnrolmentID, arg.Progress)
	var i CourseProgress
	err := row.Scan(&i.CourseprogressID, &i.EnrolmentID, &i.Progress)
	return i, err
}
