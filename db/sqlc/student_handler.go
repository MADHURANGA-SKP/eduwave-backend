package db

import (
	"context"
	"database/sql"
)

// //CreateStudentParam contains the input parameters of the creation of the student
// type CreateStudentParam struct {
// 	UserName string `json:"user_name"`
// }

// //CreateStudentResponse contains the result of the Student Creation in databse
// type CreateStudentResponse struct {
// 	Student Student `json:"student"`
// }

// //CreateStudent db handler for api call to creation of the student in database
// func (store *Store) CreateStudent(ctx context.Context, arg CreateStudentParam) (CreateStudentResponse, error) {
// 	var result CreateStudentResponse

// 	err := store.execTx(ctx, func(q *Queries) error {
// 		var err error

// 		result.Student, err = q.CreateStudent(ctx, arg.UserName)

// 		if err != nil {
// 			return err
// 		}

// 		return nil
// 	})

// 	return result, err
// }

//DeleteStudent db handler for api call to celete ctudent from teh database
func (store *Store) DeleteStudent(ctx context.Context, studentID int64) error {
	return store.Queries.DeleteStudent(ctx, studentID)
}

//GetStudentParam contains the input paramters of the retrive data
type GetStudentParam struct {
	StudentID int64 `json:"student_id"`
}

//GetStudentResponse contains the results of the Reriving data
type GetStudentResponse struct {
	Student Student `json:"student"`
}

//GetStudentParams db handler for api call to Get ctudent details from the database
func (store *Store) GetStudent(ctx context.Context, arg GetStudentParam) (GetStudentResponse, error) {
	var result GetStudentResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Student, err = q.GetStudent(ctx, arg.StudentID)

		if err != nil {
			return err
		}

		return nil
	})
	return result, err

}

//ListStudent db handler for api call to List a ctudent data the database
func (store *Store) ListStudents(ctx context.Context, params ListStudentsParams) ([]Student, error) {
	return store.Queries.ListStudents(ctx, params)
}

//UpdateStudentParams contains the input paramters of the updating  data
type UpdateStudentParam struct {
	UserName sql.NullString `json:"user_name"`
}

//UpdateStudentResponse contains the result of the updating the data
type UpdateStudentResponse struct {
	Student Student `json:"student"`
}

//UpdateStudent db handler for api call to cpdate a student data in database
func (store *Store) UpdateStudent(ctx context.Context, arg UpdateStudentParams) (UpdateStudentResponse, error) {
	var result UpdateStudentResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		updateStudent, err := q.UpdateStudent(ctx, UpdateStudentParams{
			UserName: arg.UserName,
		})

		if err != nil {
			return err
		}

		result.Student = updateStudent
		return nil
	})
	return result, err
}

