// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
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

type UserRole string

const (
	UserRoleAdmin   UserRole = "admin"
	UserRoleStudent UserRole = "student"
	UserRoleTeacher UserRole = "teacher"
)

func (e *UserRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserRole(s)
	case string:
		*e = UserRole(s)
	default:
		return fmt.Errorf("unsupported scan type for UserRole: %T", src)
	}
	return nil
}

type NullUserRole struct {
	UserRole UserRole `json:"user_role"`
	Valid    bool     `json:"valid"` // Valid is true if UserRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserRole) Scan(value interface{}) error {
	if value == nil {
		ns.UserRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UserRole), nil
}

type Admin struct {
	AdminID  int64          `json:"admin_id"`
	UserName sql.NullString `json:"user_name"`
}

type Assignment struct {
	AssignmentID   int64         `json:"assignment_id"`
	CourseID       sql.NullInt64 `json:"course_id"`
	Type           string        `json:"type"`
	Title          string        `json:"title"`
	Description    string        `json:"description"`
	SubmissionDate time.Time     `json:"submission_date"`
}

type Course struct {
	CourseID    int64         `json:"course_id"`
	TeacherID   sql.NullInt64 `json:"teacher_id"`
	Title       string        `json:"title"`
	Type        string        `json:"type"`
	Description string        `json:"description"`
	CreatedAt   time.Time     `json:"created_at"`
}

type CourseEnrolment struct {
	EnrolmentID int64         `json:"enrolment_id"`
	CourseID    sql.NullInt64 `json:"course_id"`
	RequestID   sql.NullInt64 `json:"request_id"`
	StudentID   sql.NullInt64 `json:"student_id"`
}

type CourseProgress struct {
	CourseprogressID int64         `json:"courseprogress_id"`
	EnrolmentID      sql.NullInt64 `json:"enrolment_id"`
	Progress         string        `json:"progress"`
}

type Request struct {
	RequestID  int64         `json:"request_id"`
	StudentID  sql.NullInt64 `json:"student_id"`
	TeacherID  sql.NullInt64 `json:"teacher_id"`
	CourseID   sql.NullInt64 `json:"course_id"`
	IsActive   sql.NullBool  `json:"is_active"`
	IsPending  sql.NullBool  `json:"is_pending"`
	IsAccepted sql.NullBool  `json:"is_accepted"`
	IsDeclined sql.NullBool  `json:"is_declined"`
}

type Resource struct {
	ResourceID   int64         `json:"resource_id"`
	CourseID     sql.NullInt64 `json:"course_id"`
	AssignmentID sql.NullInt64 `json:"assignment_id"`
	Title        string        `json:"title"`
	Type         TypeResource  `json:"type"`
	ContentUrl   string        `json:"content_url"`
}

type Student struct {
	StudentID int64          `json:"student_id"`
	UserName  sql.NullString `json:"user_name"`
}

type Submission struct {
	SubmissionID int64         `json:"submission_id"`
	AssignmentID sql.NullInt64 `json:"assignment_id"`
	StudentID    sql.NullInt64 `json:"student_id"`
}

type Teacher struct {
	TeacherID      int64          `json:"teacher_id"`
	AdminID        sql.NullInt64  `json:"admin_id"`
	FullName       string         `json:"full_name"`
	Email          string         `json:"email"`
	UserName       sql.NullString `json:"user_name"`
	HashedPassword string         `json:"hashed_password"`
	IsActive       bool           `json:"is_active"`
	CreatedAt      time.Time      `json:"created_at"`
}

type User struct {
	UserName          string    `json:"user_name"`
	Role              UserRole  `json:"role"`
	HashedPassword    string    `json:"hashed_password"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	IsEmailVerified   bool      `json:"is_email_verified"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}