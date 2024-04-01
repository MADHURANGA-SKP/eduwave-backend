// Assignemts cotroller created
package api

import (
	"database/sql"
	db "eduwave-back-end/db/sqlc"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

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

type GetAssignmentRequest struct {
	AssignmentID int64         `json:"assignment_id"`
	ResourceID   sql.NullInt64 `json:"resource_id"`
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

	courseID, err := strconv.ParseInt(ctx.Param("course_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("invalid course_id")))
		return
	}

	arg := db.UpdateAssignmentParams{
		ResourceID:       sql.NullInt64{Int64: courseID, Valid: true},
		Type:           req.Type,
		Title:          req.Title,
		Description:    req.Description,
		SubmissionDate: req.SubmissionDate,
	}

	assignment, err := server.store.UpdateAssignment(ctx, db.UpdateAssignmentParam{
		ResourceID: arg.ResourceID,
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
	ResourceID   sql.NullInt64 `json:"resource_id"`
}

func (server *Server) deleteAssignment(ctx *gin.Context) {
	assignmentID, err := strconv.ParseInt(ctx.Param("assignment_id"), 10, 64)
	resourceID, err := strconv.ParseInt(ctx.Param("resource_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("invalid assignment_id")))
		return
	}

	err = server.store.DeleteAssignment(ctx, db.DeleteAssignmentParam{
		AssignmentID: assignmentID,
		ResourceID : sql.NullInt64{Int64: resourceID, Valid: true},
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Assignment deleted successfully"})
}
