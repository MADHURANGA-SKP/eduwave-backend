// courses controller

package api

import (
	"net/http"

	db "eduwave-back-end/db/sqlc"

	"github.com/gin-gonic/gin"
)

// CreateCourseRequest defines the request body structure for creating a course
type CreateCourseRequest struct {
	Title       string `json:"title" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// CreateCourse creates a new course
func (server *Server) CreateCourse(ctx *gin.Context) {
	var req CreateCourseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.CreateCoursesParams{
		Title:       req.Title,
		Type:        req.Type,
		Description: req.Description,
	}

	course, err := server.store.CreateCourses(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, course)
}

// GetCourseRequest defines the request body structure for getting a course
type GetCourseRequest struct {
	CourseID int64 `uri:"course_id" binding:"required,min=1"`
}

// GetCourse retrieves a course by ID
func (server *Server) GetCourse(ctx *gin.Context) {
	var req GetCourseRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course, err := server.store.GetCourses(ctx, db.GetCoursesParams{
		CourseID: req.CourseID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, course)
}

// ListCoursesRequest defines the request body structure for listing courses
type ListCoursesRequest struct {
	Limit  int32 `form:"limit" binding:"required,min=1,max=100"`
	Offset int32 `form:"offset" binding:"required,min=0"`
}

// ListCourses lists courses
func (server *Server) ListCourses(ctx *gin.Context) {
	var req ListCoursesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.ListCoursesParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	courses, err := server.store.ListCourses(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, courses)
}
