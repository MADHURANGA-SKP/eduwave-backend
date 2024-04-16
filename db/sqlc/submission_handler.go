package db

import "context"

//CreateSubmissionParam contains the input parameters of the creations of the data
type CreateSubmissionParam struct {
	AssignmentID int64 `json:"assignment_id"`
	UserID       int64 `json:"user_id"`
}

//CreateAssignmentResponse contains the result of the creation the data
type CreateSubmissionResponse struct {
	Submission Submission `json:"submission"`
}

//CreateAssignment db handler for api call to Update assignment data in database
func (store *Store) CreateSubmission(ctx context.Context, arg CreateSubmissionParam) (CreateSubmissionResponse, error) {
	var result CreateSubmissionResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Submission, err = q.CreateSubmission(ctx, CreateSubmissionParams{
			AssignmentID: arg.AssignmentID,
			UserID: arg.UserID,
		})

		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}

//GetSubmissionParam contains the input parameters of getting the data
type GetsubmissionsByAssignmentParam struct {
	AssignmentID int64 `form:"assignment_id"`
}

type GetsubmissionsByAssignmentResponse struct {
	Submission Submission `json:"submission"`
}

//GetSubmission db handler for apu call to retrive submission data from the database
func (store *Store) GetsubmissionsByAssignment(ctx context.Context, arg GetsubmissionsByAssignmentParam) (GetsubmissionsByAssignmentResponse, error) {
	var result GetsubmissionsByAssignmentResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Submission, err = q.GetsubmissionsByAssignment(ctx, arg.AssignmentID)

		if err != nil {
			return err
		}

		return err
	})
	return result, err
}

//GetSubmissionParam contains the input parameters of getting the data
type GetsubmissionsByUserParam struct {
	AssignmentID int64 `form:"assignment_id"`
    UserID       int64 `form:"user_id"`
}

type GetsubmissionsByUserResponse struct {
	Submission Submission `json:"submission"`
}

//GetSubmission db handler for apu call to retrive submission data from the database
func (store *Store) GetsubmissionsByUser(ctx context.Context, arg GetsubmissionsByUserParam) (GetsubmissionsByUserResponse, error) {
	var result GetsubmissionsByUserResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Submission, err = q.GetsubmissionsByUser(ctx, arg.UserID)

		if err != nil {
			return err
		}

		return err
	})
	return result, err
}

//ListSubmissions db handler for api call to update a assignment data of the database
func (store *Store) Listsubmissions(ctx context.Context, params ListsubmissionsParams) ([]Submission, error) {
	return store.Queries.Listsubmissions(ctx, params)
}