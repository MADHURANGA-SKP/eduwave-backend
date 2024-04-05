package db

import "context"

//CreateTeacherParam contains the input parameters of the creations of the data
type CreateTeacherParam struct {
	UserID         int64  `json:"user_id"`
    FullName       string `json:"full_name"`
    Email          string `json:"email"`
    Qualification  string `json:"qualification"`
    UserName       string `json:"user_name"`
    HashedPassword string `json:"hashed_password"`
    IsActive       bool   `json:"is_active"`
}

//CreateTeacherResponse contains the result of the creation the data
type CreateTeacherResponse struct {
	Teacher Teacher `json:"teacher"`
}

//CreateTeacher db handler for api call to create ceacher data in database
func (store *Store) CreateTeacher(ctx context.Context, arg CreateTeacherParam) (CreateTeacherResponse, error) {
	var result CreateTeacherResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Teacher, err = q.CreateTeacher(ctx, CreateTeacherParams{
			UserID: arg.UserID,
			FullName:       arg.FullName,
			Email:          arg.Email,
			Qualification: arg.Qualification,
			UserName:       arg.UserName,
			HashedPassword: arg.HashedPassword,
			IsActive:       arg.IsActive,
		})

		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}

//DeleteAssignment db handler for api call to delete a admin from the database
func (store *Store) DeleteTeacher(ctx context.Context, teacherID int64) error {
	return store.Queries.DeleteTeacher(ctx, teacherID)
}

//GetTeacherParam contains the input paramters of the retrive data
type GetTeacherParam struct {
	TeacherID int64 `uri:"id,min=1"`
}

//GetTeacherResponse contains the results of the Reriving data
type GetTeacherResponse struct {
	Teacher Teacher `json:"teacher"`
}

//GetTeacherParams db handler for api call to Get teacher details from the database
func (store *Store) GetTeacher(ctx context.Context, arg GetTeacherParam) (GetTeacherResponse, error) {
	var result GetTeacherResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Teacher, err = q.GetTeacher(ctx, arg.TeacherID)

		if err != nil {
			return err
		}

		return nil
	})
	return result, err

}

//ListTeacher db handler for api call to List a assignment data of the database
func (store *Store) ListTeachers(ctx context.Context, params ListTeacherParams) ([]Teacher, error) {
	return store.Queries.ListTeacher(ctx, params)
}

//UpdateTeacherParams contains the input parameters of updating data
type UpdateTeacherParam struct {
	TeacherID      int64  `json:"teacher_id"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	UserName       string `json:"user_name"`
	HashedPassword string `json:"hashed_password"`
	IsActive       bool   `json:"is_active"`
}

//UpdateTeachersResponse contains the result of the upating data
type UpdateTeacherResponse struct {
	Teacher Teacher `json:"teacher"`
}

//UpdateTeacher db handler for api call t o update teacher data in database
func(store *Store) UpdateTeacher(ctx context.Context, arg UpdateTeacherParam)(UpdateTeacherResponse, error){
	var result UpdateTeacherResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Teacher, err = q.UpdateTeacher(ctx, UpdateTeacherParams{
			TeacherID: arg.TeacherID,
			FullName: arg.FullName,
			Email: arg.Email,
			UserName: arg.UserName,
			HashedPassword: arg.HashedPassword,
			IsActive: arg.IsActive,
		})

		if err != nil {
			return err
		}

		return err
	})

	return result, err
}