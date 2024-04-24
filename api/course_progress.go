// course progreass controller
package api

import (
	"database/sql"
	"net/http"

	db "eduwave-back-end/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createCourseProgressRequest struct {
	EnrolmentID int64  `json:"enrolment_id"`
    Progress    string `json:"progress"`
}

// @Summary Create course progress
// @Description Creates course progress
// @Tags CourseProgress
// @Accept json
// @Produce json
// @Param progress body createCourseProgressRequest true "Course progress object"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /createprogress [post]
// createCourseProgressRequest defines the request body structure for creating course progress
func (server *Server) createCourseProgress(ctx *gin.Context) {
	var req createCourseProgressRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCourseProgressPram{
		EnrolmentID: req.EnrolmentID,
		Progress: req.Progress,
	}

	courseProgress, err := server.store.CreateCourseProgress(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, courseProgress)
}

type getCourseProgressRequest struct {
	CourseprogressID int64         `form:"courseprogress_id"`
    EnrolmentID      int64 `form:"enrolment_id"`
}

// @Summary Get course progress by ID
// @Description Returns course progress by ID
// @Tags CourseProgress
// @Accept json
// @Produce json
// @Param courseprogress_id path int true "Course progress ID"
// @Param enrolment_id query int true "Enrolment ID"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /courseProgress/get [get]
// getCourseProgress returns course progress by ID
func (server *Server) getCourseProgress(ctx *gin.Context) {
	var req getCourseProgressRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}


	courseProgress, err := server.store.GetCourseProgress(ctx,db.GetCourseProgressParam{
		CourseprogressID: req.CourseprogressID,
		EnrolmentID: req.EnrolmentID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, courseProgress)
}

type listCourseProgressRequest struct {
	EnrolmentID int64 `form:"enrolment_id"`
    PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=10,max=100"`
}

// @Summary List course progress
// @Description Returns a list of course progress
// @Tags CourseProgress
// @Accept json
// @Produce json
// @Param enrolment_id query int true "Enrolment ID"
// @Param limit query int false "Limit the number of results"
// @Param offset query int false "Offset for paginated results"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /courseprogress/byenrolment [get]
// listCourseProgress returns a list of course progress
func (server *Server) ListCourseProgress(ctx *gin.Context) {
	var req listCourseProgressRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListCourseProgressParams{
		EnrolmentID: req.EnrolmentID,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	courseProgressList, err := server.store.ListCourseProgress(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, courseProgressList)
}

// UpdateCourseProgressRequest defines the request body structure for listing courses
type UpdateCourseProgressRequest struct {
	EnrolmentID int64  `json:"enrolment_id"`
    Progress    string `json:"progress"`
}

// @Summary Update a UpdateCourseProgress
// @Description Updates a UpdateCourseProgress with provided details
// @Accept json
// @Produce json
// @Param request body UpdateCourseProgressRequest true "Updated UpdateCourseProgress details"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /courseprogress/edit [put]
// UpdateCourse updates the selected course
func (server *Server) UpdateCourseProgress(ctx *gin.Context) {
	var req UpdateCourseProgressRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateCourseProgressParams{
		EnrolmentID: req.EnrolmentID,
		Progress: req.Progress,
	}

	courses, err := server.store.UpdateCourseProgress(ctx, db.UpdateCourseProgressParam(arg))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, courses)
}