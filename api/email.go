package api

import (
	"bytes"
	db "eduwave-back-end/db/sqlc"
	"html/template"
	"math/rand"
	"net/http"
	"net/smtp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func generateSecretCode() string {
	rand.Seed(time.Now().UnixNano())
	min := 100000 // Minimum value for a 6-digit code
	max := 999999 // Maximum value for a 6-digit code
	code := rand.Intn(max-min+1) + min
	return strconv.Itoa(code)
}

func sendUserVerificationEmail(to string) (string, error) {
	// Generate a secret code
	secretCode := generateSecretCode()

	// Sender data.
	from := "kumarihkbk.20@itfac.mrt.ac.lk"
	password := "wyoi mstq fcum fuqy"

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Prepare email content
	t, err := template.New("emailTemplate").Parse(emailTemplate)
	if err != nil {
		return "", err
	}

	var body bytes.Buffer
	err = t.Execute(&body, struct {
		EmailID    int64
		SecretCode string
	}{
		EmailID:    123, // Replace with actual email ID
		SecretCode: secretCode,
	})
	if err != nil {
		return "", err
	}

	// Send email with properly formatted HTML content
	err = sendEmail(from, password, smtpHost, smtpPort, []string{to}, body.Bytes(), "Email Verification")
	if err != nil {
		return "", err
	}

	return secretCode, nil
}

type createVerifyEmailRequest struct {
	UserName   string `json:"user_name"`
	SecretCode string `json:"secret_code"`
}

// VerifyEmailHandler handles the verification request sent by the user clicking the button in the email
func (server *Server) VerifyEmailHandler(ctx *gin.Context) {
	var req createVerifyEmailRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, db.GetUserParam{
		UserName: req.UserName,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if req.UserName != user.User.UserName {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "invalid username"})
		return
	}

	// Fetch the user's record along with the secret code from the database
	verifyEmail, err := server.store.GetVerifyEmail(ctx, db.GetVerifyEmailParam{
		UserName:   req.UserName,
		SecretCode: req.SecretCode,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Verify the provided secret code
	if req.SecretCode == verifyEmail.VerifyEmail.SecretCode {
		ctx.JSON(http.StatusOK, gin.H{"message": "Successfully Registered"})

	}

	// Update the user's record in the database to mark the email as verified
	if req.SecretCode == verifyEmail.VerifyEmail.SecretCode {
		// Update IsUsed to true
		arg := db.UpdateVerifyEmailParams{
			IsUsed:     true,
			EmailID:    verifyEmail.VerifyEmail.EmailID,
			SecretCode: verifyEmail.VerifyEmail.SecretCode,
		}
		_, err = server.store.UpdateVerifyEmail(ctx, db.UpdateVerifyEmailParam(arg))

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		//update IsEmailVerified to true
		_, err = server.store.UpdateUser(ctx, db.UpdateUserParam{
			FullName:        user.User.FullName,
			UserName:        req.UserName,
			Email:           user.User.Email,
			HashedPassword:  user.User.HashedPassword,
			IsEmailVerified: true,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Email verification successful"})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid secret code"})
}

func sendEmail(from, password, smtpHost, smtpPort string, to []string, body []byte, subject string) error {
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Compose the email message
	msg := "Subject: " + subject + "\r\n" +
		"MIME-version: 1.0;\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
		string(body)

	// Send the email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(msg))
	return err
}

// HTML email template
const emailTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Email Verification</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f5f5f5;
            margin: 0;
            padding: 20px;
        }
        .container {
            max-width: 600px;
            margin: 0 auto;
            background-color: #fff;
            padding: 20px;
            border-radius: 8px;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
        }
        h1 {
            text-align: center;
            color: #333;
        }
        p {
            font-size: 18px;
            line-height: 1.5;
            margin-bottom: 20px;
            color: #666;
        }
        .verification-code {
            display: inline-block;
            padding: 8px 16px;
            background-color: #007bff;
            color: #fff;
            font-size: 20px;
            border-radius: 4px;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Email Verification</h1>
        <p>Welcome! To complete your account setup, please verify your email address by entering the secret code below:</p>
        <p>Secret Code: <span class="verification-code">{{.SecretCode}}</span></p>
        <p>If you didn't request this verification, please ignore this email.</p>
        <p>Thank you!</p>
    </div>
</body>
</html>
`
