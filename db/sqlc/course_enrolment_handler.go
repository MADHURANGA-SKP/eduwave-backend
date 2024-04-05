package db

import "context"

//CreateCourseEnrolmentParam contains input paramters of create Progress
type CreateEnrolmentsParam struct {
	CourseID  int64 `json:"course_id"`
	RequestID int64 `json:"request_id"`
	UserID    int64 `json:"user_id"`
}

//CreateCoureseProgresresponse contains the result of the Progress data
type  CreateEnrolmentsResponse struct {
	CourseEnrolment CourseEnrolment `json:"course_enrolments"`
}

//CreateCourseEnrolments db handler for api call to create course progress data in database
func (store *Store) CreateCourseEnrolments(ctx context.Context, arg CreateEnrolmentsParam) (CreateEnrolmentsResponse, error) {
	var result CreateEnrolmentsResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.CourseEnrolment, err = q.CreateEnrolments(ctx, CreateEnrolmentsParams{
			CourseID: arg.CourseID,
			RequestID: arg.RequestID,
			UserID: arg.UserID,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}