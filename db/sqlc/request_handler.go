package db

import (
	"context"
	"database/sql"
)

//CreateRequestParam contains the input parameters of the creation of data
type CreateRequestParam struct {
	UserID     int64        `json:"user_id"`
    CourseID   int64        `json:"course_id"`
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
			UserID: arg.UserID,
			CourseID: arg.CourseID,
			IsActive:   arg.IsAccepted,
			IsPending:  arg.IsPending,
			IsAccepted: arg.IsAccepted,
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
	UserID    int64 `json:"user_id"`
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

		result.Request, err = q.GetRequest(ctx, arg.UserID)

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

//UpdateRequestsParam contains the input parameters og the updating of the data
type UpdateRequestsParam struct {
	UserID     int64        `json:"user_id"`
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
func (store *Store) UpdateRequests(ctx context.Context, arg UpdateRequestsParam) (UpdateRequestResponse, error) {
	var result UpdateRequestResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		updateRequests, err := q.UpdateRequests(ctx, UpdateRequestsParams{
			UserID: arg.UserID,
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

//ListRequestByUser db handler for api call to list all request data of the database by user
func (store *Store) ListRequestByUser(ctx context.Context, params ListRequestByUserParams) ([]Request, error) {
	return store.Queries.ListRequestByUser(ctx, params)
}