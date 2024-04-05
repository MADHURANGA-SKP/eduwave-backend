package db

import "context"

//CreateAdminParam contains the input parameters of the Createing the data
type CreateAdminParam struct {
	UserID         int64  `json:"user_id"`
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
			UserID: arg.UserID,
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
	AdminID int64 `uri:"admin_id,min=1"`
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