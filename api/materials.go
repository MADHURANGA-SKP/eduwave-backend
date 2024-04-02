package api

import (
	"net/http"

	db "eduwave-back-end/db/sqlc"

	"github.com/gin-gonic/gin"
)

// CreateMaterialRequest defines the request body structure for creating a material
type CreateMaterialRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	CourseID    int64  `json:"course_id" binding:"required"`
}

// CreateMaterial creates a new material
func (server *Server) CreateMaterial(ctx *gin.Context) {
	var req CreateMaterialRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.CreateMaterialParams{
		Title:       req.Title,
		Description: req.Description,
	}

	material, err := server.store.CreateMaterial(ctx, db.CreateMaterialParam(arg))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, material)
}

// GetMaterialsRequest defines the request body structure for getting materials
type GetMaterialsRequest struct {
	CourseID int64 `uri:"course_id" binding:"required,min=1"`
}

// GetMaterials retrieves materials for a given course ID
func (server *Server) GetMaterials(ctx *gin.Context) {
	var req GetMaterialsRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg :=	db.GetMaterialParam{
		CourseID: req.CourseID,
	}

	materials, err := server.store.ListMaterial(ctx, db.ListMaterialParams{
		CourseID: arg.CourseID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, materials)
}

// UpdateMaterialRequest defines the request body structure for updating a material
type UpdateMaterialRequest struct {
	MaterialID  int64  `json:"material_id" binding:"required,min=1"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	CourseID    int64  `json:"course_id" binding:"required"`
}

// UpdateMaterial updates a material
func (server *Server) UpdateMaterial(ctx *gin.Context) {
	var req UpdateMaterialRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.UpdateMaterialParams{
		MaterialID: req.MaterialID,
		Title:       req.Title,
		Description: req.Description,
		CourseID:    req.CourseID,
	}

	material, err := server.store.UpdateMaterials(ctx, db.UpdateMaterialParam(arg))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, material)
}

// DeleteMaterialRequest defines the request body structure for deleting a material
type DeleteMaterialRequest struct {
	MaterialID int64 `uri:"material_id" binding:"required,min=1"`
	CourseID   int64 `uri:"course_id" binding:"required,min=1"`
}

// DeleteMaterial deletes a material
func (server *Server) DeleteMaterial(ctx *gin.Context) {
	var req DeleteMaterialRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.DeleteMaterialParams{
		MaterialID: req.MaterialID,
		CourseID:   req.CourseID,
	}

	err := server.store.DeleteMaterial(ctx, db.DeleteMaterialParam{
		MaterialID: arg.MaterialID,
		CourseID: arg.CourseID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Material deleted successfully"})
}
