package api

import (
	"net/http"

	db "eduwave-back-end/db/sqlc"

	"github.com/gin-gonic/gin"
)

// createCourseProgressRequest defines the request body structure for creating course progress
type createCourseProgressRequest struct {
	EnrolmentID int64  `json:"enrolment_id" binding:"required"`
	Progress    string `json:"progress" binding:"required"`
}

// courseProgressResponse defines the response body structure for course progress
type courseProgressResponse struct {
	CourseProgressID int64  `json:"course_progress_id"`
	EnrolmentID      int64  `json:"enrolment_id"`
	Progress         string `json:"progress"`
}

// createCourseProgress creates course progress
func (server *Server) createCourseProgress(ctx *gin.Context) {
	var req createCourseProgressRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCourseProgressParams{
		EnrolmentID: req.EnrolmentID,
		Progress:    req.Progress,
	}

	courseProgress, err := server.store.CreateCourseProgress(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := courseProgressResponse{
		CourseProgressID: courseProgress.CourseprogressID,
		EnrolmentID:      courseProgress.EnrolmentID,
		Progress:         courseProgress.Progress,
	}
	ctx.JSON(http.StatusOK, rsp)
}

// getCourseProgressRequest defines the request body structure for getting course progress by ID
type getCourseProgressRequest struct {
	ID          int64 `uri:"id" binding:"required,min=1"`
	EnrolmentID int64 `uri:"enrolment_id" binding:"required,min=1"`
}

// getCourseProgress returns course progress by ID
func (server *Server) getCourseProgress(ctx *gin.Context) {
	var req getCourseProgressRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetCourseProgressParams{
		CourseprogressID: req.ID,
		EnrolmentID:      req.EnrolmentID,
	}

	courseProgress, err := server.store.GetCourseProgress(ctx, arg)
	if err != nil {
		if db.IsErrNoRows(err) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := courseProgressResponse{
		CourseProgressID: courseProgress.CourseprogressID,
		EnrolmentID:      courseProgress.EnrolmentID,
		Progress:         courseProgress.Progress,
	}
	ctx.JSON(http.StatusOK, rsp)
}

// listCourseProgressRequest defines the request body structure for listing course progress
type listCourseProgressRequest struct {
	EnrolmentID int64 `form:"enrolment_id" binding:"required,min=1"`
	Limit       int32 `form:"limit" binding:"required,min=1,max=100"`
	Offset      int32 `form:"offset" binding:"required,min=0"`
}

// listCourseProgress returns a list of course progress
func (server *Server) listCourseProgress(ctx *gin.Context) {
	var req listCourseProgressRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListCourseProgressParams{
		EnrolmentID: req.EnrolmentID,
		Limit:       req.Limit,
		Offset:      req.Offset,
	}

	courseProgressList, err := server.store.ListCourseProgress(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var resp []courseProgressResponse
	for _, courseProgress := range courseProgressList {
		resp = append(resp, courseProgressResponse{
			CourseProgressID: courseProgress.CourseprogressID,
			EnrolmentID:      courseProgress.EnrolmentID,
			Progress:         courseProgress.Progress,
		})
	}

	ctx.JSON(http.StatusOK, resp)
}
