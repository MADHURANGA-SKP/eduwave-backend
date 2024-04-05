package db

import "context"

//GetRoleParam contains the input parameters of getting the data
type GetRoleParam struct {
	RoleID   int64  `json:"role_id"`
}

type GetRoleResponse struct {
	Role Role `json:"Role"`
}

//GetRole db handler for apu call to retrive Role data from the database
func (store *Store) GetRole(ctx context.Context, arg GetRoleParam) (GetRoleResponse, error) {
	var result GetRoleResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Role, err = q.GetRole(ctx, arg.RoleID)

		if err != nil {
			return err
		}

		return err
	})
	return result, err
}