// course enrolments controller
package api

import (
	"net/http"

	db "eduwave-back-end/db/sqlc"

	"github.com/gin-gonic/gin"
)

type listEnrolmentsRequest struct {
	PageID   int32 `form:"page_id,min=1"`
	PageSize int32 `form:"page_size,min=10,max=10"`
}

// @Summary List enrolments
// @Description List enrolments for a student in a course
// @ID list-enrolments
// @Accept  json
// @Produce  json
// @Param limit query int true "Limit"
// @Param offset query int true "Offset"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /enrolments [get]
func (server *Server) listEnrolments(ctx *gin.Context) {
	var req listEnrolmentsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListEnrolmentsParams{
		Limit:    req.PageSize,
		Offset:   (req.PageID - 1) * req.PageSize,
	}

	enrolments, err := server.store.ListEnrolments(ctx, arg)
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
// @Description Create a new enrolment with the course
// @ID create-enrolment
// @Accept json
// @Produce json
// @Param request body CreateEnrolmentsRequest true "enrolment details"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /enrol [post]
func (server *Server) CreateCourseEnrolment(ctx *gin.Context) {
	var req CreateEnrolmentsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateEnrolmentsParams{
		CourseID:  req.CourseID,
		RequestID: req.RequestID,
		UserID:    req.UserID,
	}

	assignment, err := server.store.CreateCourseEnrolments(ctx, db.CreateEnrolmentsParam(arg))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, assignment)
}
