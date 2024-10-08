package db

import "context"

//CreateCourseProgresParam contains input paramters of create Progress
type CreateCourseProgressPram struct {
	EnrolmentID int64  `json:"enrolment_id"`
    Progress    string `json:"progress"`
}

//CreateCoureseProgresresponse contains the result of the Progress data
type CreateCoureseProgresResponse struct {
	CourseProgress CourseProgress `json:"course_progress"`
}

//CreateCourseProgress db handler for api call to create course progress data in database
func (store *Store) CreateCourseProgress(ctx context.Context, arg CreateCourseProgressPram) (CreateCoureseProgresResponse, error) {
	var result CreateCoureseProgresResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.CourseProgress, err = q.CreateCourseProgress(ctx, CreateCourseProgressParams{
			EnrolmentID: arg.EnrolmentID,
			Progress: arg.Progress,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

//GetCourseProgressParam contains input parameters to get courseprogress data
type GetCourseProgressParam struct {
	CourseprogressID int64         `json:"courseprogress_id"`
	EnrolmentID      int64 `json:"enrolment_id"`
}

//GetCourseProgressResponse contains the result of the updating of the data
type GetCourseProgressResponse struct {
	CourseProgress CourseProgress `json:"course_progress"`
}

//GetCourseProgress db handler for api call to retrive a progress data from the databse
func (store *Store) GetCourseProgress(ctx context.Context, arg GetCourseProgressParam) (GetCourseProgressResponse, error) {
	var result GetCourseProgressResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.CourseProgress, err = q.GetCourseProgress(ctx, GetCourseProgressParams{
			CourseprogressID: arg.CourseprogressID,
			EnrolmentID:      arg.EnrolmentID,
		})

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}


//ListCoursebyuser db handler for api call to listcourse by created user data of the database
func (store *Store) ListCourseProgress(ctx context.Context, params ListCourseProgressParams) ([]CourseProgress, error) {
	return store.Queries.ListCourseProgress(ctx, params)
}


//UpdateCourseProgressParam contains the input parameters of updating coruse progress data 
type UpdateCourseProgressParam struct {
	EnrolmentID int64  `json:"enrolment_id"`
	Progress    string `json:"progress"`
}

//UpdateCourseResponse Contains the result of the creating course data
type UpdateCourseProgressResponse struct {
	CourseProgress CourseProgress `json:"course_progress"`
}

//UpdateCourse dn handler for api call to update course data in databse
func(store *Store) UpdateCourseProgress(ctx context.Context, arg UpdateCourseProgressParam)(UpdateCourseProgressResponse, error){
	var result UpdateCourseProgressResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error 

		result.CourseProgress, err =q.UpdateCourseProgress(ctx, UpdateCourseProgressParams{
			EnrolmentID: arg.EnrolmentID,
			Progress: arg.Progress,
		})

		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}
