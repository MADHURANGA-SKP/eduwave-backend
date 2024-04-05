// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type TypeResource string

const (
	TypeResourcePdf   TypeResource = "pdf"
	TypeResourceVideo TypeResource = "video"
	TypeResourceImage TypeResource = "image"
	TypeResourceDoc   TypeResource = "doc"
)

func (e *TypeResource) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TypeResource(s)
	case string:
		*e = TypeResource(s)
	default:
		return fmt.Errorf("unsupported scan type for TypeResource: %T", src)
	}
	return nil
}

type NullTypeResource struct {
	TypeResource TypeResource `json:"type_resource"`
	Valid        bool         `json:"valid"` // Valid is true if TypeResource is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTypeResource) Scan(value interface{}) error {
	if value == nil {
		ns.TypeResource, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.TypeResource.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTypeResource) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.TypeResource), nil
}

type Admin struct {
	AdminID        int64     `json:"admin_id"`
	UserID         int64     `json:"user_id"`
	Role           string    `json:"role"`
	UserName       string    `json:"user_name"`
	HashedPassword string    `json:"hashed_password"`
	FullName       string    `json:"full_name"`
	Email          string    `json:"email"`
	CreatedAt      time.Time `json:"created_at"`
}

type Assignment struct {
	AssignmentID   int64     `json:"assignment_id"`
	ResourceID     int64     `json:"resource_id"`
	Type           string    `json:"type"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	SubmissionDate time.Time `json:"submission_date"`
	CreatedAt      time.Time `json:"created_at"`
}

type Course struct {
	CourseID    int64     `json:"course_id"`
	TeacherID   int64     `json:"teacher_id"`
	Title       string    `json:"title"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Image       []byte    `json:"image"`
}

type CourseEnrolment struct {
	EnrolmentID int64 `json:"enrolment_id"`
	CourseID    int64 `json:"course_id"`
	RequestID   int64 `json:"request_id"`
	StudentID   int64 `json:"student_id"`
}

type CourseProgress struct {
	CourseprogressID int64  `json:"courseprogress_id"`
	EnrolmentID      int64  `json:"enrolment_id"`
	Progress         string `json:"progress"`
}

type Material struct {
	MaterialID  int64     `json:"material_id"`
	CourseID    int64     `json:"course_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type Permission struct {
	PermissionID   int64  `json:"permission_id"`
	PermissionName string `json:"permission_name"`
}

type Request struct {
	RequestID  int64        `json:"request_id"`
	StudentID  int64        `json:"student_id"`
	TeacherID  int64        `json:"teacher_id"`
	CourseID   int64        `json:"course_id"`
	IsActive   sql.NullBool `json:"is_active"`
	IsPending  sql.NullBool `json:"is_pending"`
	IsAccepted sql.NullBool `json:"is_accepted"`
	IsDeclined sql.NullBool `json:"is_declined"`
	CreatedAt  time.Time    `json:"created_at"`
}

type Resource struct {
	ResourceID int64        `json:"resource_id"`
	MaterialID int64        `json:"material_id"`
	Title      string       `json:"title"`
	Type       TypeResource `json:"type"`
	ContentUrl string       `json:"content_url"`
	CreatedAt  time.Time    `json:"created_at"`
}

type Role struct {
	RoleID   int64  `json:"role_id"`
	RoleName string `json:"role_name"`
}

type RolePermission struct {
	RoleID       int64 `json:"role_id"`
	PermissionID int64 `json:"permission_id"`
}

type Session struct {
	SessionID    uuid.UUID `json:"session_id"`
	UserName     string    `json:"user_name"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type Student struct {
	StudentID int64     `json:"student_id"`
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
}

type Submission struct {
	SubmissionID int64 `json:"submission_id"`
	AssignmentID int64 `json:"assignment_id"`
	StudentID    int64 `json:"student_id"`
}

type Teacher struct {
	TeacherID      int64     `json:"teacher_id"`
	UserID         int64     `json:"user_id"`
	Role           string    `json:"role"`
	FullName       string    `json:"full_name"`
	Email          string    `json:"email"`
	UserName       string    `json:"user_name"`
	HashedPassword string    `json:"hashed_password"`
	IsActive       bool      `json:"is_active"`
	CreatedAt      time.Time `json:"created_at"`
	Qualification  string    `json:"qualification"`
}

type TeachersUser struct {
	TeachersTeacherID int64 `json:"teachers_teacher_id"`
	UsersUserID       int64 `json:"users_user_id"`
}

type User struct {
	UserID            int64     `json:"user_id"`
	UserName          string    `json:"user_name"`
	Role              string    `json:"role"`
	HashedPassword    string    `json:"hashed_password"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	IsEmailVerified   bool      `json:"is_email_verified"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type UserRole struct {
	UserID int64 `json:"user_id"`
	RoleID int64 `json:"role_id"`
}

type VerifyEmail struct {
	EmailID    int64     `json:"email_id"`
	UserName   string    `json:"user_name"`
	Email      string    `json:"email"`
	SecretCode string    `json:"secret_code"`
	IsUsed     bool      `json:"is_used"`
	CreatedAt  time.Time `json:"created_at"`
	ExpiredAt  time.Time `json:"expired_at"`
}
