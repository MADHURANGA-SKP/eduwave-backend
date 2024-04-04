package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

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