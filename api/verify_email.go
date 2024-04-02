package api

import (
	"bytes"
	"net/http"
	"os"
	"text/template"

	db "eduwave-back-end/db/sqlc"
	"eduwave-back-end/val"

	"eduwave-back-end/mail"

	"github.com/gin-gonic/gin"
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

func (server *Server) createVerifyEmail(ctx *gin.Context) {
	var req createVerifyEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Validate input data
	if err := val.ValidateEmail(req.Email); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Create VerifyEmailParam from request data
	createEmailParam := db.CreateVerifyEmailParams{
		UserName:   req.UserName,
		Email:      req.Email,
		SecretCode: req.SecretCode,
	}

	// Call the store method to create a verification email
	email, err := server.store.CreateVerifyEmail(ctx, createEmailParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Respond with the created verification email
	resp := newVerifyEmailResponse(email)
	ctx.JSON(http.StatusOK, resp)
}

type UpdateVerifyEmailRequest struct {
	EmailID    int64  `json:"email_id"`
	SecretCode string `json:"secret_code"`
}

func (server *Server) updateVerifyEmail(ctx *gin.Context) {
	var req UpdateVerifyEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	UpdateVerifyEmailParams := db.UpdateVerifyEmailParams{
		EmailID:    req.EmailID,
		SecretCode: req.SecretCode,
	}

	email, err := server.store.UpdateVerifyEmail(ctx, UpdateVerifyEmailParams)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := newVerifyEmailResponse(email)
	ctx.JSON(http.StatusOK, resp)
}

func createVerificationEmailHandler(c *gin.Context) {
	// Parse request body
	var request struct {
		UserName   string `json:"user_name"`
		Email      string `json:"email"`
		SecretCode string `json:"secret_code"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	// Read sender information from app.env file
	from := os.Getenv("SENDER_EMAIL")
	password := os.Getenv("SENDER_PASSWORD")

	// Receiver email address.
	to := []string{request.Email}

	// Prepare email content
	t, _ := template.New("emailTemplate").Parse(emailTemplate)
	var body bytes.Buffer
	err := t.Execute(&body, nil)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to execute email template"})
		return
	}

	// Send email
	err = mail.SendEmail(from, password, to, "Email Verification", body.String(), nil)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to send email"})
		return
	}

	emailSent = true

	// Respond with verification message
	c.String(200, "Email sent successfully")
}

func serveHTMLPageHandler(c *gin.Context) {
	// Serve the HTML page
	c.HTML(200, "verify.html", nil)
}

func verifyEmailHandler(c *gin.Context) {
	if emailSent {
		c.String(200, "Email verified")
	} else {
		c.String(400, "Email not yet sent")
	}
}

const emailTemplate = `Subject: Registartion in Eduwave"
MIME-version: 1.0;
Content-Type: text/html; charset="UTF-8";

<!DOCTYPE html>
<html>
<head>
    <title>Email Verification</title>
    <script>
        function redirectToVerification() {
            window.location.href = "http://localhost:8080/v1/verify_email?email_id=%d&secret_code=%s";
        }
    </script>
</head>
<body>
    <h3>Email Verification</h3>
    <p>Click the button to verify your account. Redirecting...</p>
    <button onclick="redirectToVerification()">Click here to continue</button>
</body>
</html>
`
