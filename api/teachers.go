// teachers controller
package api

import (
	"database/sql"
	db "eduwave-back-end/db/sqlc"
	"eduwave-back-end/token"
	"eduwave-back-end/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createTeacherRequest struct {
    UserID         int64  `json:"user_id"`
    AdminID        int64  `json:"admin_id"`
    FullName       string `json:"full_name"`
    Email          string `json:"email"`
    Qualification  string `json:"qualification"`
    UserName       string `json:"user_name"`
    HashedPassword string `json:"hashed_password"`
    IsActive       bool   `json:"is_active"`
}

// @Summary Create a new teacher
// @Description Create a new teacher with the provided details
// @Tags teachers
// @Accept json
// @Produce json
// @Param request body createTeacherRequest true "Teacher details"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /teacher [post]
func (server *Server) createTeacher(ctx *gin.Context) {
	var req createTeacherRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateTeacherParam{
		UserID: req.UserID,
		AdminID: req.AdminID,
		FullName: req.FullName,
		UserName: authPayload.UserName,
		Email: req.Email,
		Qualification: req.Qualification,
		HashedPassword: hashedPassword,
		IsActive: req.IsActive,
	}

	teacher, err := server.store.CreateTeacher(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, teacher)
}

type GetTeacherRequest struct {
	TeacherID int64 `uri:"teacher_id,min=1"`
}

// @Summary Get a teacher by ID
// @Description Retrieve a teacher by their ID
// @Tags teachers
// @Accept json
// @Produce json
// @Param teacher_id path int true "Teacher ID"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /teacher/{teacher_id} [get]
func (server *Server) getTeacher(ctx *gin.Context) {
	var req GetTeacherRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetTeacherParam{TeacherID: req.TeacherID}

	teacher, err := server.store.GetTeacher(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		return
	}

	ctx.JSON(http.StatusOK, teacher)
}

type ListTeacherParams struct {
	PageID   int32 `form:"page_id,min=1"`
	PageSize int32 `form:"page_size,min=5,max=10"`
}

// @Summary List teachers
// @Description Retrieve a list of teachers
// @Tags teachers
// @Accept json
// @Produce json
// @Param page_id query int false "Page ID" default:"1"
// @Param page_size query int false "Page Size" default:"10"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /teachers [get]
func (server *Server) listTeachers(ctx *gin.Context) {
	var req ListTeacherParams
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListTeacherParams{
        Limit:  req.PageSize,
        Offset: (req.PageID - 1) * req.PageSize,
    }

	teachers, err := server.store.ListTeacher(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, teachers)
}


type UpeateTeacherRequest struct {
	TeacherID      int64  `json:"teacher_id"`
    FullName       string `json:"full_name"`
    Email          string `json:"email"`
    UserName       string `json:"user_name"`
    HashedPassword string `json:"hashed_password"`
    IsActive       bool   `json:"is_active"`
}

// @Summary Update a teacher
// @Description Update an existing teacher with the provided details
// @Tags teachers
// @Accept json
// @Produce json
// @Param teacher_id path int true "Teacher ID"
// @Param request body UpeateTeacherRequest true "Updated teacher details"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /teachers/{teacher_id} [put]
func (server *Server) updateTeacher(ctx *gin.Context) {
	var req UpeateTeacherRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	teacherID, err := strconv.Atoi(ctx.Param("teacher_id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, errorResponse(err))
        return
    }

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.UpdateTeacherParams{
		TeacherID: int64(teacherID),
		FullName: req.FullName,
		Email: req.Email,
		UserName: authPayload.UserName,
		HashedPassword: hashedPassword,
    }

	teacher, err := server.store.UpdateTeacher(ctx, db.UpdateTeacherParam(arg))
	if err != nil {
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, teacher)
}

// @Summary Delete a teacher
// @Description Delete a teacher by their ID
// @Tags teachers
// @Accept json
// @Produce json
// @Param teacher_id path int true "Teacher ID"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /teachers/{teacher_id} [delete]
func (server *Server) deleteTeacher(ctx *gin.Context) {
	teacherID, err := strconv.ParseInt(ctx.Param("teacher_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = server.store.DeleteTeacher(ctx, teacherID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}
