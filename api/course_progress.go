// course progreass controller
package api

import (
	"database/sql"
	"net/http"

	db "eduwave-back-end/db/sqlc"

	"github.com/gin-gonic/gin"
)

// createCourseProgressRequest defines the request body structure for creating course progress
type createCourseProgressRequest struct {
	Progress    string `json:"progress" binding:"required"`
}


// createCourseProgress creates course progress
func (server *Server) createCourseProgress(ctx *gin.Context) {
	var req createCourseProgressRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCourseProgressPram{
		Progress:    req.Progress,
	}

	courseProgress, err := server.store.CreateCourseProgress(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, courseProgress)
}

// getCourseProgressRequest defines the request body structure for getting course progress by ID
type getCourseProgressRequest struct {
	CourseprogressID int64         `json:"courseprogress_id"`
	EnrolmentID      sql.NullInt64 `json:"enrolment_id"`
}

// getCourseProgress returns course progress by ID
func (server *Server) getCourseProgress(ctx *gin.Context) {
	var req getCourseProgressRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetCourseProgressParams{
		CourseprogressID: req.CourseprogressID,
		EnrolmentID:      req.EnrolmentID,
	}

	courseProgress, err := server.store.GetCourseProgress(ctx, db.GetCourseProgressParam(arg))
	if err != nil {
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, courseProgress)
}

// listCourseProgressRequest defines the request body structure for listing course progress
type listCourseProgressRequest struct {
	EnrolmentID sql.NullInt64 `form:"enrolment_id" binding:"required,min=1"`
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

	ctx.JSON(http.StatusOK, courseProgressList)
}
