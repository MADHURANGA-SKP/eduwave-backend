// requests controller
package api

import (
	"database/sql"
	db "eduwave-back-end/db/sqlc"
	"eduwave-back-end/token"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createRequestRequest struct {
	UserID     int64 `json:"user_id"`
	CourseID   int64 `json:"course_id"`
	IsActive   bool  `json:"is_active"`
	IsPending  bool  `json:"is_pending"`
	IsAccepted bool  `json:"is_accepted"`
	IsDeclined bool  `json:"is_declined"`
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

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if req.UserID != authPayload.UserID {
		err := errors.New("permission denied. account doesn't belongs to the authenticated user")
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	arg := db.CreateRequestParams{
		UserID:     req.UserID,
		CourseID:   req.CourseID,
		IsActive:   sql.NullBool{Bool: false, Valid: false},
		IsPending:  sql.NullBool{Bool: true, Valid: true},
		IsAccepted: sql.NullBool{Bool: false, Valid: false},
		IsDeclined: sql.NullBool{Bool: false, Valid: false},
	}

	request, err := server.store.CreateRequest(ctx, db.CreateRequestParam{
		UserID:     arg.UserID,
		CourseID:   arg.CourseID,
		IsActive:   arg.IsActive,
		IsPending:  arg.IsPending,
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
	UserID   int64 `form:"user_id"`
    CourseID int64 `form:"course_id"`
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
func (server *Server) GetRequest(ctx *gin.Context) {
	var req getRequestRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetRequestParam{
		UserID: req.UserID,
		CourseID: req.CourseID,
	}

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
	PageSize int32 `form:"page_size" binding:"required,min=10,max=100"`
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
func (server *Server) ListRequest(ctx *gin.Context) {
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
	UserID     int64 `json:"user_id"`
    CourseID   int64 `json:"course_id"`
	IsActive   bool `json:"is_active"`
	IsPending  bool `json:"is_pending"`
	IsAccepted bool `json:"is_accepted"`
	IsDeclined bool `json:"is_declined"`
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
func (server *Server) UpdateRequests(ctx *gin.Context) {
	var req updateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// userID, err := strconv.Atoi(ctx.Param("user_id"))
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	// }

	arg := db.UpdateRequestsParam{
		UserID:     req.UserID,
		CourseID: req.CourseID,
		IsActive:   req.IsActive,
		IsPending:  req.IsPending,
		IsAccepted: req.IsAccepted,
		IsDeclined: req.IsDeclined,
	}

	request, err := server.store.UpdateRequests(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, request)
}

type userRequestResponse struct {
	Username      string    `json:"user_name"`
	FullName      string    `json:"full_name"`
}

func newRequestResponse(user db.User) userRequestResponse {
	return userRequestResponse{
		Username: user.UserName,
		FullName: user.FullName,
	}
}

type ListRequestByUserRequest struct {
	UserID int64 `form:"user_id"`
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=10,max=100"`
}

type courseDetails []db.Course

// @Summary List requests By user
// @Description List requests based on User
// @Produce json
// @Param user_if query int "User ID"
// @Param limit query int  "Limit"
// @Param offset query int  "Offset"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /requests/byuser [get]
func (server *Server) ListRequestByUser(ctx *gin.Context) {
	var req ListRequestByUserRequest
	var courseDetails courseDetails 
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUserById(ctx, db.GetUserByIdParam{
		UserID: req.UserID,
	})
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if user.User.UserID != authPayload.UserID {
		err := errors.New("account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	arg := db.ListRequestByUserParams{
		UserID: authPayload.UserID,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	request, err := server.store.ListRequestByUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	for _, request := range request {
		course, err := server.store.GetCourse(ctx, db.GetCourseParam{
		  CourseID: request.CourseID,
		})
		if request.CourseID != course.Course.CourseID {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		courseDetails = append(courseDetails, course.Course)
	}

	userRsp := newRequestResponse(user.User)

	response := gin.H{
		"user"	: userRsp,
		"number_of_request": request,
		"requested_Course_details" : courseDetails,
	}

	ctx.JSON(http.StatusOK, response)
}



type ListRequestByCourseRequest struct {
	CourseID int64 `form:"course_id"`
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=10,max=100"`
}

// @Summary List requests By course
// @Description List requests based on course
// @Produce json
// @Param course_id query int "course ID"
// @Param limit query int  "Limit"
// @Param offset query int  "Offset"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /requests/bycourse [get]
func (server *Server) ListRequestByCourse(ctx *gin.Context) {
	var req ListRequestByCourseRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListRequestByCourseParams{
		CourseID: req.CourseID,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	requests, err := server.store.ListRequestByCourse(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, requests)
}
