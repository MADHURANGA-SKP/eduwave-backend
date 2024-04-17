package api

import (
	"net/http"

	db "eduwave-back-end/db/sqlc"

	"github.com/gin-gonic/gin"
)

// CreateMaterialRequest defines the request body structure for creating a material
type CreateMaterialRequest struct {
	CourseID    int64  `json:"course_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// @Summary Create a new material
// @Description Create a new material
// @ID create-material
// @Accept  json
// @Produce  json
// @Param request body CreateMaterialRequest true "Create Material Request"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /material [post]
// CreateMaterial creates a new material
func (server *Server) CreateMaterial(ctx *gin.Context) {
	var req CreateMaterialRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateMaterialParams{
		CourseID: req.CourseID,
		Title:       req.Title,
		Description: req.Description,
	}

	material, err := server.store.CreateMaterial(ctx, db.CreateMaterialParam(arg))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, material)
}

// GetMaterialsRequest defines the request body structure for getting materials
type GetMaterialsRequest struct {
	MaterialID int64 `form:"material_id,min=1"`
}

// @Summary Get materials for a course
// @Description Get materials for a course
// @ID get-materials
// @Accept  json
// @Produce  json
// @Param course_id path int true "Course ID"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /material/get [get]
// GetMaterials retrieves materials for a given course ID
func (server *Server) GetMaterials(ctx *gin.Context) {
	var req GetMaterialsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg :=	db.GetMaterialParam{MaterialID: req.MaterialID}

	materials, err := server.store.GetMaterial(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
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

// @Summary Update a material
// @Description Update a material
// @ID update-material
// @Accept  json
// @Produce  json
// @Param material_id path int true "Material ID"
// @Param request body UpdateMaterialRequest true "Update Material Request"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /material/edit [put]
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
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, material)
}

// DeleteMaterialRequest defines the request body structure for deleting a material
type DeleteMaterialRequest struct {
	MaterialID   int64 `form:"material_id"`
}

// @Summary Delete a material
// @Description Delete a material
// @ID delete-material
// @Accept  json
// @Produce  json
// @Param material_id path int true "Material ID"
// @Param course_id path int true "Course ID"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500
// @Router /material/delete [delete]
// DeleteMaterial deletes a material
func (server *Server) DeleteMaterial(ctx *gin.Context) {
	var req DeleteMaterialRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteMaterial(ctx, db.DeleteMaterialParam{
		MaterialID: req.MaterialID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Material deleted successfully"})
}
