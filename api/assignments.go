// Assignemts cotroller created
package api

import (
	db "eduwave-back-end/db/sqlc"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type createAssignmentRequest struct {
	ResourceID     int64     `json:"resource_id"`
	Type           string    `json:"type"`
    Title          string    `json:"title"`
    Description    string    `json:"description"`
    SubmissionDate time.Time `json:"submission_date"`
}

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
func (server *Server) createAssignment(ctx *gin.Context) {
	var req createAssignmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAssignmentParams{
		ResourceID: req.ResourceID,
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
	AssignmentID int64         `form:"assignment_id"`
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
// @Router /assignment/get [get]
func (server *Server) getAssignment(ctx *gin.Context) {
	var req GetAssignmentRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.GetAssignmentParam{AssignmentID: req.AssignmentID}

	assignment, err := server.store.GetAssignment(ctx, arg)
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
	AssignmentID   int64     `json:"assignment_id"`
	Type           string    `json:"type"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	SubmissionDate time.Time `json:"submission_date"`
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
// @Router /assignments/edit [put]
func (server *Server) updateAssignment(ctx *gin.Context) {
	var req UpdateAssignmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateAssignmentParams{
		AssignmentID: req.AssignmentID,
		Type:           req.Type,
		Title:          req.Title,
		Description:    req.Description,
		SubmissionDate: req.SubmissionDate,
	}

	assignment, err := server.store.UpdateAssignment(ctx, db.UpdateAssignmentParam(arg))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, assignment)
}

type DeleteAssignmentRequest struct {
	AssignmentID int64         `form:"assignment_id"`
}

// @Summary Delete an assignment
// @Description Delete an assignment by its ID and resource ID
// @ID delete-assignment
// @Accept json
// @Produce json
// @Param assignment_id path int true "Assignment ID"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /assignment/delete [delete]
func (server *Server) deleteAssignment(ctx *gin.Context) {
	var req DeleteAssignmentRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	err := server.store.DeleteAssignment(ctx, db.DeleteAssignmentParam{
		AssignmentID: req.AssignmentID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Assignment deleted successfully"})
}

type GetassignmentByResourceRequest struct {
	ResourceID  int64 `form:"resource_id"`
}

// @Summary Get Assignment By Resource
// @Description Get Assignment By Resource
// @ID update-assignment
// @Accept json
// @Produce json
// @Param request body GetassignmentByResourceRequest true "Get Assignment By Resource details"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /assignment/byresource [put]
func(server *Server) GetassignmentByResource(ctx *gin.Context){
	var req GetassignmentByResourceRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
	}

	arg := db.GetAssignmentByResourceParam{
		ResourceID: req.ResourceID,
	}

	assignmentbyresource, err := server.store.GetAssignmentByResource(ctx, arg)
	if err != nil {
		if errors.Is(err ,db.ErrRecordNotFound){
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, assignmentbyresource)
}
