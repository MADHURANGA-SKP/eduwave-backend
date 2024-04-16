package db

import (
	"context"
	"time"
)

// CreateAssignmentParam contains the input parameters of the creations of the data
type CreateAssignmentParam struct {
	ResourceID     int64     `json:"resource_id"`
	Type           string    `json:"type"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	SubmissionDate time.Time `json:"submission_date"`
}

// CreateAssignmentResponse contains the result of the creation the data
type CreateAssignmentResponse struct {
	Assignment Assignment `json:"assignment"`
}

// CreateAssignment db handler for api call to Update assignment data in database
func (store *Store) CreateAssignment(ctx context.Context, arg CreateAssignmentParam) (CreateAssignmentResponse, error) {
	var result CreateAssignmentResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Assignment, err = q.CreateAssignment(ctx, CreateAssignmentParams{
			ResourceID:     arg.ResourceID,
			Type:           arg.Type,
			Title:          arg.Title,
			Description:    arg.Description,
			SubmissionDate: arg.SubmissionDate,
		})

		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}

// DeleteAssignmentParam contains the input parameters of the geting  the data
type DeleteAssignmentParam struct {
	AssignmentID int64 `json:"assignment_id"`
	ResourceID   int64 `json:"resource_id"`
}

// DeleteAssignment db handler for api call to delete a admin from the database
func (store *Store) DeleteAssignment(ctx context.Context, arg DeleteAssignmentParam) error {
	return store.Queries.DeleteAssignment(ctx, DeleteAssignmentParams{
		AssignmentID: arg.AssignmentID,
		ResourceID:   arg.ResourceID,
	})
}

// GetAssignmentParam contains the input parameters of the geting  the data
type GetAssignmentParam struct {
	AssignmentID int64 `uri:"assignment_id"`
}

// GetAssignmentResponse contains the result of the geting the data
type GetAssignmentResponse struct {
	Assignment Assignment `json:"assignment"`
}

// GetAssignment db handler for api call to retrive a assignment data from the database
func (store *Store) GetAssignment(ctx context.Context, arg GetAssignmentParam) (GetAssignmentResponse, error) {
	var result GetAssignmentResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Assignment, err = q.GetAssignment(ctx, arg.AssignmentID)

		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}

// UpdateAssignmentParam contains the input parameters of the updating of the data
type UpdateAssignmentParam struct {
	AssignmentID   int64     `json:"assignment_id"`
	Type           string    `json:"type"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	SubmissionDate time.Time `json:"submission_date"`
}

// UpdateAssignmentResponse contains the result of the updating of the data
type UpdateAssignmentResponse struct {
	Assignment Assignment `json:"assignment"`
}

// UpdateAssignment db handler for api call to update a assignment data of the database
func (store *Store) UpdateAssignment(ctx context.Context, arg UpdateAssignmentParam) (UpdateAssignmentResponse, error) {
	var result UpdateAssignmentResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Assignment, err = q.UpdateAssignment(ctx, UpdateAssignmentParams{
			AssignmentID:   arg.AssignmentID,
			Type:           arg.Type,
			Title:          arg.Title,
			Description:    arg.Description,
			SubmissionDate: arg.SubmissionDate,
		})

		if err != nil {
			return err
		}

		return nil

	})
	return result, err
}
