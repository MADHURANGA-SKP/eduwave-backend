// Assignemts cotroller created
package api

import (
	db "eduwave-back-end/db/sqlc"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new assignment
// @Description Create a new assignment with the provided details
// @ID create-assignment
// @Accept json
// @Produce json
// @Param request body createAssignmentRequest true "Assignment details"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /assignments [post]
type createAssignmentRequest struct {
	Type           string    `json:"type" binding:"required"`
	Title          string    `json:"title" binding:"required"`
	Description    string    `json:"description" binding:"required"`
	SubmissionDate time.Time `json:"submission_date" binding:"required"`
	CourseID       int64     `json:"course_id" binding:"required"`
}

func (server *Server) createAssignment(ctx *gin.Context) {
	var req createAssignmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAssignmentParams{
		Type:           req.Type,
		Title:          req.Title,
		Description:    req.Description,
		SubmissionDate: req.SubmissionDate,
	}

	assignment, err := server.store.CreateAssignment(ctx, db.CreateAssignmentParam(arg))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, assignment)
}

// @Summary Get an assignment by ID
// @Description Get an assignment by its ID and resource ID
// @ID get-assignment
// @Accept json
// @Produce json
// @Param assignment_id path int true "Assignment ID"
// @Param resource_id path int true "Resource ID"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /assignments/{assignment_id}/{resource_id} [get]
type GetAssignmentRequest struct {
	AssignmentID int64         `json:"assignment_id"`
	ResourceID   int64 `json:"resource_id"`
}

func (server *Server) getAssignment(ctx *gin.Context) {
	var req GetAssignmentRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	assignment, err := server.store.GetAssignment(ctx, db.GetAssignmentParam{
		AssignmentID: req.AssignmentID,
		ResourceID: req.ResourceID,
	})
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, assignment)
}

// @Summary Update an assignment
// @Description Update an assignment with the provided details
// @ID update-assignment
// @Accept json
// @Produce json
// @Param request body UpdateAssignmentRequest true "Updated assignment details"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /assignments/{assignment_id} [put]
type UpdateAssignmentRequest struct {
	Type           string    `json:"type"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	SubmissionDate time.Time `json:"submission_date"`
}

func (server *Server) updateAssignment(ctx *gin.Context) {
	var req UpdateAssignmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateAssignmentParams{
		Type:           req.Type,
		Title:          req.Title,
		Description:    req.Description,
		SubmissionDate: req.SubmissionDate,
	}

	assignment, err := server.store.UpdateAssignment(ctx, db.UpdateAssignmentParam{
		Type: arg.Type,
		Title: arg.Title,
		Description: arg.Description,
		SubmissionDate: arg.SubmissionDate,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, assignment)
}

type DeleteAssignmentRequest struct {
	AssignmentID int64         `json:"assignment_id"`
	ResourceID   int64 `json:"resource_id"`
}

// @Summary Delete an assignment
// @Description Delete an assignment by its ID and resource ID
// @ID delete-assignment
// @Accept json
// @Produce json
// @Param assignment_id path int true "Assignment ID"
// @Param resource_id path int true "Resource ID"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /assignments/{assignment_id}/{resource_id} [delete]
func (server *Server) deleteAssignment(ctx *gin.Context) {
	assignmentID, err := strconv.ParseInt(ctx.Param("assignment_id"), 10, 64)
	resourceID, err := strconv.ParseInt(ctx.Param("resource_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("invalid assignment_id")))
		return
	}

	err = server.store.DeleteAssignment(ctx, db.DeleteAssignmentParam{
		AssignmentID: assignmentID,
		ResourceID : resourceID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Assignment deleted successfully"})
}
