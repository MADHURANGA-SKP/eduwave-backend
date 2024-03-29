// submission controller
package api

import (
	"database/sql"
	"net/http"

	db "eduwave-back-end/db/sqlc"

	"github.com/gin-gonic/gin"
)

// getSubmissionRequest defines the request body structure for getting a submission
type getSubmissionRequest struct {
	AssignmentID int64 `uri:"assignment_id" binding:"required,min=1"`
	StudentID    int64 `uri:"student_id" binding:"required,min=1"`
}

// getSubmission retrieves a submission
func (server *Server) getSubmission(ctx *gin.Context) {
	var req getSubmissionRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetsubmissionsParams{
		AssignmentID: sql.NullInt64{Int64: req.AssignmentID, Valid: true},
		StudentID:    sql.NullInt64{Int64: req.StudentID, Valid: true},
	}

	submission, err := server.store.Getsubmissions(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, submission)
}

// listSubmissionsRequest defines the request body structure for listing submissions
type listSubmissionsRequest struct {
	AssignmentID int64 `uri:"assignment_id" binding:"required,min=1"`
	Limit        int32 `form:"limit" binding:"required,min=1,max=100"`
	Offset       int32 `form:"offset" binding:"required,min=0"`
}

// listSubmissions lists submissions
func (server *Server) listSubmissions(ctx *gin.Context) {
	var req listSubmissionsRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListsubmissionsParams{
		AssignmentID: sql.NullInt64{Int64: req.AssignmentID, Valid: true},
		Limit:        req.Limit,
		Offset:       req.Offset,
	}

	submissions, err := server.store.Listsubmissions(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, submissions)
}
