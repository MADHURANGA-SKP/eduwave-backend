// course enrolments controller
package api

import (
	"net/http"

	db "eduwave-back-end/db/sqlc"

	"github.com/gin-gonic/gin"
)

type listEnrolmentsRequest struct {
	UserID   int64 `json:"user_id"`
	CourseID int64 `json:"course_id"`
    Limit    int32 `json:"limit"`
    Offset   int32 `json:"offset"`
}

// @Summary List enrolments
// @Description List enrolments for a student in a course
// @ID list-enrolments
// @Accept  json
// @Produce  json
// @Param student_id query int true "Student ID"
// @Param course_id query int true "Course ID"
// @Param limit query int true "Limit"
// @Param offset query int true "Offset"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /courseEnrolments [get]
func (server *Server) listEnrolments(ctx *gin.Context) {
	var req listEnrolmentsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	enrolments, err := server.store.ListEnrolments(ctx, db.ListEnrolmentsParams{
		UserID: req.UserID,
		CourseID:  req.CourseID,
		Limit:     req.Limit,
		Offset:    req.Offset,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, enrolments)
}

type CreateEnrolmentsRequest struct {
	CourseID  int64 `json:"course_id"`
	RequestID int64 `json:"request_id"`
	UserID    int64 `json:"user_id"`
}

// @Summary Create a new enrolment
// @Description Create a new enrolment with the provided details
// @ID create-enrolment
// @Accept json
// @Produce json
// @Param request body CreateEnrolmentsRequest true "enrolment details"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /enrolments [post]
func (server *Server) CreateCourseEnrolment(ctx *gin.Context) {
	var req CreateEnrolmentsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateEnrolmentsParams{
		CourseID: req.CourseID,
		RequestID: req.RequestID,
		UserID: req.UserID,
	}

	assignment, err := server.store.CreateCourseEnrolments(ctx, db.CreateEnrolmentsParam(arg))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, assignment)
}
