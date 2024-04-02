// students controller
package api

import (
	"database/sql"
	"net/http"
	"strconv"

	db "eduwave-back-end/db/sqlc"

	"github.com/gin-gonic/gin"
)

// // createUserRequest defines the request body structure for creating a new student
// type createStudentRequest struct {
// 	UserName string `json:"user_name" binding:"required"`
// }

// // createStudent creates a new student
// func (server *Server) createStudent(ctx *gin.Context) {
// 	var req createStudentRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	arg := db.CreateStudentParam{
// 		UserName: req.UserName,
// 	}

// 	student, err := server.store.CreateStudent(ctx, arg)
// 	if err != nil {
// 		errCode := db.ErrorCode(err)
// 		if errCode == db.ForeignKeyViolation || errCode == db.UniqueViolations {
// 			ctx.JSON(http.StatusForbidden, errorResponse(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, student)
// }

type ListStudentRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// @Summary List students
// @Description Returns a list of students with pagination
// @ID list-students
// @Produce json
// @Param page_id query int true "Page ID" minimum(1)
// @Param page_size query int true "Page Size" minimum(5) maximum(10)
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /students [get]
// listStudents returns a list of students
func (server *Server) listStudents(ctx *gin.Context) {
	var req ListStudentRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListStudentsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	rsp, err := server.store.ListStudents(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, rsp)
}

// @Summary Delete student by ID
// @Description Deletes a student by their ID
// @ID delete-student
// @Produce json
// @Param id path int64 true "Student ID"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /students/{student_id} [delete]
// deleteStudent deletes a student by ID
func (server *Server) deleteStudent(ctx *gin.Context) {
	// Parse student ID from path parameter
	studentID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Call the database store function to delete the student
	err = server.store.DeleteStudent(ctx, studentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Respond with success
	ctx.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}

type UpdateStudentRequest struct {
	StudentID int64  `json:"student_id"`
    UserName  string `json:"user_name"`
}

// @Summary Update student
// @Description Update a student by their ID
// @ID update-student
// @Accept json
// @Produce json
// @Param id path int64 true "Student ID"
// @Param student body UpdateStudentRequest true "Student data"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /students/:student_id [put]
// updateStudent updates a student by ID
func (server *Server) updateStudent(ctx *gin.Context) {
	var req UpdateStudentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateStudentParams{
		StudentID: req.StudentID,
		UserName:  req.UserName,
	}

	// Call the database store function to update the student
	updatedStudent, err := server.store.UpdateStudent(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, updatedStudent)
}
