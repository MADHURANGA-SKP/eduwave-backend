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
	Email      string `json:"email"`
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

	if req.UserName != user.User.UserName {
		ctx.JSON(http.StatusUnauthorized, gin.H{"messege": "invalid username"})
		return
	}

	// Fetch the user's record along with the secret code from the database
	verifyEmail, err := server.store.GetVerifyEmail(ctx, db.GetVerifyEmailParam{
		SecretCode: req.SecretCode,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Verify the provided secret code
	if req.SecretCode == verifyEmail.VerifyEmail.SecretCode {
		ctx.JSON(http.StatusOK, gin.H{"message": "Successfully Registered"})
		return
	}

	// Update the user's record in the database to mark the email as verified
	if req.SecretCode == verifyEmail.VerifyEmail.SecretCode {
		// Create verify email data in the database
		arg := db.UpdateVerifyEmailParams{
			IsUsed: true,
		}
		_, err = server.store.UpdateVerifyEmail(ctx, db.UpdateVerifyEmailParam(arg))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, user)
	}
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
<html>
<head>
    <title>Email Verification</title>
    <script>
        function redirectToVerification(emailID, secretCode) {
            var xhr = new XMLHttpRequest();
			var url = "/v1/verify_email";
			var params = "email_id=" + emailID + "&secret_code=" + secretCode;
			xhr.open("POST", url, true);
			xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
			xhr.onreadystatechange = function () {
				if (xhr.readyState == 4 && xhr.status == 200) {
					alert(xhr.responseText);
				}
			}
			xhr.send(params);
        }
    </script>
</head>
<body>
    <h3>Email Verification</h3>
	<p>verify your account with the  secret code: <strong>{{.SecretCode}}</strong></p>
</body>
</html>
`
