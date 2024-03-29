// teachers controller
package api

import (
	db "eduwave-back-end/db/sqlc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createTeacherRequest struct {
	FullName       string `json:"full_name" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	UserName       string `json:"user_name"`
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

	teacher, err := server.store.CreateTeacher(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, teacher)
}

func (server *Server) getTeacher(ctx *gin.Context) {
	teacherID, err := strconv.ParseInt(ctx.Param("teacher_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	teacher, err := server.store.GetTeacher(ctx, teacherID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, teacher)
}

func (server *Server) listTeachers(ctx *gin.Context) {
	adminID, err := strconv.ParseInt(ctx.Query("admin_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	limit, err := strconv.ParseInt(ctx.DefaultQuery("limit", "10"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	offset, err := strconv.ParseInt(ctx.DefaultQuery("offset", "0"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListTeacherParams{
		AdminID: adminID,
		Limit:   int32(limit),
		Offset:  int32(offset),
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
