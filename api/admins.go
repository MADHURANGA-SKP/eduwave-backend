package api

import (
	"database/sql"
	"net/http"

	db "eduwave-back-end/db/sqlc"

	"github.com/gin-gonic/gin"
)

// deleteAdminRequest defines the request body structure for deleting an admin
type deleteAdminRequest struct {
	AdminID int64 `uri:"admin_id" binding:"required,min=1"`
}

// deleteAdmin deletes an admin
func (server *Server) deleteAdmin(ctx *gin.Context) {
	var req deleteAdminRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteAdmin(ctx, req.AdminID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Admin deleted successfully"})
}

// getAdminRequest defines the request body structure for getting an admin
type getAdminRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// getAdmin retrieves an admin
func (server *Server) getAdmin(ctx *gin.Context) {
	var req getAdminRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	admin, err := server.store.GetAdmin(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, admin)
}

// updateAdminRequest defines the request body structure for updating an admin
type updateAdminRequest struct {
	AdminID  int64  `json:"admin_id" binding:"required,min=1"`
	UserName string `json:"user_name"`
}

// updateAdmin updates an admin
func (server *Server) updateAdmin(ctx *gin.Context) {
	var req updateAdminRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateAdminParams{
		AdminID:  req.AdminID,
		UserName: sql.NullString{String: req.UserName, Valid: req.UserName != ""},
	}

	admin, err := server.store.UpdateAdmin(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, admin)
}
