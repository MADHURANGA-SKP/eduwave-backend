package api

import (
	"database/sql"
	"net/http"
	"strconv"

	db "eduwave-back-end/db/sqlc"
	"eduwave-back-end/token"
	"eduwave-back-end/util"

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
	AdminID int64 `json:"admin_id" binding:"required,min=1"`
}

// getAdmin retrieves an admin
func (server *Server) getAdmin(ctx *gin.Context) {
	var req getAdminRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetAdminParam{
		AdminID: req.AdminID,
	}

	admin, err := server.store.GetAdmin(ctx, arg)
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
	AdminID        int64  `json:"admin_id"`
    FullName       string `json:"full_name"`
    UserName       string `json:"user_name"`
    Email          string `json:"email"`
    HashedPassword string `json:"hashed_password"`
}

// updateAdmin updates an admin
func (server *Server) updateAdmin(ctx *gin.Context) {
	var req updateAdminRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	adminID, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, errorResponse(err))
        return
    }

	arg := db.UpdateAdminParams{
		AdminID: int64(adminID),
		FullName: req.FullName,
		UserName: req.UserName,
		Email: req.Email,
		HashedPassword: hashedPassword,
	}

	admin, err := server.store.UpdateAdmin(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, admin)
}

// CreateAdminRequest defines the request body structure for updating an admin
type CreateAdminRequest struct {
	FullName       string `json:"full_name"`
    UserName       string `json:"user_name"`
    Email          string `json:"email"`
    HashedPassword string `json:"hashed_password"`
}

// CreateAdmin Creates an admin
func (server *Server) createAdmin(ctx *gin.Context) {
	var req CreateAdminRequest
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
	arg := db.CreateAdminParams{
		FullName: req.FullName,
		UserName: authPayload.Username,
		Email: req.Email,
		HashedPassword: hashedPassword,
	}

	admin, err := server.store.CreateAdmin(ctx, db.CreateAdminParam(arg))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, admin )
}
