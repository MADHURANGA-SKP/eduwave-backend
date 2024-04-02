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

type deleteAdminRequest struct {
	AdminID int64 `uri:"admin_id" binding:"required,min=1"`
}

// @Summary Delete an admin
// @Description Deletes an admin by ID
// @ID delete-admin
// @Param admin_id path int true "Admin ID"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500 
// @Router /admin/{admin_id} [delete]
// deleteAdminRequest defines the request body structure for deleting an admin
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

type getAdminRequest struct {
	AdminID int64 `uri:"admin_id,min=1"`
}

// @Summary Get admin by ID
// @Description Retrieves an admin by ID
// @ID get-admin
// @Param admin_id path int true "Admin ID"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500 
// @Router /admin/{admin_id} [get]
// getAdminRequest defines the request body structure for getting an admin
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

type updateAdminRequest struct {
	AdminID        int64  `json:"admin_id"`
    FullName       string `json:"full_name"`
    UserName       string `json:"user_name"`
    Email          string `json:"email"`
    HashedPassword string `json:"hashed_password"`
}

// @Summary Update an admin
// @Description Updates an admin by ID
// @ID update-admin
// @Param admin_id path int true "Admin ID"
// @Accept json
// @Produce json
// @Param request body updateAdminRequest true "Admin data to update"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500 
// @Router /admin/{admin_id} [put]
// updateAdminRequest defines the request body structure for updating an admin
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

	adminID, err := strconv.Atoi(ctx.Param("admin_id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, errorResponse(err))
        return
    }

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.UpdateAdminParams{
		AdminID: int64(adminID),
		FullName: req.FullName,
		UserName: authPayload.UserName,
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

type CreateAdminRequest struct {
	FullName       string `json:"full_name"`
    UserName       string `json:"user_name"`
    Email          string `json:"email"`
    HashedPassword string `json:"hashed_password"`
}

// @Summary Create an admin
// @Description Creates a new admin
// @ID create-admin
// @Accept json
// @Produce json
// @Param request body CreateAdminRequest true "Admin data to create"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /admin [post]
// CreateAdminRequest defines the request body structure for updating an admin
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
		UserName: authPayload.UserName,
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
