package db

import (
	"context"
	"database/sql"
)

// CreateUserParam contains the input parameters of data
type CreateUserParam struct {
	UserName       string `json:"user_name"`
	FullName       string `json:"full_name"`
	HashedPassword string `json:"hashed_password"`
	Email          string `json:"email"`
	Role           string `json:"role"`
	Qualification  string `json:"qualification"`
}

// CreateUserResponse contains the result of the creation
type CreateUserResponse struct {
	User User `json:"user"`
}

// CreateUser db handler fro api call to create a user in database
func (store *Store) CreateUser(ctx context.Context, arg CreateUserParam) (CreateUserResponse, error) {
	var result CreateUserResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.User, err = q.CreateUser(ctx, CreateUserParams{
			UserName:       arg.UserName,
			FullName:       arg.FullName,
			HashedPassword: arg.HashedPassword,
			Email:          arg.Email,
			Role:           arg.Role,
			Qualification:  arg.Qualification,
		})

		if err != nil {
			return err
		}
		// result.Student, err = q.CreateStudent(ctx, arg.)
		return nil
	})

	return result, err
}

// GetUserParam contains the input parameters of the geting the data
type GetUserParam struct {
	UserName string `json:"user_name"`
}

// GetUserResponse contains the result of the geting the data
type GetUserResponse struct {
	User User `json:"user"`
}

// GetUser db handler for api call to retrive a admin data from the database
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

// UpdateUserParam contains the input parameters of the update the data
type UpdateUserParam struct {
	HashedPassword  sql.NullString `json:"hashed_password"`
	FullName        sql.NullString `json:"full_name"`
	Email           sql.NullString `json:"email"`
	UserName        string         `json:"user_name"`
	IsEmailVerified sql.NullBool   `json:"is_email_verified"`
}

// UpdateUserResponse contains the result of the updating data
type UpdateUserResponse struct {
	User User `json:"user"`
}

// UpdateUser db handler for api call to update user data in database
func (store *Store) UpdateUser(ctx context.Context, arg UpdateUserParam) (UpdateUserResponse, error) {
	var result UpdateUserResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.User, err = q.UpdateUser(ctx, UpdateUserParams{
			HashedPassword:  sql.NullString{String: arg.HashedPassword.String, Valid: true},
			FullName:        sql.NullString{String: arg.FullName.String, Valid: true},
			Email:           sql.NullString{String: arg.Email.String, Valid: true},
			IsEmailVerified: sql.NullBool{Bool: true, Valid: true},
			UserName:        arg.UserName,
		})

		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}

// ListUserParams contains the input parameters for list users
func (store *Store) ListUsers(ctx context.Context, params ListUserParams) ([]User, error) {
	return store.Queries.ListUser(ctx, params)
}
