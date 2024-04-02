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

// @Summary Create a new resource
// @Description Creates a new resource
// @ID create-resource
// @Accept  json
// @Produce  json
// @Param request body createResourceRequest true "Create Resource Request"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /resource [post]
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
	ResourceID int64 `json:"resource_id"`
    MaterialID int64 `json:"material_id"`
}

// @Summary Delete a resource
// @Description Deletes a resource
// @ID delete-resource
// @Produce  json
// @Param resource_id path int true "Resource ID"
// @Param material_id path int true "Material ID"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /resource/{resource_id}/{material_id} [delete]
// deleteResource deletes a resource
func (server *Server) deleteResource(ctx *gin.Context) {
	var req deleteResourceRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.DeleteResourceParams{
		ResourceID: req.ResourceID,
		MaterialID: req.MaterialID,
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
	MaterialID int64 `json:"matirial_id"`
    ResourceID int64         `json:"resource_id"`
}

// @Summary Get a resource
// @Description Retrieves a resource
// @ID get-resource
// @Produce  json
// @Param material_id path int true "Material ID"
// @Param resource_id path int true "Resource ID"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /resource/{resource_id}/{material_id} [get]
// getResource retrieves a resource
func (server *Server) getResource(ctx *gin.Context) {
	var req getResourceRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resource, err := server.store.GetResource(ctx, db.GetResourceParam{
		MaterialID: req.MaterialID,
		ResourceID: req.ResourceID,
	})
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

// @Summary Update a resource
// @Description Updates a resource
// @ID update-resource
// @Accept  json
// @Produce  json
// @Param request body updateResourceRequest true "Update Resource Request"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /resource/{resource_id} [put]
// updateResourceRequest defines the request body structure for updating a resource
type updateResourceRequest struct {
	MaterialID sql.NullInt64 `json:"matirial_id"`
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

	resource, err := server.store.UpdateResource(ctx, db.UpdateResourceParam{
		MaterialID: req.ResourceID,
		Title: req.Title,
		Type: req.Type,
		ContentUrl: req.ContentUrl,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, resource)
}
