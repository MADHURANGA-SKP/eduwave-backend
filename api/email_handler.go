package api

import (
	"net/http"

	db "eduwave-back-end/db/sqlc"
	"eduwave-back-end/val"

	"github.com/gin-gonic/gin"
)

type VerifyEmailRequest struct {
	EmailId    int64  `json:"email_id"`
	SecretCode string `json:"secret_code"`
}
type FieldViolation struct {
	Field       string `json:"field"`
	Description string `json:"description"`
}

func (server *Server) VerifyEmail(ctx *gin.Context) {
	var req VerifyEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	violations := validateVerifyEmailRequest(&req)
	if len(violations) > 0 {
		ctx.JSON(http.StatusBadRequest, fieldViolationResponse(violations))
		return
	}

	txResult, err := server.store.VerifyEmailTx(ctx, db.VerifyEmailTxParams{
		EmailId:    req.EmailId,
		SecretCode: req.SecretCode,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := struct {
		IsVerified bool `json:"is_verified"`
	}{
		IsVerified: txResult.User.IsEmailVerified,
	}
	ctx.JSON(http.StatusOK, rsp)
}

func fieldViolationResponse(violations []*FieldViolation) gin.H {
	response := make(gin.H)
	for _, v := range violations {
		response[v.Field] = v.Description
	}
	return response
}

func validateVerifyEmailRequest(req *VerifyEmailRequest) (violations []*FieldViolation) {
	if err := val.ValidateEmailId(req.EmailId); err != nil {
		violations = append(violations, fieldViolation("email_id", err.Error()))
	}

	if err := val.ValidateSecretCode(req.SecretCode); err != nil {
		violations = append(violations, fieldViolation("secret_code", err.Error()))
	}

	return violations
}

func fieldViolation(field, description string) *FieldViolation {
	return &FieldViolation{
		Field:       field,
		Description: description,
	}
}
