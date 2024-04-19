// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateAssignment(ctx context.Context, arg CreateAssignmentParams) (Assignment, error)
	CreateCourseProgress(ctx context.Context, arg CreateCourseProgressParams) (CourseProgress, error)
	CreateCourses(ctx context.Context, arg CreateCoursesParams) (Course, error)
	CreateEnrolments(ctx context.Context, arg CreateEnrolmentsParams) (CourseEnrolment, error)
	CreateMaterial(ctx context.Context, arg CreateMaterialParams) (Material, error)
	CreateRequest(ctx context.Context, arg CreateRequestParams) (Request, error)
	CreateResource(ctx context.Context, arg CreateResourceParams) (Resource, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateSubmission(ctx context.Context, arg CreateSubmissionParams) (Submission, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateVerifyEmail(ctx context.Context, arg CreateVerifyEmailParams) (VerifyEmail, error)
	DeleteAssignment(ctx context.Context, assignmentID int64) error
	DeleteCourses(ctx context.Context, courseID int64) error
	DeleteMaterial(ctx context.Context, materialID int64) error
	DeleteRequest(ctx context.Context, requestID int64) error
	DeleteResource(ctx context.Context, resourceID int64) error
	DeleteUsers(ctx context.Context, userID int64) error
	GetAssignment(ctx context.Context, assignmentID int64) (Assignment, error)
	GetCourseProgress(ctx context.Context, arg GetCourseProgressParams) (CourseProgress, error)
	GetCourses(ctx context.Context, courseID int64) (Course, error)
	GetMaterial(ctx context.Context, materialID int64) (Material, error)
	GetRequest(ctx context.Context, arg GetRequestParams) (Request, error)
	GetResource(ctx context.Context, resourceID int64) (Resource, error)
	GetSession(ctx context.Context, sessionID uuid.UUID) (Session, error)
	// -- name: GetUser :one
	// SELECT users.user_name AS user_username, teachers.user_name AS teacher_username, admins.user_name AS admin_username
	// FROM users
	// LEFT JOIN teachers ON users.user_id = teachers.user_id
	// LEFT JOIN admins ON users.user_id = admins.user_id
	// WHERE
	//     users.user_name = $1 OR teachers.user_name = $1 OR admins.user_name = $1;
	GetUser(ctx context.Context, userName string) (User, error)
	GetVerifyEmail(ctx context.Context, secretCode string) (VerifyEmail, error)
	GetsubmissionsByAssignment(ctx context.Context, assignmentID int64) (Submission, error)
	GetsubmissionsByUser(ctx context.Context, userID int64) (Submission, error)
	ListCourseProgress(ctx context.Context, arg ListCourseProgressParams) ([]CourseProgress, error)
	ListCourses(ctx context.Context, arg ListCoursesParams) ([]Course, error)
	ListCoursesByUser(ctx context.Context, arg ListCoursesByUserParams) ([]Course, error)
	ListEnrolments(ctx context.Context, arg ListEnrolmentsParams) ([]CourseEnrolment, error)
	ListMaterial(ctx context.Context, arg ListMaterialParams) ([]Material, error)
	ListRequest(ctx context.Context, arg ListRequestParams) ([]Request, error)
	ListResource(ctx context.Context, arg ListResourceParams) ([]Resource, error)
	ListUser(ctx context.Context, arg ListUserParams) ([]User, error)
	Listsubmissions(ctx context.Context, arg ListsubmissionsParams) ([]Submission, error)
	UpdateAssignment(ctx context.Context, arg UpdateAssignmentParams) (Assignment, error)
	UpdateCourses(ctx context.Context, arg UpdateCoursesParams) (Course, error)
	UpdateMaterial(ctx context.Context, arg UpdateMaterialParams) (Material, error)
	UpdateRequests(ctx context.Context, arg UpdateRequestsParams) (Request, error)
	UpdateResource(ctx context.Context, arg UpdateResourceParams) (Resource, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateVerifyEmail(ctx context.Context, arg UpdateVerifyEmailParams) (VerifyEmail, error)
}

var _ Querier = (*Queries)(nil)
