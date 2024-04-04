package db

import "context"

//CreateCourseParam contain the input parameters of creating the course
type CreateCourseParam struct{
	Title       string `json:"title"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

//CreateCourseResponse contains the result of the create course data 
type CreateCourseResponse struct {
	Course Course `json:"course"`
}

//CreateCourse db handler for api call to create course in database
func(store *Store) CreateCourse(ctx context.Context, arg CreateCourseParam)(CreateCourseResponse, error){
	var result CreateCourseResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Course, err = q.CreateCourses(ctx, CreateCoursesParams{
			Title: arg.Title,
			Type: arg.Type,
			Description: arg.Description,
		})

		if err != nil {
			return err
		}
		return err
	})

	return result,err
}

//DeleteCourseParam contains the input parameters of the geting the data
type DeleteCourseParam struct {
	CourseID  int64         `json:"course_id"`
	TeacherID int64 `json:"teacher_id"`
}

//DeleteCourseResponse contains the result of the geting the data
type DeleteCourseResponse struct {
	Course Course `json:"course"`
}

//DeleteCourse db handler for api call to delete a course from the database
func (store *Store) DeleteCourse(ctx context.Context, arg DeleteCourseParam) error {
	return store.Queries.DeleteCourses(ctx, DeleteCoursesParams{
		CourseID: arg.CourseID,
		TeacherID: arg.TeacherID,
	})
}

//GetCourseParam contains the input parameters of the geting the data
type GetCourseParam struct {
	CourseID  int64         `json:"course_id"`
	TeacherID int64 `json:"teacher_id"`
}

//GetCourseResponse contains the result of the geting the data
type GetCourseResponse struct {
	Course Course `json:"course"`
}

//GetUser db handler for api call to retrive a admin data from the database
func (store *Store) GetCourse(ctx context.Context, arg GetCourseParam) (GetCourseResponse, error) {
	var result GetCourseResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Course, err = q.GetCourses(ctx, GetCoursesParams{
			CourseID: arg.CourseID,
			TeacherID: arg.TeacherID,
		})

		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}


//ListSubmissions db handler for api call to listcourse data of the database
func (store *Store) listCourses(ctx context.Context, params ListCoursesParams) ([]Course, error) {
	return store.Queries.ListCourses(ctx, params)
}

//UpdateCourseParam contains the input parameters of updating coruse data 
type UpdateCoursesParam struct {
	CourseID    int64  `json:"course_id"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

//UpdateCourseResponse Contains the result of the creating course data
type UpdateCoursesResponse struct {
	Course Course `json:"course"`
}

//UpdateCourse dn handler for api call to update course data in databse
func(store *Store) UpdateCourse(ctx context.Context, arg UpdateCoursesParam)(UpdateCoursesResponse, error){
	var result UpdateCoursesResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error 

		result.Course, err =q.UpdateCourses(ctx, UpdateCoursesParams{
			CourseID: arg.CourseID,
			Title: arg.Title,
			Type: arg.Type,
			Description: arg.Description,
		})

		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}