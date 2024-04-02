package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

//store provide all funtions to execute db queries and data trival and transfers
type Store struct {
	*Queries
	db *sql.DB
}

//create NewStore
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

//execTX execute a funtion within a database action
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

//CreateAdminParam contains the input parameters of the Createing the data
type CreateAdminParam struct {
	FullName       string `json:"full_name"`
    UserName       string `json:"user_name"`
    Email          string `json:"email"`
    HashedPassword string `json:"hashed_password"`
}

//CreateAdminResponse contains the result of the Createing the data
type CreateAdminResponse struct {
	Admin Admin `json:"admin"`
}

//CreateAdmin db handler for api call to retrive a admin data from the database
func (store *Store) CreateAdmin(ctx context.Context, arg CreateAdminParam) (CreateAdminResponse, error) {
	var result CreateAdminResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Admin, err = q.CreateAdmin(ctx, CreateAdminParams{
			FullName: arg.FullName,
			UserName: arg.UserName,
			Email: arg.Email,
			HashedPassword: arg.HashedPassword,
		})

		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}

//DeleteAdmin db handler for api call to delete a admin from the database
func (store *Store) DeleteAdmin(ctx context.Context, adminID int64) error {
	return store.Queries.DeleteAdmin(ctx, adminID)
}

//GetAdminParam contains the input parameters of the geting the data
type GetAdminParam struct {
	AdminID int64 `json:"admin_id"`
}

//GetAdminResponse contains the result of the geting the data
type GetAdminResponse struct {
	Admin Admin `json:"admin"`
}

//GetAdmin db handler for api call to retrive a admin data from the database
func (store *Store) GetAdmin(ctx context.Context, arg GetAdminParam) (GetAdminResponse, error) {
	var result GetAdminResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Admin, err = q.GetAdmin(ctx, arg.AdminID)

		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}

//UpdateAdminParam contains the input parameters of the updating the data
type UpdateAdminParam struct {
	AdminID        int64  `json:"admin_id"`
    FullName       string `json:"full_name"`
    UserName       string `json:"user_name"`
    Email          string `json:"email"`
    HashedPassword string `json:"hashed_password"`
}

//UpdateAdminResponse contains the result of the updating the data
type UpdateAdminResponse struct {
	Admin Admin `json:"admin"`
}

//UpdateAdmin db handler for api call to Update a admin data in database
func (store *Store) UpdateAdmin(ctx context.Context, arg UpdateAdminParams) (UpdateAdminResponse, error) {
	var result UpdateAdminResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		updateAdmin, err := q.UpdateAdmin(ctx, UpdateAdminParams{
			AdminID: arg.AdminID,
			FullName: arg.FullName,
			UserName: arg.UserName,
			Email: arg.Email,
			HashedPassword: arg.HashedPassword,
		})

		if err != nil {
			return err
		}

		if updateAdmin.AdminID == 0 {
			return err
		}

		result.Admin = updateAdmin
		return nil
	})
	return result, err
}

//CreateAssignmentParam contains the input parameters of the creations of the data
type CreateAssignmentParam struct {
	Type           string    `json:"type"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	SubmissionDate time.Time `json:"submission_date"`
}

//CreateAssignmentResponse contains the result of the creation the data
type CreateAssignmentResponse struct {
	Assignment Assignment `json:"assignment"`
}

//CreateAssignment db handler for api call to Update assignment data in database
func (store *Store) CreateAssignment(ctx context.Context, arg CreateAssignmentParam) (CreateAssignmentResponse, error) {
	var result CreateAssignmentResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Assignment, err = q.CreateAssignment(ctx, CreateAssignmentParams{
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

//DeleteAssignmentParam contains the input parameters of the geting  the data
type DeleteAssignmentParam struct {
	AssignmentID int64         `json:"assignment_id"`
	ResourceID   int64 `json:"resource_id"`
}


//DeleteAssignment db handler for api call to delete a admin from the database
func (store *Store) DeleteAssignment(ctx context.Context, arg DeleteAssignmentParam) error {
	return store.Queries.DeleteAssignment(ctx, DeleteAssignmentParams{
		AssignmentID: arg.AssignmentID,
		ResourceID: arg.ResourceID,
	})
}

//GetAssignmentParam contains the input parameters of the geting  the data
type GetAssignmentParam struct {
	AssignmentID int64         `json:"assignment_id"`
	ResourceID   int64 `json:"resource_id"`
}

//GetAssignmentResponse contains the result of the geting the data
type GetAssignmentResponse struct {
	Assignment Assignment `json:"assignment"`
}

//GetAssignment db handler for api call to retrive a assignment data from the database
func (store *Store) GetAssignment(ctx context.Context, arg GetAssignmentParam)(GetAssignmentResponse, error){
	var result GetAssignmentResponse

	err := store.execTx(ctx, func (q *Queries) error {
		var err error

		result.Assignment, err = q.GetAssignment(ctx, GetAssignmentParams{
			AssignmentID: arg.AssignmentID,
			ResourceID: arg.ResourceID,
		})

		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}

//UpdateAssignmentParam contains the input parameters of the updating of the data
type UpdateAssignmentParam struct {
	ResourceID     int64 `json:"resource_id"`
	Type           string        `json:"type"`
	Title          string        `json:"title"`
	Description    string        `json:"description"`
	SubmissionDate time.Time     `json:"submission_date"`
}

//UpdateAssignmentResponse contains the result of the updating of the data
type UpdateAssignmentResponse struct {
	Assignment Assignment `json:"assignment"`
}

//UpdateAssignment db handler for api call to update a assignment data of the database
func (store *Store) UpdateAssignment(ctx context.Context, arg UpdateAssignmentParam)(UpdateAssignmentResponse, error){
	var result UpdateAssignmentResponse

	err := store.execTx(ctx, func (q *Queries) error {
		var err error

		updateAssignment, err := q.UpdateAssignment(ctx, UpdateAssignmentParams{
			ResourceID: arg.ResourceID,
			Type: arg.Type,
			Title: arg.Title,
			Description: arg.Description,
			SubmissionDate: arg.SubmissionDate,
		})

		if err != nil {
			return err
		}

		result.Assignment = updateAssignment
		return nil

	})
	return result, err
}

//ListEnrolments db handler for api call to list enrolment data of the database
func (store *Store) ListEnrolments(ctx context.Context, params ListEnrolmentsParams) ([]CourseEnrolment, error) {
	return store.Queries.ListEnrolments(ctx, params)
}

//CreateCourseProgresParam contains input paramters of create Progress
type CreateCourseProgressPram struct {
	Progress string `Json:"progress"`
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

		result.CourseProgress, err = q.CreateCourseProgress(ctx, arg.Progress)
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

//CreateRequestParam contains the input parameters of the creation of data
type CreateRequestParam struct {
	IsActive   sql.NullBool `json:"is_active"`
	IsPending  sql.NullBool `json:"is_pending"`
	IsAccepted sql.NullBool `json:"is_accepted"`
	IsDeclined sql.NullBool `json:"is_declined"`
}

//CreateRequestResponse contains the result of the creation of the data
type CreateRequestResponse struct {
	Request Request `json:"request"`
}

//CreateRequest db handler for api call to Create request in database
func (store *Store) CreateRequest(ctx context.Context, arg CreateRequestParam) (CreateRequestResponse, error) {
	var result CreateRequestResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Request, err = q.CreateRequest(ctx, CreateRequestParams{
			IsActive:   arg.IsActive,
			IsPending:  arg.IsPending,
			IsAccepted: arg.IsPending,
			IsDeclined: arg.IsDeclined,
		})

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

//GetRequestParam contains the input parameters of the retriving  data
type DeleteRequestParam struct {
	StudentID int64 `json:"student_id"`
	RequestID int64         `json:"request_id"`
}

//DeleteRequest db handler for api call to delete a request from the database
func (store *Store) DeleteRequest(ctx context.Context, arg DeleteRequestParam) error {
	return store.Queries.DeleteRequest(ctx, DeleteRequestParams{
		StudentID: arg.StudentID,
		RequestID: arg.RequestID,
	})
}

//GetRequestParam contains the input parameters of the retriving  data
type GetRequestParam struct {
	StudentID int64 `json:"student_id"`
	RequestID int64         `json:"request_id"`
}

//GetRequestResponse contains the result of the updating of the data
type GetRequestResponse struct {
	Request Request `json:"request"`
}

//GetRequest db handler for api call to retrive a progress data in the databse
func (store *Store) GetRequest(ctx context.Context, arg GetRequestParam) (GetRequestResponse, error) {
	var result GetRequestResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Request, err = q.GetRequest(ctx, GetRequestParams{
			StudentID: arg.StudentID,
			RequestID: arg.RequestID,
		})

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

//UpdateRequestsParam contains the input parameters og the updating of the data
type UpdateRequestsParam struct {
	IsActive   sql.NullBool `json:"is_active"`
	IsPending  sql.NullBool `json:"is_pending"`
	IsAccepted sql.NullBool `json:"is_accepted"`
	IsDeclined sql.NullBool `json:"is_declined"`
}

//UpdateRequestResponse contains the result of the updaing of the data
type UpdateRequestResponse struct {
	Request Request `json:"request"`
}

//UpdateRequest db handler for api call to update a request data of the database
func (store *Store) UpdateRequest(ctx context.Context, arg UpdateRequestsParam) (UpdateRequestResponse, error) {
	var result UpdateRequestResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		updateRequests, err := q.UpdateRequests(ctx, UpdateRequestsParams{
			IsActive:   arg.IsActive,
			IsPending:  arg.IsPending,
			IsAccepted: arg.IsAccepted,
			IsDeclined: arg.IsDeclined,
		})

		if err != nil {
			return err
		}

		result.Request = updateRequests
		return nil
	})

	return result, err
}

//ListRequest db handler for api call to list all request data of the database
func (store *Store) ListRequest(ctx context.Context, params ListRequestParams) ([]Request, error) {
	return store.Queries.ListRequest(ctx, params)
}

//CreateResourceParam contains the input parameters of data
type CreateResourceParam struct {
	Title      string       `json:"title"`
	Type       TypeResource `json:"type"`
	ContentUrl string       `json:"content_url"`
}

//CreateResourceResponse contains the result of the creation of data
type CreateResourceResponse struct {
	Resource Resource `json:"resource"`
}

//CreateResource db handler fro api call to update resource data in database
func (store *Store) CreateResource(ctx context.Context, arg CreateResourceParam) (CreateResourceResponse, error) {
	var result CreateResourceResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Resource, err = q.CreateResource(ctx, CreateResourceParams{
			Title:      arg.Title,
			Type:       arg.Type,
			ContentUrl: arg.ContentUrl,
		})

		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

//DeleteResource db handler for api call to delete exact data from the database
func (store *Store) DeleteResource(ctx context.Context, params DeleteResourceParams) error {
	return store.Queries.DeleteResource(ctx, params)
}

//GetResourceParam contains the input paramters of the retriving data
type GetResourceParam struct {
	MaterialID int64 `json:"Material_id"`
	ResourceID int64         `json:"resource_id"`
}

//GetResourceResponse contains the result of the retriving data
type GetResourceResponse struct {
	Resource Resource `json:"resource"`
}

//GetResource db handler for api call to retrive a resource data from teh databse
func (store *Store) GetResource(ctx context.Context, arg GetResourceParam) (GetResourceResponse, error) {
	var result GetResourceResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Resource, err = q.GetResource(ctx, GetResourceParams{
			MaterialID: arg.MaterialID,
			ResourceID: arg.ResourceID,
		})

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

//ListResource db handler for api call to retrive a resource data from teh databse
func (store *Store) ListResource(ctx context.Context, params ListResourceParams) ([]Resource, error) {
	return store.Queries.ListResource(ctx, params)
}

//UpdateResourceParam contains the input parameters of the updating data
type UpdateResourceParam struct {
	MaterialID int64        `json:"material_id"`
	Title        string        `json:"title"`
	Type         TypeResource  `json:"type"`
	ContentUrl   string        `json:"content_url"`
}

//UpdateResourceResponse contains the result of the updating data
type UpdateResourceResponse struct {
	Resource Resource `json:"resource"`
}

//UpdateResource db handler for api call to update resource data in the database
func(store *Store) UpdateResource(ctx context.Context, arg UpdateResourceParam)(UpdateResourceResponse, error){
	var result UpdateResourceResponse

	err := store.execTx( ctx, func(q *Queries) error {
		var err error

		result.Resource, err = q.UpdateResource(ctx, UpdateResourceParams{
			MaterialID: arg.MaterialID,
			Title: arg.Title,
			Type: arg.Type,
			ContentUrl:  arg.ContentUrl,
		})

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

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

//GetSubmissionParam contains the input parameters of getting the data
type GetSubmissionsParam struct {
	AssignmentID int64 `json:"assignment_id"`
    StudentID    int64 `json:"student_id"`
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
			StudentID:    arg.StudentID,
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

//CreateTeacherParam contains the input parameters of the creations of the data
type CreateTeacherParam struct {
	AdminID 	int64	`json:"admin_id"`
	FullName       string         `json:"full_name"`
	Email          string         `json:"email"`
	UserName       string `json:"user_name"`
	HashedPassword string         `json:"hashed_password"`
	IsActive       bool           `json:"is_active"`
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
			AdminID : 		arg.AdminID,
			FullName:       arg.FullName,
			Email:          arg.Email,
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
	TeacherID int64 `json:"teacher_id"`
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

//CreateUserParam contains the input parameters of data
type CreateUserParam struct {
	UserName       string   `json:"user_name"`
	Role           string `json:"role"`
	FullName       string   `json:"full_name"`
	HashedPassword string   `json:"hashed_password"`
	Email          string   `json:"email"`
}

//CreateUserResponse contains the result of the creation
type CreateUserResponse struct {
	User User `json:"user"`
	// Admin Admin `json:"admin"`
	// Student Student `json:"student"`
}

//CreateUser db handler fro api call to create a user in database
func (store *Store) CreateUser(ctx context.Context, arg CreateUserParam) (CreateUserResponse, error) {
	var result CreateUserResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.User, err = q.CreateUser(ctx, CreateUserParams{
			UserName:       arg.UserName,
			Role:           arg.Role,
			FullName:       arg.FullName,
			HashedPassword: arg.HashedPassword,
			Email:          arg.Email,
		})

		if err != nil {
			return err
		}
		// result.Student, err = q.CreateStudent(ctx, arg.)
		return nil
	})

	return result, err
}

//GetUserParam contains the input parameters of the geting the data
type GetUserParam struct {
	UserName string `json:"user_name"`
}

//GetUserResponse contains the result of the geting the data
type GetUserResponse struct {
	User User `json:"user"`
}

//GetUser db handler for api call to retrive a admin data from the database
func (store *Store) GetUser(ctx context.Context, arg GetUserParam) (GetUserResponse, error) {
	var result GetUserResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.User, err = q.GetUser(ctx, arg.UserName)

		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}

//UpdateUserParam contains the input parameters of the update the data
type UpdateUserParam struct {
	HashedPassword sql.NullString `json:"hashed_password"`
	FullName       sql.NullString `json:"full_name"`
	Email          sql.NullString `json:"email"`
	UserName       string         `json:"user_name"`
}

//UpdateUserResponse contains the result of the updating data
type UpdateUserResponse struct {
	User User `json:"user"`
}

//UpdateUser db handler for api call to update user data in database
func (store *Store) UpdateUser(ctx context.Context, arg UpdateUserParam) (UpdateUserResponse, error) {
	var result UpdateUserResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.User, err = q.UpdateUser(ctx, UpdateUserParams{
			HashedPassword: arg.HashedPassword,
			FullName:       arg.FullName,
			Email:          arg.Email,
			UserName:       arg.UserName,
		})

		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}

//CreateSessionParam contans input paramters of the creation of the session
type CreateSessionParam struct {
	SessionID    uuid.UUID `json:"session_id"`
	UserName     string    `json:"user_name"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
}

//CreateSessionResponse contains the resut of the creating the data
type CreateSessionResponse struct {
	Session Session `json:"session"`
}

//CreateSession db handler for api call to create session data in database
func (store *Store) CreateSession(ctx context.Context, arg CreateSessionParam) (CreateSessionResponse, error) {
	var result CreateSessionResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Session, err = q.CreateSession(ctx, CreateSessionParams{
			SessionID:    arg.SessionID,
			UserName:     arg.UserName,
			RefreshToken: arg.RefreshToken,
			UserAgent:    arg.UserAgent,
			ClientIp:     arg.ClientIp,
			IsBlocked:    arg.IsBlocked,
			ExpiresAt:    arg.ExpiresAt,
		})

		if err != nil {
			return err
		}
		return nil
	})

	return result, err
}

// //GetSessionParam contains the input parameters of getting session data
// type GetSessionparam struct {
// 	SessionID    uuid.UUID `json:"session_id"`
// }

// //GetSessionResponse contain the result of the getting session data
// type GetSessionResponse struct {
// 	Session Session `json:"session"`
// }

// //GetSession db handler for api call to get session data in database
// func(store *Store) GetSession(ctx context.Context, arg GetSessionparam)(GetSessionResponse, error){
// 	var result GetSessionResponse

// 	err := store.execTx(ctx, func(q *Queries) error {
// 		var err error

// 		result.Session, err = q.GetSession(ctx, arg.SessionID)

// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// 	return result, err
// }

//CreateVerifyEmailParam contaisn the input parameters of verify email data
type CreateVerifyEmailParam struct {
	UserName   string `json:"user_name"`
	Email      string `json:"email"`
	SecretCode string `json:"secret_code"`
}

//CreateVersifyEmailResponse contains the result of Creating session data
type CreateVerifyEmailResponse struct {
	VerifyEmail VerifyEmail `json:"verify_email"`
}

//CreateSession db handler for api call to create email verification data in database
func (store *Store) CreateVerifyEmail(ctx context.Context, arg CreateVerifyEmailParam) (CreateVerifyEmailResponse, error) {
	var result CreateVerifyEmailResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.VerifyEmail, err = q.CreateVerifyEmail(ctx, CreateVerifyEmailParams{
			UserName:   arg.UserName,
			Email:      arg.Email,
			SecretCode: arg.SecretCode,
		})

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

//UpdateVersifyEmailParam contains input parameters of update Verify email data
type UpdateVerifyEmailParam struct {
	EmailID    int64  `json:"email_id"`
	SecretCode string `json:"secret_code"`
}

//UpdateVerifyEmailResponse contains the result of updating verify email data
type UpdateVerifyEmailResponse struct {
	VerifyEmail VerifyEmail `json:"verify_email"`
}

//UpdateVeridyEmail db handler for api call to update verify emaildata in database
func (store *Store) UpdateVerifyEmail(ctx context.Context, arg UpdateVerifyEmailParam) (UpdateVerifyEmailResponse, error) {
	var result UpdateVerifyEmailResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.VerifyEmail, err = q.UpdateVerifyEmail(ctx, UpdateVerifyEmailParams{
			EmailID:    arg.EmailID,
			SecretCode: arg.SecretCode,
		})

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

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

//CreateMaterialParam contains the input parameters of  creating Material data
type CreateMaterialParam struct{
	Title       string `json:"title"`
	Description string `json:"description"`
}

//CreateMaterialResponse contains the result of the creating Material data 
type CreateMaterialReponse struct {
	Material Material `json:"Material"`
}

//CreateMaterial db handler for api call to create Material data in databse
func(store *Store) CreateMaterial(ctx context.Context, arg CreateMaterialParam)(CreateMaterialReponse, error){
	var result CreateMaterialReponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Material, err = q.CreateMaterial(ctx, CreateMaterialParams{
			Title: arg.Title,
			Description: arg.Description,
		})
		if err != nil {
			return err
		}
		return nil
	})
	return result, err
}

//DeleteMaterialParam contaisn the input parameters of delete Material data
type DeleteMaterialParam struct {
	MaterialID int64         `json:"Material_id"`
	CourseID   int64`json:"course_id"`
}

//DeleteMatirila db handler for api call to delete Material data in database
func(store *Store) DeleteMaterial(ctx context.Context, arg DeleteMaterialParam)error {
	return store.Queries.DeleteMaterial(ctx, DeleteMaterialParams{
		MaterialID: arg.MaterialID,
		CourseID: arg.CourseID,
	})
}

//GetMaterialparam contains the input parameters of the get Material data
type GetMaterialParam struct {
	CourseID int64 `json:"course_id"`
}

//GetMaterialResponse contains the result of the get matrial data
type GetMaterialResponse struct {
	Material Material `json:"Material"`
}

//GetMaterial db handler for api call to get Material data in database
func(store *Store) GetMaterial(ctx context.Context, arg GetMaterialParam)(GetMaterialResponse, error){
	var result GetMaterialResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Material, err = q.GetMaterial(ctx, arg.CourseID)
		if err != nil {
			return err
		}
		return nil
	})
	return result, err
}

//ListMaterial db handler fro api call to list Material data in database
func(store *Store) ListMaterial(ctx context.Context, params ListMaterialParams)([]Material, error){
	return store.Queries.ListMaterial(ctx, params)
}

//UpdateMaterialParam contains the input parameters of the Update Material data
type UpdateMaterialParam struct {
	MaterialID  int64         `json:"Material_id"`
	CourseID    int64 `json:"course_id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
}

//UpdateMaterialResponse contains the result of the Updated Material data
type UpdateMaterialResponse struct {
	Material Material `json:"Material"`
}

//UpdateMatririal db handler for api call to the update Material data in database
func(store *Store) UpdateMaterials(ctx context.Context, arg UpdateMaterialParam)(UpdateMaterialResponse, error){
	var result UpdateMaterialResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Material, err = q.UpdateMaterial(ctx, UpdateMaterialParams{
			MaterialID: arg.MaterialID,
			CourseID: arg.CourseID,
			Title: arg.Title,
			Description: arg.Description,
		})
		if err != nil {
			return err
		}
		return err
	})
	return result, err
}
