// submission controller
package api

import (
	"database/sql"
	"net/http"

	db "eduwave-back-end/db/sqlc"

	"github.com/gin-gonic/gin"
)

// CreateSubmissionRequest defines the request body structure for creating a submission
type CreateSubmissionRequest struct {
	AssignmentID int64 `json:"assignment_id"`
	UserID       int64 `json:"user_id"`
}

// @Summary Create a new Submission
// @Description Creates a new submission by user depend on assignment
// @Accept json
// @Produce json
// @Param request body CreateSubmissionRequest true "assignmnet_id and user_id"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /submissions [post]
// CreateSubmission creates a new course
func (server *Server) CreateSubmission(ctx *gin.Context) {
	var req CreateSubmissionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateSubmissionParams{
		AssignmentID: req.AssignmentID,
		UserID: req.UserID,
	}

	course, err := server.store.CreateSubmission(ctx, db.CreateSubmissionParam(arg))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, course)
}


// GetsubmissionsByAssignmentRequest defines the request body structure for getting a submission
type GetsubmissionsByAssignmentRequest struct {
	AssignmentID int64 `form:"assignment_id"`
}

// @Summary Get a submission By Assignment
// @Description Retrieves a submission by assignment and student ID
// @ID GetSubmissionsByAssignment
// @Produce json
// @Param assignment_id path int true "Assignment ID"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /submissions/byassignemnt [get]
// GetSubmissionsByAssignment retrieves a submission by assignment
func (server *Server) GetSubmissionsByAssignment(ctx *gin.Context) {
	var req GetsubmissionsByAssignmentRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetsubmissionsByAssignmentParam{
		AssignmentID: req.AssignmentID,
	}

	submission, err := server.store.GetsubmissionsByAssignment(ctx, arg)
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

// GetsubmissionsByUserRequest defines the request body structure for getting a submission
type GetsubmissionsByUserRequest struct {
	UserID int64 `form:"user_id"`
}

// @Summary Get a submission By User
// @Description Retrieves a submission by assignment and student ID
// @ID GetSubmissionsByUser
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /submissions/byuser [get]
// GetSubmissionsByUser retrieves a submission by user
func (server *Server) GetSubmissionsByUser(ctx *gin.Context) {
	var req GetsubmissionsByUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetsubmissionsByUserParam{
		UserID: req.UserID,
	}

	submission, err := server.store.GetsubmissionsByUser(ctx, arg)
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
	PageID   int32 `form:"page_id" binding:"required,min=1"`
    PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
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
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListsubmissionsParams{
		Limit:  req.PageSize,
        Offset: (req.PageID - 1) * req.PageSize,
	}

	submissions, err := server.store.Listsubmissions(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, submissions)
}
