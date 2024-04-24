// courses controller

package api

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	db "eduwave-back-end/db/sqlc"
	"eduwave-back-end/token"

	"github.com/gin-gonic/gin"
)

// uploadSingleFile handles uploading a single file
func(server *Server) uploadSingleImage(ctx *gin.Context, file multipart.File, header *multipart.FileHeader) (string, error) {
	// Validate file extension (optional)
	allowedExtensions := map[string]bool{
	  ".jpg":  true,
	  ".jpeg": true,
	  ".png":  true,
	}
	fileExt := filepath.Ext(header.Filename)
	if !allowedExtensions[fileExt] {
	  return "", fmt.Errorf("unsupported file extension: %s", fileExt)
	}
  
	// Generate unique filename
	originalFileName := strings.TrimSuffix(filepath.Base(header.Filename), filepath.Ext(header.Filename))
	now := time.Now()
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
  
	// Create upload directory if it doesn't exist
	uploadDir := "uploads/images"
	if err := os.MkdirAll(uploadDir, 0755); err != nil{
	  return "", err
	}
  
	// Create destination file path
	filePath := filepath.Join(uploadDir, filename)
  
	// Save uploaded file
	out, err := os.Create(filePath)
	if err != nil {
	  return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
	  return "", err
	}
  
	return filename, nil
  }
  
// CreateCourseRequest defines the request body structure for creating a course
type CreateCourseRequest struct {
	UserID      int64  `json:"user_id"`
	Title       string `form:"title" binding:"required"`
	Type        string `form:"type" binding:"required"`
	Description string `form:"description" binding:"required"`
	Image       string `json:"image"`
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
// @Router /course/:user_id [post]
// CreateCourse creates a new course
func (server *Server) CreateCourse(ctx *gin.Context) {
	var req CreateCourseRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID, err := strconv.Atoi(ctx.Param("user_id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, errorResponse(err))
        return
    }

	// Handle image upload (if included in the request)
	var imageFile string
	file, header, err := ctx.Request.FormFile("image") 
	if err == nil {
		imageFile, err = server.uploadSingleImage(ctx, file, header)
	  if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	  }
	}
	// Update request struct with image filename (if uploaded)
	req.Image = imageFile

	arg := db.CreateCoursesParams{
		UserID: int64(userID),
		Title:       req.Title,
		Type:        req.Type,
		Description: req.Description,
		Image:       req.Image,
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
	CourseID  int64    `form:"course_id"`
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
	Image       string `json:"image"`
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

	// // Access uploaded file information
	// file, err := ctx.FormFile("image") // Replace "file" with your actual form field name
	// if err != nil {
	//   ctx.JSON(http.StatusBadRequest, errorResponse(err))
	//   return
	// }
  
	// // Open the uploaded file on the server
	// openedFile, err := file.Open()
	// if err != nil {
	//   ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	//   return
	// }
	// defer openedFile.Close() // Close the file after processing
  
	// // Read the file contents into a byte array
	// fileData, err := io.ReadAll(openedFile)
	// if err != nil {
	//   ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	//   return
	// }

	arg := db.UpdateCoursesParams{
		CourseID:    req.CourseID,
		Title:       req.Title,
		Type:        req.Type,
		Description: req.Description,
		Image:       req.Image,
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

// getTotalStudentsByCourse retrieves the total number of students enrolled in each course.
func (server *Server) getTotalStudentsByCourse(ctx *gin.Context) []gin.H {
	var result []gin.H

	courses, err := server.store.ListCourses(ctx, db.ListCoursesParams{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return result
	}

	for _, course := range courses {
		enrolments, err := server.store.ListEnrolments(ctx, db.ListEnrolmentsParams{
			CourseID: course.CourseID,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			continue
		}

		courseInfo := gin.H{
			"courseId":     course,
			"StudentCount": len(enrolments),
		}
		result = append(result, courseInfo)
	}

	return result
}

func (server *Server) getTotalCoursesByUserID(ctx *gin.Context) []gin.H {
	var result []gin.H

	users, err := server.store.ListUsers(ctx, db.ListUserParams{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return result
	}

	for _, user := range users {
		enrolments, err := server.store.ListEnrolmentsByUser(ctx, db.ListEnrolmentsByUserParams{
			UserID: user.UserID,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			continue
		}

		userInfo := gin.H{
			"userId":      user.UserID,
			"CourseCount": len(enrolments),
		}
		result = append(result, userInfo)
	}

	return result
}


type ListCoursesByUserRequest struct {
	UserID int64 `form:"user_id"`
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// @Summary List courses
// @Description Lists courses with pagination by usercreation
// @Produce json
// @Param limit query int true "Number of items to return"
// @Param offset query int true "Offset for pagination"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /course/byuser [get]
// ListCoursesRequest defines the request body structure for listing courses
// ListCourses lists courses
func (server *Server) ListCoursesByUser(ctx *gin.Context) {
	var req ListCoursesByUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if req.UserID != authPayload.UserID {
		err := errors.New("account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := db.ListCoursesByUserParams{
		UserID: authPayload.UserID,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	courses, err := server.store.ListCoursesByUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, courses)
}


// GetCourseByStudentRequest defines the request body structure for getting a course
type GetCourseWithRequestdetails struct {
	CourseID  int64    `form:"course_id"`
}

type Request []db.Request

// @Summary Get a course by ID
// @Description Retrieves a course by its ID
// @Produce json
// @Param course_id path int true "Course ID"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /course/get [get]
// GetCourseWithRequestdetails retrieves a course by ID
func (server *Server) GetCourseWithRequestdetails(ctx *gin.Context) {
	var req GetCourseWithRequestdetails
	var requestdata Request
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	for i := 0; i < len(requestdata); i++{
		if requestdata[i].UserID == authPayload.UserID {
			requestD, err := server.store.GetRequest(ctx, db.GetRequestParam{UserID: authPayload.UserID,})
			if err != nil {
				if errors.Is(err, db.ErrRecordNotFound){
					ctx.JSON(http.StatusNotFound, errorResponse(err))
					return 
				}
				response := gin.H{
					"is_accepted": requestD.Request.IsAccepted,
					"is_pending":  requestD.Request.IsPending,
				}
				ctx.JSON(http.StatusOK, response)
				return 
		}
	}

	arg := db.GetCourseParam{CourseID: req.CourseID}
    course, err := server.store.GetCourse(ctx, arg)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, errorResponse(err))
        return
    }

    response := gin.H{
        "course_details": course,
    }

    ctx.JSON(http.StatusOK, response)
	}
}

// GetCourseByUserRequest defines the request body structure for getting a course
type GetCourseByUserRequest struct {
	UserID   int64 `form:"user_id"`
	CourseID int64 `form:"course_id"`
}

// @Summary Get a course by ID
// @Description Retrieves a course by its ID
// @Produce json
// @Param course_id path int true "Course ID"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /course/byuser [get]
// GetCourseByUser retrieves a course by ID
func (server *Server) GetCourseByUserCourse(ctx *gin.Context) {
	var req GetCourseByUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetCourseByUserParam{
		UserID: req.UserID,
		CourseID: req.CourseID,
	}

	course, err := server.store.GetCourseByUserCourse(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, course)
}