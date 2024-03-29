package api

import (
	"database/sql"
	db "eduwave-back-end/db/sqlc"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createRequestRequest struct {
	IsActive   bool `json:"is_active" binding:"required"`
	IsPending  bool `json:"is_pending" binding:"required"`
	IsAccepted bool `json:"is_accepted" binding:"required"`
	IsDeclined bool `json:"is_declined" binding:"required"`
}

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

	request, err := server.store.CreateRequest(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, request)
}

type getRequestRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getRequest(ctx *gin.Context) {
	var req getRequestRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	request, err := server.store.GetRequest(ctx, req.ID)
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

func (server *Server) listRequests(ctx *gin.Context) {
	var req listRequestRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListRequestParams{
		StudentID: req.StudentID,
		TeacherID: req.TeacherID,
		CourseID:  req.CourseID,
		Limit:     req.Limit,
		Offset:    req.Offset,
	}

	requests, err := server.store.ListRequest(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, requests)
}

type deleteRequestRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteRequest(ctx *gin.Context) {
	var req deleteRequestRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteRequest(ctx, req.ID)
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

func (server *Server) updateRequest(ctx *gin.Context) {
	var req updateRequestRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateRequestsParams{
		StudentID:  sql.NullInt64{Int64: req.ID, Valid: true},
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
