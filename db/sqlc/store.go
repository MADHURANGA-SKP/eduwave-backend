package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

//store provide all funtions to execute db queries and data trival and transfers
type Store struct {
	*Queries
	db *sql.DB
}

//create NewStore
func NewStore(db *sql.DB) *Store{
	return &Store{
		db: db,
		Queries: New(db),
	}
}

//execTX execute a funtion within a database action
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error{
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q :=New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

//DeleteAdmin db handler for api call to delete a admin from the database
func (store *Store) DeleteAdmin(ctx context.Context, adminID int64) error {
    return store.Queries.DeleteAdmin(ctx, adminID)
}

//GetAdminParam contains the input parameters of the geting the data 
type GetAdminParam struct{
	AdminID  int64 `json:"admin_id"`
}

//GetAdminResponse contains the result of the geting the data
type GetAdminResponse struct{
	Admin Admin `json:"admin"`
}

//GetAdmin db handler for api call to retrive a admin data from the database
func (store *Store) GetAdmin(ctx context.Context, arg GetAdminParam)(GetAdminResponse, error){
	var result GetAdminResponse 

	err := store.execTx(ctx, func(q *Queries) error{
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
type UpdateAdminParam struct{
	UserName sql.NullString `json:"user_name"`
}

//UpdateAdminResponse contains the result of the updating the data
type UpdateAdminResponse struct{
	Admin Admin `json:"admin"`
}
 
//UpdateAdmin db handler for api call to Update a admin data in database
func (store *Store) UpdateAdmin(ctx context.Context, arg UpdateAdminParams)(UpdateAdminResponse, error){
	var result UpdateAdminResponse

	err := store.execTx(ctx, func(q *Queries) error{
		var err error
		updateAdmin, err := q.UpdateAdmin(ctx, UpdateAdminParams{
			UserName: arg.UserName,
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
type CreateAssignmentParam struct{
	Type           string        `json:"type"`
	Title          string        `json:"title"`
	Description    string        `json:"description"`
	SubmissionDate time.Time     `json:"submission_date"`
}

//CreateAssignmentResponse contains the result of the creation the data
type CreateAssignmentResponse struct{
	Assignment Assignment `json:"assignment"`
}

//CreateAssignment db handler for api call to Update assignment data in database
func (store *Store) CreateAssignment(ctx context.Context, arg CreateAssignmentParam)(CreateAssignmentResponse, error){
		var result CreateAssignmentResponse

		err := store.execTx(ctx, func(q *Queries) error{
			var err error
			result.Assignment, err = q.CreateAssignment(ctx, CreateAssignmentParams{
				Type: arg.Type,
				Title: arg.Title,
				Description: arg.Description,
				SubmissionDate: arg.SubmissionDate,
			});
			
			if err != nil {
				return err
			}
	
		
			return nil
		})
		return result, err
}

//DeleteAssignment db handler for api call to delete a admin from the database
func (store *Store) DeleteAssignment(ctx context.Context, assignmentID int64) error {
	return store.Queries.DeleteAssignment(ctx, assignmentID)
}

//GetAssignmentParam contains the input parameters of the geting  the data
type GetAssignmentParam struct {
	CourseID sql.NullInt64 `json:"course_id"`
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
		
		result.Assignment, err = q.GetAssignment(ctx, arg.CourseID)

		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}

//UpdateAssignmentParam contains the input parameters of the updating of the data 
type UpdateAssignmentParam struct {
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

//ListEnrolments db handler for api call to update a assignment data of the database
func (store *Store) ListEnrolments(ctx context.Context, params ListEnrolmentsParams) ([]CourseEnrolment, error) {
    return store.Queries.ListEnrolments(ctx, params)
}

//GetCourseProgressParam db handler for api call to Get course progress data from the database
type GetCourseProgressParam struct {
	CourseprogressID int64         `json:"courseprogress_id"`
	EnrolmentID      sql.NullInt64 `json:"enrolment_id"`
}

//GetCourseProgressResponse contains the result of the updating of the data
type GetCourseProgressResponse struct {
	CourseProgress CourseProgress `json:"course_progress"`
}

//GetCourseProgress db handler for api call to retrive a progress data from the databse
func(store *Store) GetCourseProgress(ctx context.Context, arg GetCourseProgressParam)(GetCourseProgressResponse, error){
	var result GetCourseProgressResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.CourseProgress, err = q.GetCourseProgress(ctx, GetCourseProgressParams{
			CourseprogressID : arg.CourseprogressID,
			EnrolmentID: arg.EnrolmentID,
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
func(store *Store) CreateRequest(ctx context.Context, arg CreateRequestParam)(CreateRequestResponse, error){
	var result CreateRequestResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error 

		result.Request, err = q.CreateRequest(ctx, CreateRequestParams{
			IsActive: arg.IsActive,
			IsPending: arg.IsPending,
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

//DeleteRequest db handler for api call to delete a request from the database
func (store *Store) DeleteRequest(ctx context.Context, requestID int64) error {
	return store.Queries.DeleteRequest(ctx, requestID)
}

//GetRequestParam contains the input parameters of the retriving  data
type GetRequestParam struct {
	RequestID int64 `json:"Request_id"`
}	

//GetRequestResponse contains the result of the updating of the data
type GetRequestResponse struct {
	Request Request `json:"request"`
}	

//GetRequest db handler for api call to retrive a progress data in the databse
func(store *Store) GetRequest(ctx context.Context, arg GetRequestParam)(GetRequestResponse, error){
	var result GetRequestResponse
	
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		
		result.Request, err = q.GetRequest(ctx, arg.RequestID)
		
		if err != nil {
			return err
		}	
		
		return nil
	})	
	
	return result, err
}	

//UpdateRequestsParam contains the input parameters og the updating of the data
type UpdateRequestsParam struct {
	IsActive   sql.NullBool  `json:"is_active"`
	IsPending  sql.NullBool  `json:"is_pending"`
	IsAccepted sql.NullBool  `json:"is_accepted"`
	IsDeclined sql.NullBool  `json:"is_declined"`
}

//UpdateRequestResponse contains the result of the updaing of the data
type UpdateRequestResponse struct {
	Request Request `json:"request"`
}

//UpdateRequest db handler for api call to update a request data of the database
func(store *Store) UpdateRequest(ctx context.Context, arg UpdateRequestsParam)(UpdateRequestResponse, error){
	var result UpdateRequestResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error 

		updateRequests, err := q.UpdateRequests(ctx, UpdateRequestsParams{
			IsActive: arg.IsActive,
			IsPending: arg.IsPending,
			IsAccepted: arg.IsAccepted,
			IsDeclined: arg.IsDeclined,
		})

		if err != nil {
			return err
		}
		
		result.Request = updateRequests
		return nil
	})

	return result,err
}


//ListRequest db handler for api call to List all Request data of the database
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
func(store *Store) CreateResource(ctx context.Context, arg CreateResourceParam)(CreateResourceResponse, error){
	var result CreateResourceResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error 
		result.Resource, err = q.CreateResource(ctx, CreateResourceParams{
			Title: arg.Title,
			Type: arg.Type,
			ContentUrl: arg.ContentUrl,
		})

		if err != nil {
			return err
		}

		return err
	})

	return result, err
}


// //DeleteResource db handler fro api call to delete exact data from the database
func (store *Store) DeleteResource(ctx context.Context, params DeleteResourceParams) error{
	return store.Queries.DeleteResource(ctx, params)
}

//GetResourceParam contains the input paramters of the retriving data
type GetResourceParam struct {
	AssignmentID sql.NullInt64 `json:"assignment_id"`
	CourseID     sql.NullInt64 `json:"course_id"`
}

//GetResourceResponse contains the result of the retriving data
type GetResourceResponse struct {
	Resource Resource `json:"resource"`
}

//GetResource db handler for api call to retrive a resource data from teh databse
func(store *Store) GetResource(ctx context.Context, arg GetResourceParam)(GetResourceResponse, error){
	var result GetResourceResponse

	err := store.execTx(ctx, func (q *Queries) error {
		var err error

		result.Resource, err = q.GetResource(ctx, GetResourceParams{
			AssignmentID: arg.AssignmentID,
			CourseID: arg.CourseID,
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
    return store.Queries.ListResource(ctx, params )
}


//UpdateResourceParam contains the input parameters of the updating data
type UpdateResourceParam struct {
	Title        string        `json:"title"`
	Type         TypeResource  `json:"type"`
	ContentUrl   string        `json:"content_url"`
}

//UpdateResourceResponse contains the result of the updating data
type UpdateResourceResponse struct {
	Resource Resource `json:"resource"`
}

//UpdateResource db handler for api call to Update resource data in the database
func(store *Store) UpdateResource(ctx context.Context, arg UpdateResourceParam)(UpdateResourceResponse, error){
	var result UpdateResourceResponse

	err := store.execTx( ctx, func(q *Queries) error {
		var err error 

		result.Resource, err = q.UpdateResource(ctx, UpdateResourceParams{
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

//CreateStudentParam contains the input parameters of the creation of the student
type CreateStudentParam struct {
	UserName sql.NullString `json:"user_name"`
}

//CreateStudentResponse contains the result of the Student Creation in databse
type CreateStudentResponse struct {
	Student Student `json:"student"`
}

//CreateStudent db handler for api call to Creation of the student in database
func(store *Store) CreateStudent(ctx context.Context, arg CreateStudentParam)(CreateStudentResponse, error){
	var result CreateStudentResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Student, err = q.CreateStudent(ctx, arg.UserName)

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
} 

//DeleteStudent db handler for api call to Delete Student from teh database
func(store *Store) DeleteStudent(ctx context.Context, studentID int64) error {
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

//GetStudentParams db handler for api call to Get Student details from the database
func(store *Store) GetStudent(ctx context.Context, arg GetStudentParam)(GetStudentResponse, error){
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

//ListStudent db handler for api call to List a Student data the database
func (store *Store) ListStudents(ctx context.Context, params ListStudentsParams) ([]Student, error) {
    return store.Queries.ListStudents(ctx, params )
}

//UpdateStudentParams contains the input paramters of the updating  data
type UpdateStudentParam struct {
	UserName  sql.NullString `json:"user_name"`
}

//UpdateStudentResponse contains the result of the updating the data
type UpdateStudentResponse struct{
	Student Student `json:"student"`
}
 
//UpdateStudent db handler for api call to Update a student data in database
func (store *Store) UpdateStudent(ctx context.Context, arg UpdateStudentParams)(UpdateStudentResponse, error){
	var result UpdateStudentResponse

	err := store.execTx(ctx, func(q *Queries) error{
		var err error
		updateStudent, err := q.UpdateStudent(ctx, UpdateStudentParams{
			UserName: arg.UserName,
		})
 
		if err != nil {
			return err
		}

		if updateStudent.StudentID == 0 {
			return err
		}

		result.Student = updateStudent
		return nil
	})
	return result, err
}




