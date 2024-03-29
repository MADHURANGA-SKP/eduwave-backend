// teachers controller
package api

import (
	"database/sql"
	db "eduwave-back-end/db/sqlc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createTeacherRequest struct {
	FullName       string `json:"full_name" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	UserName       sql.NullString `json:"user_name"`
	HashedPassword string `json:"hashed_password" binding:"required"`
	IsActive       bool   `json:"is_active"`
}

func (server *Server) createTeacher(ctx *gin.Context) {
	var req createTeacherRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateTeacherParams{
		FullName:       req.FullName,
		Email:          req.Email,
		UserName:       req.UserName,
		HashedPassword: req.HashedPassword,
		IsActive:       req.IsActive,
	}

	teacher, err := server.store.CreateTeacher(ctx, db.CreateTeacherParam(arg))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, teacher)
}

type GetTeacherRequest struct {
	TeacherID int64 `json:"teacher_id"`
}

func (server *Server) getTeacher(ctx *gin.Context) {
	var req GetTeacherRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetTeacherParam{req.TeacherID}

	teacher, err := server.store.GetTeacher(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, teacher)
}

type ListTeacherParams struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
    PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

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

func (server *Server) updateTeacher(ctx *gin.Context) {
	var req createTeacherRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	teacherID, err := strconv.ParseInt(ctx.Param("teacher_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateTeacherParams{
		TeacherID:      teacherID,
		FullName:       req.FullName,
		Email:          req.Email,
		UserName:       req.UserName,
		HashedPassword: req.HashedPassword,
		IsActive:       req.IsActive,
	}

	teacher, err := server.store.UpdateTeacher(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, teacher)
}

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
