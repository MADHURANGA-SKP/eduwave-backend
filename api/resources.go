// resources controller
package api

import (
	"database/sql"
	"net/http"

	db "eduwave-back-end/db/sqlc"

	"github.com/gin-gonic/gin"
)

// createResourceRequest defines the request body structure for creating a new resource
type createResourceRequest struct {
	Title      string          `json:"title" binding:"required"`
	Type       db.TypeResource `json:"type" binding:"required"`
	ContentUrl string          `json:"content_url" binding:"required"`
}

// createResource creates a new resource
func (server *Server) createResource(ctx *gin.Context) {
	var req createResourceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateResourceParams{
		Title:      req.Title,
		Type:       req.Type,
		ContentUrl: req.ContentUrl,
	}

	resource, err := server.store.CreateResource(ctx, db.CreateResourceParam(arg))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, resource)
}

// deleteResourceRequest defines the request body structure for deleting a resource
type deleteResourceRequest struct {
	ResourceID int64         `json:"resource_id"`
    MatirialID sql.NullInt64 `json:"matirial_id"`
}

// deleteResource deletes a resource
func (server *Server) deleteResource(ctx *gin.Context) {
	var req deleteResourceRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.DeleteResourceParams{
		ResourceID: req.ResourceID,
		MatirialID: req.MatirialID,
	}

	err := server.store.DeleteResource(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}

// getResourceRequest defines the request body structure for getting a resource
type getResourceRequest struct {
	MatirialID sql.NullInt64 `json:"matirial_id"`
    ResourceID int64         `json:"resource_id"`
}

// getResource retrieves a resource
func (server *Server) getResource(ctx *gin.Context) {
	var req getResourceRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetResourceParams{
		MatirialID: req.MatirialID,
		ResourceID: req.ResourceID,
	}

	resource, err := server.store.GetResource(ctx, db.GetResourceParam(arg))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, resource)
}

// updateResourceRequest defines the request body structure for updating a resource
type updateResourceRequest struct {
	MatirialID sql.NullInt64 `json:"matirial_id"`
	ResourceID int64         `json:"resource_id"`
	Title      string        `json:"title"`
	Type       db.TypeResource  `json:"type"`
	ContentUrl string        `json:"content_url"`
}

// updateResource updates a resource
func (server *Server) updateResource(ctx *gin.Context) {
	var req updateResourceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateResourceParams{
		MatirialID: req.MatirialID,
		ResourceID: req.ResourceID,
		Title:      req.Title,
		Type:       req.Type,
		ContentUrl: req.ContentUrl,
	}

	resource, err := server.store.UpdateResource(ctx, db.UpdateResourceParam{
		Title: arg.Title,
		Type: arg.Type,
		ContentUrl:arg.ContentUrl,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, resource)
}
