// course enrolments controller
package api

import (
	"errors"
	"net/http"

	db "eduwave-back-end/db/sqlc"

	"github.com/gin-gonic/gin"
)

type listEnrolmentsRequest struct {
	CourseID int64 `form:"course_id"`
	PageID   int32 `form:"page_id,min=1"`
	PageSize int32 `form:"page_size,min=10,max=100"`
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
func (server *Server) ListEnrolments(ctx *gin.Context) {
	var req listEnrolmentsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListEnrolmentsParams{
		CourseID: req.CourseID,
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


type GetEnrolmentRequest struct {
	UserID   int64 `form:"user_id"`
    CourseID int64 `form:"course_id"`
}
// @Summary Get User enrolment details by Userid and courseid
// @Description Get an user  enrolment detials by user id and courseid
// @ID get-user
// @Accept json
// @Produce json
// @Param user_id path int true "user_id" 
// @Param course_id path int true "course_id" 
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /enrolment/get [get]
func (server *Server) GetEnrolment(ctx *gin.Context) {
	var req GetEnrolmentRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.GetEnrolmentParam{
		UserID: req.UserID,
		CourseID: req.CourseID,
	}

	enrolment, err := server.store.GetEnrolment(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, enrolment)
}