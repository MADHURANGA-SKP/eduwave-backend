package db

import (
	"context"
	"database/sql"
	"fmt"
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


//GetAccountTxParams contains the input parameters of the Geting of the data 
type GetAdminParam struct{
	AdminID  int64 `json:"admin_id" binding:"required"`
	UserName  string `json:"user_name" binding:"required"`
}

//GetAccountResult contains the result of the Geting of the data
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


//UpdateteTodoTxParams contains the input parameters of the Updating of the data 
type UpdateAdminParam struct{
	AdminID  int64          `json:"admin_id"`
	UserName sql.NullString `json:"user_name"`
}


//UpdateTodoResult contains the result of the Updating of the data
type UpdateAdminResponse struct{
	Admin Admin `json:"admin"`
}

 
//UpdateAdmin db handler for api call to Update a admin data in database
func (store *Store) UpdateAdmin(ctx context.Context, arg UpdateAdminParams)(UpdateAdminResponse, error){
	var result UpdateAdminResponse

	err := store.execTx(ctx, func(q *Queries) error{
		var err error
		updateAdmin, err := q.UpdateAdmin(ctx, UpdateAdminParams{
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