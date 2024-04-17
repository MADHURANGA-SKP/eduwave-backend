// courses controller

package api

import (
	"net/http"

	db "eduwave-back-end/db/sqlc"

	"github.com/gin-gonic/gin"
)

// CreateCourseRequest defines the request body structure for creating a course
type CreateCourseRequest struct {
	UserID      int64  `json:"user_id"`
	Title       string `json:"title" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Description string `json:"description" binding:"required"`
	Image       []byte `json:"image"`
}

// @Summary Create a new course
// @Description Creates a new course
// @Accept json
// @Produce json
// @Param request body CreateCourseRequest true "Course details"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /course [post]
// CreateCourse creates a new course
func (server *Server) CreateCourse(ctx *gin.Context) {
	var req CreateCourseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCoursesParams{
		UserID: req.UserID,
		Title:       req.Title,
		Type:        req.Type,
		Description: req.Description,
		Image: req.Image,
	}

	course, err := server.store.CreateCourses(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, course)
}

// GetCourseRequest defines the request body structure for getting a course
type GetCourseRequest struct {
	CourseID  int64    `form:"course_id,min=1"`
}

// @Summary Get a course by ID
// @Description Retrieves a course by its ID
// @Produce json
// @Param course_id path int true "Course ID"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /course/get [get]
// GetCourse retrieves a course by ID
func (server *Server) GetCourse(ctx *gin.Context) {
	var req GetCourseRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetCourseParam{CourseID: req.CourseID}

	course, err := server.store.GetCourse(ctx, db.GetCourseParam(arg))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, course)
}

type ListCoursesRequest struct {
	PageID   int32 `form:"page_id,min=1"`
	PageSize int32 `form:"page_size,min=10,max=10"`
}

// @Summary List courses
// @Description Lists courses with pagination
// @Produce json
// @Param limit query int true "Number of items to return"
// @Param offset query int true "Offset for pagination"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /courses [get]
// ListCoursesRequest defines the request body structure for listing courses
// ListCourses lists courses
func (server *Server) ListCourses(ctx *gin.Context) {
	var req ListCoursesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListCoursesParams{
		Limit:  req.PageSize,
        Offset: (req.PageID - 1) * req.PageSize,
	}

	courses, err := server.store.ListCourses(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, courses)
}

// UpdateCoursesRequest defines the request body structure for listing courses
type UpdateCoursesRequest struct {
	CourseID    int64  `json:"course_id"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Image       []byte `json:"image"`
}

// @Summary Update a course
// @Description Updates a course with provided details
// @Accept json
// @Produce json
// @Param request body UpdateCoursesRequest true "Updated course details"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /course/edit [put]
// UpdateCourse updates the selected course
func (server *Server) UpdateCourses(ctx *gin.Context) {
	var req UpdateCoursesRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateCoursesParams{
		CourseID: req.CourseID,
		Title: req.Title,
		Type: req.Type,
		Description: req.Description,
		Image: req.Image,
	}

	courses, err := server.store.UpdateCourses(ctx, db.UpdateCoursesParam(arg))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, courses)
}

// deleteCourseRequest defines the request body structure for deleting an Course
type deleteCourseRequest struct {
	CourseID  int64 `form:"course_id"`
}

// @Summary Delete a course
// @Description Deletes a course by ID
// @Produce json
// @Param course_id path int true "Course ID"
// @Param teacher_id query int true "Teacher ID"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /course/delete [delete]
// deleteCourse deletes an Course
func (server *Server) DeleteCourse(ctx *gin.Context) {
	var req deleteCourseRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteCourse(ctx, db.DeleteCourseParam{
		CourseID: req.CourseID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}
