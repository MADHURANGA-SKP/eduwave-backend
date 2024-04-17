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
	UserName       string `json:"user_name" binding:"required,alphanum"`
	FullName       string `json:"full_name"`
	HashedPassword string `json:"hashed_password" binding:"required,min=6"`
	Email          string `json:"email" binding:"required,email"`
	Role           string `json:"role"`
	Qualification  string `json:"qualification"`
}

type userResponse struct {
	Username      string    `json:"user_name"`
	FullName      string    `json:"full_name"`
	Email         string    `json:"email"`
	Role          string    `json:"role"`
	Qualification string    `json:"qualification"`
	CreatedAt     time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:      user.UserName,
		FullName:      user.FullName,
		Email:         user.Email,
		Role:          user.Role,
		Qualification: user.Qualification,
		CreatedAt:     user.CreatedAt,
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
		UserName:       req.UserName,
		FullName:       req.FullName,
		HashedPassword: hashedPassword,
		Email:          req.Email,
		Role:           "student",
		Qualification:  req.Qualification,
	})
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolations {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Send verification email
	secretCode, err := sendUserVerificationEmail(req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Create verify email data in the database
	verifyEmail := db.CreateVerifyEmailParam{
		UserName:   req.UserName,
		Email:      req.Email,
		SecretCode: secretCode,
	}
	_, err = server.store.CreateVerifyEmail(ctx, verifyEmail)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Respond to the client
	rsp := newUserResponse(user.User)
	ctx.JSON(http.StatusOK, gin.H{
		"user":       rsp,
		"secretCode": secretCode, // Send the secret code back to the client
	})
}

type UpdateUserRequest struct {
	HashedPassword    string    `json:"hashed_password"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	IsEmailVerified   bool      `json:"is_email_verified"`
	UserName          string    `json:"user_name"`
}

// @Summary Update a user
// @Description Updates a user with provided details
// @Accept json
// @Produce json
// @Param request body UpdateUserRequest true "Updated user details"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /user/edit [Patch]
// Updateuser updates the selected user
// updateStudent updates a student by ID
func (server *Server) UpdateUser(ctx *gin.Context) {
	var req UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	username, err := server.store.GetUser(ctx, db.GetUserParam{
		UserName: req.UserName,
	},
	)

	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if username.User.UserName != req.UserName {
		ctx.JSON(http.StatusForbidden, "connot update other user's info")
		return
	}

	hashedPassword, err := util.HashPassword(req.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := &db.UpdateUserParams{
		HashedPassword:    hashedPassword,
		PasswordChangedAt: time.Now(),
		FullName:          req.FullName,
		Email:             req.Email,
		IsEmailVerified:   req.IsEmailVerified,
		UserName:          req.UserName,
	}

	user, err := server.store.UpdateUser(ctx, db.UpdateUserParam{
		HashedPassword:    arg.HashedPassword,
		PasswordChangedAt: arg.PasswordChangedAt,
		FullName:          arg.FullName,
		Email:             arg.Email,
		IsEmailVerified:   arg.IsEmailVerified,
		UserName:          arg.UserName,
	})
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolations {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	rsp := newUserResponse(user.User)

	ctx.JSON(http.StatusOK, rsp)
}

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
	},
	)

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
		user.User.UserID,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		user.User.UserName,
		user.User.Role,
		user.User.UserID,
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

type createAdminRequest struct {
	UserName       string `json:"user_name" binding:"required,alphanum"`
	FullName       string `json:"full_name"`
	HashedPassword string `json:"hashed_password" binding:"required,min=6"`
	Email          string `json:"email" binding:"required,email"`
	Role           string `json:"role"`
	Qualification  string `json:"qualification"`
}

type userAdminResponse struct {
	Username      string    `json:"user_name"`
	FullName      string    `json:"full_name"`
	Email         string    `json:"email"`
	Role          string    `json:"role"`
	Qualification string    `json:"qualification"`
	CreatedAt     time.Time `json:"created_at"`
}

func newAdminResponse(user db.User) userAdminResponse {
	return userAdminResponse{
		Username:      user.UserName,
		FullName:      user.FullName,
		Email:         user.Email,
		Role:          user.Role,
		Qualification: user.Qualification,
		CreatedAt:     user.CreatedAt,
	}
}

// @Summary Create a admin's user
// @Description Create a new admin's user with the provided details
// @ID create admin user
// @Accept  json
// @Produce  json
// @Param request body createUserRequest true "User creation request"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /admin/signup [post]
func (server *Server) createAdminUser(ctx *gin.Context) {
	var req createAdminRequest
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
		UserName:       req.UserName,
		FullName:       req.FullName,
		HashedPassword: hashedPassword,
		Email:          req.Email,
		Role:           req.Role,
		Qualification:  req.Qualification,
	})
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolations {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newAdminResponse(user.User)
	ctx.JSON(http.StatusOK, rsp)
}

// ListUserRequest contains the impurt parameters for list rolsbased user data
type ListUserRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// @Summary ListUser
// @Description ListUser with the provided admin based
// @ID list-user
// @Accept  json
// @Produce  json
// @Param request body ListUserRequest true "admin list request"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /listadmin [get]
func (server *Server) ListUser(ctx *gin.Context) {
	var req ListUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.ListUserParams{
		Role:   "admin",
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	userlist, err := server.store.ListUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, userlist)
}

// ListUserStudentRequest contains the impurt parameters for list rolebased user data
type ListUserStudentRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// @Summary ListUserStudent
// @Description ListUserStudent with the provided student based
// @ID list-student
// @Accept  json
// @Produce  json
// @Param request body ListUserStudentRequest true "student list request"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /liststudent [get]
func (server *Server) ListUserStudent(ctx *gin.Context) {
	var req ListUserStudentRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.ListUserParams{
		Role:   "student",
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	userlist, err := server.store.ListUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, userlist)
}

// ListUserTeacherRequest contains the impurt parameters for list rolebased user data
type ListUserTeacherRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// @Summary ListUserTeacher
// @Description ListUserTeacher with the provided teacher based
// @ID list-teacher
// @Accept  json
// @Produce  json
// @Param request body ListUserTeacherRequest true "teacher list request"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /listteacher [get]
func (server *Server) ListUserTeacher(ctx *gin.Context) {
	var req ListUserTeacherRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.ListUserParams{
		Role:   "teacher",
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	userlist, err := server.store.ListUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, userlist)
}


type GetUserRequest struct {
	UserName string         `form:"user_name"`
}
// @Summary Get an user details by username
// @Description Get an user by its username
// @ID get-user
// @Accept json
// @Produce json
// @Param user_name 
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /user/get [get]
func (server *Server) GetUser(ctx *gin.Context) {
	var req GetUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.GetUserParam{UserName: req.UserName}

	assignment, err := server.store.GetUser(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, assignment)
}

func(server *Server) GetSample(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{"message": "hellow world"})
}
