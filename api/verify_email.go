package api

import (
	db "eduwave-back-end/db/sqlc"
)

var emailSent bool

type createVerifyEmailRequest struct {
	UserName   string `json:"user_name"`
	Email      string `json:"email"`
	SecretCode string `json:"secret_code"`
}

type verifyEmailResponse struct {
	EmailID    int64  `json:"email_id"`
	UserName   string `json:"user_name"`
	Email      string `json:"email"`
	SecretCode string `json:"secret_code"`
	IsUsed     bool   `json:"is_used"`
	CreatedAt  string `json:"created_at"`
	ExpiredAt  string `json:"expired_at"`
}

func newVerifyEmailResponse(email db.VerifyEmail) verifyEmailResponse {
	return verifyEmailResponse{
		EmailID:    email.EmailID,
		UserName:   email.UserName,
		Email:      email.Email,
		SecretCode: email.SecretCode,
		IsUsed:     email.IsUsed,
		CreatedAt:  email.CreatedAt.String(),
		ExpiredAt:  email.ExpiredAt.String(),
	}
}

// func (server *Server) createVerifyEmail(ctx *gin.Context) {
// 	var req createVerifyEmailRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	// Validate input data
// 	if err := val.ValidateEmail(req.Email); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	// Create VerifyEmailParam from request data
// 	createEmailParam := db.CreateVerifyEmailParam{
// 		UserName:   req.UserName,
// 		Email:      req.Email,
// 		SecretCode: req.SecretCode,
// 	}

// 	// Call the store method to create a verification email
// 	email, err := server.store.CreateVerifyEmail(ctx, createEmailParam)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	// Respond with the created verification email
// 	resp := newVerifyEmailResponse(email)
// 	ctx.JSON(http.StatusOK, resp)
// }

// func (server *Server) updateVerifyEmail(ctx *gin.Context) {
// 	var req createVerifyEmailRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	// Validate input data
// 	if err := val.ValidateEmail(req.Email); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	updateVerifyEmailParam := db.UpdateVerifyEmailParam{
// 		EmailID:    req.EmailID,
// 		SecretCode: req.SecretCode,
// 	}

// 	email, err := server.store.UpdateVerifyEmail(ctx, updateVerifyEmailParam)

// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	resp := newVerifyEmailResponse(email)
// 	ctx.JSON(http.StatusOK, resp)
// }
