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
	AssignmentID int64 `json:"assignment_id"`
    StudentID    int64 `json:"student_id"`
}

// @Summary Get a submission
// @Description Retrieves a submission by assignment and student ID
// @ID getSubmission
// @Produce json
// @Param assignment_id path int true "Assignment ID"
// @Param student_id path int true "Student ID"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /submissions/{assignment_id}/{student_id} [get]
// getSubmission retrieves a submission
func (server *Server) getSubmission(ctx *gin.Context) {
	var req getSubmissionRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetsubmissionsParams{
		AssignmentID: req.AssignmentID,
		StudentID: req.StudentID,
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

// @Summary List submissions
// @Description Lists submissions for a given assignment ID
// @ID listSubmissions
// @Produce json
// @Param assignment_id path int true "Assignment ID"
// @Param limit query int true "Limit" minimum(1) maximum(100)
// @Param offset query int true "Offset" minimum(0)
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /submissions [get]
// listSubmissions lists submissions
func (server *Server) listSubmissions(ctx *gin.Context) {
	var req listSubmissionsRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListsubmissionsParams{
		AssignmentID: req.AssignmentID,
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
