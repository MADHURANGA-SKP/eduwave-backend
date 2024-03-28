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

//GetAdminParam contains the input parameters of the Geting of the data 
type GetAdminParam struct{
	AdminID  int64 `json:"admin_id"`
}

//GetAdminResponse contains the result of the Geting of the data
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

//UpdateAdminParam contains the input parameters of the Updating of the data 
type UpdateAdminParam struct{
	UserName sql.NullString `json:"user_name"`
}

//UpdateAdminResponse contains the result of the Updating of the data
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

//CreateAssignmentParam contains the input parameters of the Updating of the data
type CreateAssignmentParam struct{
	Type           string        `json:"type"`
	Title          string        `json:"title"`
	Description    string        `json:"description"`
	SubmissionDate time.Time     `json:"submission_date"`
}

//CreateAssignmentResponse contains the result of the Geting of the data
type CreateAssignmentResponse struct{
	Assignment Assignment `json:"assignment"`
}

//CreateAssignment db handler for api call to Update a admin data in database
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

//GetAssignmentParam contains the input parameters of the Geting of the data
type GetAssignmentParam struct {
	CourseID sql.NullInt64 `json:"course_id"`
}

//GetAssignmentResponse contains the result of the Geting of the data
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

//UpdateAssignmentParam contains the input parameters of the Updating of the data 
type UpdateAssignmentParam struct {
	Type           string        `json:"type"`
	Title          string        `json:"title"`
	Description    string        `json:"description"`
	SubmissionDate time.Time     `json:"submission_date"`
}

//UpdateAssignmentResponse contains the result of the Updating of the data
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

//GetCourseProgressResponse contains the result of the Updating of the data
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

//CreateRequestParam contains the unput parameters of the Updating of data 
type CreateRequestParam struct {
	IsActive   sql.NullBool `json:"is_active"`
	IsPending  sql.NullBool `json:"is_pending"`
	IsAccepted sql.NullBool `json:"is_accepted"`
	IsDeclined sql.NullBool `json:"is_declined"`
}

//CreateRequestResponse contains the result of the Creation of the data 
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

//GetRequestParam db handler for api call to Get Request data from the database
type GetRequestParam struct {
	RequestID int64 `json:"Request_id"`
}	

//GetRequestResponse contains the result of the Updating of the data
type GetRequestResponse struct {
	Request Request `json:"request"`
}	

//GetRequest db handler for api call to retrive a progress data from the databse
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

