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
	UserID     int64        `json:"user_id"`
    CourseID   int64        `json:"course_id"`
	IsActive   bool `json:"is_active"`
	IsPending  bool `json:"is_pending"`
	IsAccepted bool `json:"is_accepted"`
	IsDeclined bool `json:"is_declined"`
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
		UserID: req.UserID,
		CourseID: req.CourseID,
		IsActive:   sql.NullBool{Bool: req.IsActive, Valid: req.IsActive},
		IsPending:  sql.NullBool{Bool: req.IsPending, Valid: req.IsPending},
		IsAccepted: sql.NullBool{Bool: req.IsAccepted, Valid: req.IsAccepted},
		IsDeclined: sql.NullBool{Bool: req.IsDeclined, Valid: req.IsDeclined},
	}

	request, err := server.store.CreateRequest(ctx, db.CreateRequestParam{
		UserID: arg.UserID,
		CourseID: arg.CourseID,
		IsActive: arg.IsActive,
		IsPending: arg.IsPending,
		IsAccepted: arg.IsAccepted,
		IsDeclined: arg.IsDeclined,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, request)
}

type getRequestRequest struct {
	RequestID int64 `form:"request_id"`
}

// @Summary Get a request by ID
// @Description Get a request by its ID
// @Produce json
// @Param request_id path int true "Request ID"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /request/get [get]
func (server *Server) getRequest(ctx *gin.Context) {
	var req getRequestRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
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
	PageID   int32 `form:"page_id" binding:"required,min=1"`
    PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
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
		Limit:  req.PageSize,
        Offset: (req.PageID - 1) * req.PageSize,
	}

	requests, err := server.store.ListRequest(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, requests)
}

type deleteRequestRequest struct {
    RequestID int64 `form:"request_id"`
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
// @Router /request/delete [delete]
func (server *Server) deleteRequest(ctx *gin.Context) {
	var req deleteRequestRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteRequest(ctx, req.RequestID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Request deleted successfully"})
}

type updateRequest struct {
	UserID     int64        `json:"user_id"`
    IsActive   sql.NullBool `json:"is_active"`
    IsPending  sql.NullBool `json:"is_pending"`
    IsAccepted sql.NullBool `json:"is_accepted"`
    IsDeclined sql.NullBool `json:"is_declined"`
}

// @Summary Update a request
// @Description Update a request with the provided parameters
// @Accept json
// @Produce json
// @Param request body updateRequest true "Requested data"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /request/edit [put]
func (server *Server) updateRequest(ctx *gin.Context) {
	var req updateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// userID, err := strconv.Atoi(ctx.Param("user_id"))
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	// }


	arg := db.UpdateRequestsParams{
		UserID: req.UserID,
		IsActive:   sql.NullBool{Bool: req.IsActive.Bool, Valid: true},
		IsPending:  sql.NullBool{Bool: req.IsPending.Bool, Valid: true},
		IsAccepted: sql.NullBool{Bool: req.IsAccepted.Bool, Valid: true},
		IsDeclined: sql.NullBool{Bool: req.IsDeclined.Bool, Valid: true},
	}

	request, err := server.store.UpdateRequests(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, request)
}
