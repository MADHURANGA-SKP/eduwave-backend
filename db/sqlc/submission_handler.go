package db

import "context"

//GetSubmissionParam contains the input parameters of getting the data
type GetSubmissionsParam struct {
	AssignmentID int64 `json:"assignment_id"`
    UserID       int64 `json:"user_id"`
}

type GetSubmissionResponse struct {
	Submission Submission `json:"submission"`
}

//GetSubmission db handler for apu call to retrive submission data from the database
func (store *Store) GetSubmission(ctx context.Context, arg GetSubmissionsParam) (GetSubmissionResponse, error) {
	var result GetSubmissionResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Submission, err = q.Getsubmissions(ctx, GetsubmissionsParams{
			AssignmentID: arg.AssignmentID,
			UserID: arg.UserID,
		})

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