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
// @Router /courseProgress [post]
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
	CourseprogressID int64         `json:"courseprogress_id"`
    EnrolmentID      int64 `json:"enrolment_id"`
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
// @Router /courseProgress/{courseprogress_id} [get]
// getCourseProgress returns course progress by ID
func (server *Server) getCourseProgress(ctx *gin.Context) {
	var req getCourseProgressRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
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
	EnrolmentID int64 `json:"enrolment_id"`
    Limit       int32 `json:"limit"`
    Offset      int32 `json:"offset"`
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
// @Router /courseProgress [get]
// listCourseProgress returns a list of course progress
func (server *Server) listCourseProgress(ctx *gin.Context) {
	var req listCourseProgressRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListCourseProgressParams{
		EnrolmentID: req.EnrolmentID,
		Limit:       req.Limit,
		Offset:      req.Offset,
	}

	courseProgressList, err := server.store.ListCourseProgress(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, courseProgressList)
}
