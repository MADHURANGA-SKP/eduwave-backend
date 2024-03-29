// course enrolments controller
package api

import (
	"database/sql"
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

func (server *Server) listEnrolments(ctx *gin.Context) {
	var req listEnrolmentsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	enrolments, err := server.store.ListEnrolments(ctx, db.ListEnrolmentsParams{
		StudentID: sql.NullInt64{Int64: req.StudentID, Valid: true},
		CourseID:  sql.NullInt64{Int64: req.CourseID, Valid: true},
		Limit:     req.Limit,
		Offset:    req.Offset,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, enrolments)
}
