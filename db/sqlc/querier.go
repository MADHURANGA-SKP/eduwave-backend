// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateAssignment(ctx context.Context, arg CreateAssignmentParams) (Assignment, error)
	CreateCourseProgress(ctx context.Context, progress string) (CourseProgress, error)
	CreateRequest(ctx context.Context, arg CreateRequestParams) (Request, error)
	CreateResource(ctx context.Context, arg CreateResourceParams) (Resource, error)
	CreateStudent(ctx context.Context, userName sql.NullString) (Student, error)
	CreateTeacher(ctx context.Context, arg CreateTeacherParams) (Teacher, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAdmin(ctx context.Context, adminID int64) error
	DeleteAssignment(ctx context.Context, assignmentID int64) error
	DeleteRequest(ctx context.Context, requestID int64) error
	DeleteResource(ctx context.Context, arg DeleteResourceParams) error
	DeleteStudent(ctx context.Context, studentID int64) error
	DeleteTeacher(ctx context.Context, teacherID int64) error
	GetAdmin(ctx context.Context, adminID int64) (Admin, error)
	GetAssignment(ctx context.Context, courseID sql.NullInt64) (Assignment, error)
	GetCourseProgress(ctx context.Context, arg GetCourseProgressParams) (CourseProgress, error)
	GetRequest(ctx context.Context, requestID int64) (Request, error)
	GetResource(ctx context.Context, arg GetResourceParams) (Resource, error)
	GetStudent(ctx context.Context, studentID int64) (Student, error)
	GetTeacher(ctx context.Context, teacherID int64) (Teacher, error)
	GetUser(ctx context.Context, userName string) (User, error)
	Getsubmissions(ctx context.Context, arg GetsubmissionsParams) (Submission, error)
	ListCourseProgress(ctx context.Context, arg ListCourseProgressParams) ([]CourseProgress, error)
	ListEnrolments(ctx context.Context, arg ListEnrolmentsParams) ([]CourseEnrolment, error)
	ListRequest(ctx context.Context, arg ListRequestParams) ([]Request, error)
	ListResource(ctx context.Context, arg ListResourceParams) ([]Resource, error)
	ListStudents(ctx context.Context, arg ListStudentsParams) ([]Student, error)
	ListTeacher(ctx context.Context, arg ListTeacherParams) ([]Teacher, error)
	Listsubmissions(ctx context.Context, arg ListsubmissionsParams) ([]Submission, error)
	UpdateAdmin(ctx context.Context, arg UpdateAdminParams) (Admin, error)
	UpdateAssignment(ctx context.Context, arg UpdateAssignmentParams) (Assignment, error)
	UpdateRequests(ctx context.Context, arg UpdateRequestsParams) (Request, error)
	UpdateResource(ctx context.Context, arg UpdateResourceParams) (Resource, error)
	UpdateStudent(ctx context.Context, arg UpdateStudentParams) (Student, error)
	UpdateTeacher(ctx context.Context, arg UpdateTeacherParams) (Teacher, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
