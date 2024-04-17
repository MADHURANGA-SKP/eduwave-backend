package db

import "context"

// CreateVerifyEmailParam contaisn the input parameters of verify email data
type CreateVerifyEmailParam struct {
	UserName   string `json:"user_name"`
	Email      string `json:"email"`
	SecretCode string `json:"secret_code"`
}

// CreateVersifyEmailResponse contains the result of Creating session data
type CreateVerifyEmailResponse struct {
	VerifyEmail VerifyEmail `json:"verify_email"`
}

// CreateSession db handler for api call to create email verification data in database
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

// UpdateVersifyEmailParam contains input parameters of update Verify email data
type UpdateVerifyEmailParam struct {
	IsUsed     bool   `json:"is_used"`
	EmailID    int64  `json:"email_id"`
	SecretCode string `json:"secret_code"`
}

// UpdateVerifyEmailResponse contains the result of updating verify email data
type UpdateVerifyEmailResponse struct {
	VerifyEmail VerifyEmail `json:"verify_email"`
}

// UpdateVeridyEmail db handler for api call to update verify emaildata in database
func (store *Store) UpdateVerifyEmail(ctx context.Context, arg UpdateVerifyEmailParam) (UpdateVerifyEmailResponse, error) {
	var result UpdateVerifyEmailResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.VerifyEmail, err = q.UpdateVerifyEmail(ctx, UpdateVerifyEmailParams{
			IsUsed:     arg.IsUsed,
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

// CreateVerifyEmailParam contaisn the input parameters of verify email data
type GetVerifyEmailParam struct {
	SecretCode string `json:"secret_code"`
}

// CreateVersifyEmailResponse contains the result of Creating session data
type GetVerifyEmailResponse struct {
	VerifyEmail VerifyEmail `json:"verify_email"`
}

// GetTeacherParams db handler for api call to Get teacher details from the database
func (store *Store) GetVerifyEmail(ctx context.Context, arg GetVerifyEmailParam) (GetVerifyEmailResponse, error) {
	var result GetVerifyEmailResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.VerifyEmail, err = q.GetVerifyEmail(ctx, arg.SecretCode)

		if err != nil {
			return err
		}

		return nil
	})
	return result, err

}
