package worker

import (
	db "eduwave-back-end/db/sqlc"
	"eduwave-back-end/mail"
)

// SimpleTaskDistributor represents a task distributor
type SimpleTaskDistributor struct {
	store  db.Store
	mailer mail.EmailSender
}

const (
	QueueCritical       = "critical"
	QueueDefault        = "default"
	TaskSendVerifyEmail = "send_verify_email"
)

// PayloadSendVerifyEmail represents the payload for sending a verification email task
type PayloadSendVerifyEmail struct {
	Username string `json:"username"`
}
