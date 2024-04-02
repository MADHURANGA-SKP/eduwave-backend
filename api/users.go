package api

import (
	"errors"
	"net/http"
	"time"

	db "eduwave-back-end/db/sqlc"
	"eduwave-back-end/util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createUserRequest struct {
	UserName       string `json:"user_name"`
	Role           string `json:"role"`
	FullName       string `json:"full_name"`
	HashedPassword string `json:"hashed_password"`
	Email          string `json:"email"`
}

type userResponse struct {
	Username          string    `json:"user_name"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:          user.UserName,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}

// @Summary Create a new user
// @Description Create a new user with the provided details
// @ID create-user
// @Accept  json
// @Produce  json
// @Param request body createUserRequest true "User creation request"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /signup [post]
func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	user, err := server.store.CreateUser(ctx, db.CreateUserParam{
		UserName: req.UserName,
		Role: req.Role,
		FullName: req.FullName,
		HashedPassword:hashedPassword,
		Email: req.Email,
	})
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolations {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newUserResponse(user.User)
	ctx.JSON(http.StatusOK, rsp)
}

// type UpdateUserRequest struct {
// 	HashedPassword sql.NullString `json:"hashed_password"`
// 	FullName       sql.NullString `json:"full_name"`
// 	Email          sql.NullString `json:"email"`
// 	UserName       string         `json:"user_name"`
// }

// // updateStudent updates a student by ID
// func (server *Server) UpdateUser(ctx *gin.Context){
// 	var req UpdateUserRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 	}

// 	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
// 	arg := db.UpdateUserParams{
// 		HashedPassword: req.HashedPassword,
// 		FullName: req.FullName,
// 		Email: req.Email,
// 		UserName: authPayload.Username,
// 	}

// 	// Call the database store function to update the student
// 	updatedStudent, err := server.store.UpdateUser(ctx,)
// 	if err != nil {
// 		if err == sql.ErrNoRows{
// 			ctx.JSON(http.StatusNotFound, errorResponse(err))
// 		}

// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, updatedStudent)
// }

type loginUserRequest struct {
	UserName       string `json:"user_name"`
	HashedPassword string `json:"hashed_password"`
}

type loginUserResponse struct {
	SessionID             uuid.UUID    `json:"session_id"`
	AccessToken           string       `json:"access_token"`
	AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
	RefreshToken          string       `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
	User                  userResponse `json:"user"`
}

// @Summary Log in user
// @Description Log in a user with the provided credentials
// @ID login-user
// @Accept  json
// @Produce  json
// @Param request body loginUserRequest true "Login request"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /login [post]
func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest


	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, db.GetUserParam{
		UserName: req.UserName,
	})
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.HashedPassword, user.User.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		user.User.UserName,
		user.User.Role,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		user.User.UserName,
		user.User.Role,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	session, err := server.store.CreateSession(ctx, db.CreateSessionParam{
		SessionID:    refreshPayload.ID,
		UserName:     user.User.UserName,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginUserResponse{
		SessionID:             session.Session.SessionID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  newUserResponse(user.User),
	}

	ctx.JSON(http.StatusOK, rsp)
}
