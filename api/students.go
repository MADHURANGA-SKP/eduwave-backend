// students controller
package api

import (
	"net/http"
	"strconv"

	db "eduwave-back-end/db/sqlc"

	"github.com/gin-gonic/gin"
)

// createUserRequest defines the request body structure for creating a new student
type createStudentRequest struct {
	UserName string `json:"user_name" binding:"required"`
}

// studentResponse defines the response body structure for a student
type studentResponse struct {
	StudentID int64  `json:"student_id"`
	UserName  string `json:"user_name"`
}

// createStudent creates a new student
func (server *Server) createStudent(ctx *gin.Context) {
	var req createStudentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	student, err := server.store.CreateStudent(ctx, db.CreateStudentParams{
		UserName: req.UserName,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := studentResponse{
		StudentID: student.StudentID,
		UserName:  student.UserName,
	}
	ctx.JSON(http.StatusOK, rsp)
}

// listStudents returns a list of students
func (server *Server) listStudents(ctx *gin.Context) {
	userName := ctx.Query("user_name")
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

	students, err := server.store.ListStudents(ctx, db.ListStudentsParams{
		UserName: userName,
		Limit:    int32(limit),
		Offset:   int32(offset),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var resp []studentResponse
	for _, student := range students {
		resp = append(resp, studentResponse{
			StudentID: student.StudentID,
			UserName:  student.UserName,
		})
	}

	ctx.JSON(http.StatusOK, resp)
}

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

// updateStudent updates a student by ID
func (server *Server) updateStudent(ctx *gin.Context) {
	// Parse student ID from path parameter
	studentID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Parse request body to get updated student information
	var req createStudentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Call the database store function to update the student
	updatedStudent, err := server.store.UpdateStudent(ctx, db.UpdateStudentParams{
		StudentID: studentID,
		UserName:  req.UserName,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Respond with the updated student information
	rsp := studentResponse{
		StudentID: updatedStudent.StudentID,
		UserName:  updatedStudent.UserName,
	}
	ctx.JSON(http.StatusOK, rsp)
}
