// course enrolments controller
package api

import (
	"net/http"

	db "eduwave-back-end/db/sqlc"

	"github.com/gin-gonic/gin"
)

type listEnrolmentsRequest struct {
	StudentID int64 `form:"student_id" binding:"required"`
	CourseID  int64 `form:"course_id" binding:"required"`
	Limit     int32 `form:"limit" binding:"required"`
	Offset    int32 `form:"offset" binding:"required"`
}

// @Summary List enrolments
// @Description List enrolments for a student in a course
// @ID list-enrolments
// @Accept  json
// @Produce  json
// @Param student_id query int true "Student ID"
// @Param course_id query int true "Course ID"
// @Param limit query int true "Limit"
// @Param offset query int true "Offset"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /courseEnrolments [get]
func (server *Server) listEnrolments(ctx *gin.Context) {
	var req listEnrolmentsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	enrolments, err := server.store.ListEnrolments(ctx, db.ListEnrolmentsParams{
		StudentID: req.StudentID,
		CourseID:  req.CourseID,
		Limit:     req.Limit,
		Offset:    req.Offset,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, enrolments)
}
