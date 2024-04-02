// requests controller
package api

import (
	"database/sql"
	db "eduwave-back-end/db/sqlc"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createRequestRequest struct {
	IsActive   sql.NullBool `json:"is_active"`
	IsPending  sql.NullBool `json:"is_pending"`
	IsAccepted sql.NullBool `json:"is_accepted"`
	IsDeclined sql.NullBool `json:"is_declined"`
}

// createRequest represents the request body for create a request.
// @Summary Create a new request
// @Description Create a new request with the given parameters
// @Accept json
// @Produce json
// @Param request body createRequestRequest true "Request data"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /requests [post]
func (server *Server) createRequest(ctx *gin.Context) {
	var req createRequestRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateRequestParams{
		IsActive:   req.IsActive,
		IsPending:  req.IsPending,
		IsAccepted: req.IsAccepted,
		IsDeclined: req.IsDeclined,
	}

	request, err := server.store.CreateRequest(ctx, db.CreateRequestParam(arg))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, request)
}

type getRequestRequest struct {
	RequestID int64 `uri:"Request_id"`
}

// @Summary Get a request by ID
// @Description Get a request by its ID
// @Produce json
// @Param request_id path int true "Request ID"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /request/{request_id} [get]
func (server *Server) getRequest(ctx *gin.Context) {
	var req getRequestRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetRequestParam{RequestID: req.RequestID}

	request, err := server.store.GetRequest(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, request)
}

type listRequestRequest struct {
	StudentID int64 `form:"student_id"`
	TeacherID int64 `form:"teacher_id"`
	CourseID  int64 `form:"course_id"`
	Limit     int32 `form:"limit" binding:"required,min=5,max=10"`
	Offset    int32 `form:"offset" binding:"required,min=0"`
}

// @Summary List requests
// @Description List requests based on provided parameters
// @Produce json
// @Param student_id query int false "Student ID"
// @Param teacher_id query int false "Teacher ID"
// @Param course_id query int false "Course ID"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /requests [get]
func (server *Server) listRequests(ctx *gin.Context) {
	var req listRequestRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListRequestParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	requests, err := server.store.ListRequest(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, requests)
}

type deleteRequestRequest struct {
	StudentID int64 `json:"student_id"`
	RequestID int64         `json:"request_id"`
}

// @Summary Delete a request
// @Description Delete a request by student and request ID
// @Produce json
// @Param student_id path int true "Student ID"
// @Param request_id path int true "Request ID"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /request/{student_id}/{request_id} [delete]
func (server *Server) deleteRequest(ctx *gin.Context) {
	var req deleteRequestRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteRequest(ctx, db.DeleteRequestParam{
		StudentID: req.StudentID,
		RequestID: req.RequestID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Request deleted successfully"})
}

type updateRequestRequest struct {
	ID         int64 `json:"id" binding:"required,min=1"`
	IsActive   bool  `json:"is_active" binding:"required"`
	IsPending  bool  `json:"is_pending" binding:"required"`
	IsAccepted bool  `json:"is_accepted" binding:"required"`
	IsDeclined bool  `json:"is_declined" binding:"required"`
}

// @Summary Update a request
// @Description Update a request with the provided parameters
// @Accept json
// @Produce json
// @Param request body updateRequestRequest true "Request data"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /request/{student_id} [put]
func (server *Server) updateRequest(ctx *gin.Context) {
	var req updateRequestRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateRequestsParams{
		StudentID:  req.ID,
		IsActive:   sql.NullBool{Bool: req.IsActive, Valid: true},
		IsPending:  sql.NullBool{Bool: req.IsPending, Valid: true},
		IsAccepted: sql.NullBool{Bool: req.IsAccepted, Valid: true},
		IsDeclined: sql.NullBool{Bool: req.IsDeclined, Valid: true},
	}

	request, err := server.store.UpdateRequests(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, request)
}
