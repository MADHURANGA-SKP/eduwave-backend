// resources controller
package api

import (
	"database/sql"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	db "eduwave-back-end/db/sqlc"

	"github.com/gin-gonic/gin"
)

// uploadSingleFile handles uploading a single file
func(server *Server) uploadSingleFile(ctx *gin.Context, file multipart.File, header *multipart.FileHeader) (string, error) {
	// Validate file extension (optional)
	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	  ".txt":  true,
	  ".docx": true,
	  ".pdf":  true,
	  ".mp4":  true,
	}
	fileExt := filepath.Ext(header.Filename)
	if !allowedExtensions[fileExt] {
	  return "", fmt.Errorf("unsupported file extension: %s", fileExt)
	}
  
	// Generate unique filename
	originalFileName := strings.TrimSuffix(filepath.Base(header.Filename), filepath.Ext(header.Filename))
	now := time.Now()
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
  
	// Create upload directory if it doesn't exist
	uploadDir := "uploads/resources"
	if err := os.MkdirAll(uploadDir, 0755); err != nil{
	  return "", err
	}
  
	// Create destination file path
	filePath := filepath.Join(uploadDir, filename)
  
	// Save uploaded file
	out, err := os.Create(filePath)
	if err != nil {
	  return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
	  return "", err
	}
  
	return filename, nil
  }

type TypeResource string

// createResourceRequest defines the request body structure for creating a new resource
type createResourceRequest struct {
	MaterialID int64        `json:"material_id"`
	Title      string       `form:"title"`
	Type       TypeResource `form:"type"`
	ContentUrl string       `form:"content_url"`
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
// @Router /resource/:material_id [post]
// createResource creates a new resource
func (server *Server) createResource(ctx *gin.Context) {
	var req createResourceRequest
  if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	materialID, err := strconv.Atoi(ctx.Param("material_id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, errorResponse(err))
        return
    }

	// Handle image upload (if included in the request)
	var imageFilename string
	file, header, err := ctx.Request.FormFile("file") 
	if err == nil {
	  imageFilename, err = server.uploadSingleFile(ctx, file, header)
	  if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	  }
	}
	// Update request struct with image filename (if uploaded)
	req.ContentUrl = imageFilename 

	arg := db.CreateResourceParams{
		MaterialID: int64(materialID),
		Title:      req.Title,
		Type:       db.TypeResource(req.Type),
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
	ResourceID int64 `form:"resource_id"`
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
// @Router /resource/delete [delete]
// deleteResource deletes a resource
func (server *Server) deleteResource(ctx *gin.Context) {
	var req deleteResourceRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteResource(ctx, db.DeleteResourceParam{
		ResourceID: req.ResourceID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Resource deleted successfully"})
}

// getResourceRequest defines the request body structure for getting a resource
type getResourceRequest struct {
    ResourceID int64         `form:"resource_id,min=1"`
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
// @Router /resource/get [get]
// getResource retrieves a resource
func (server *Server) getResource(ctx *gin.Context) {
	var req getResourceRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resource, err := server.store.GetResource(ctx, db.GetResourceParam{
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
// @Router /resource/edit [put]
// updateResourceRequest defines the request body structure for updating a resource
type updateResourceRequest struct {
	MaterialID int64        `json:"material_id"`
    ResourceID int64        `json:"resource_id"`
    Title      string       `json:"title"`
    Type       TypeResource `json:"type"`
    ContentUrl string       `json:"content_url"`
}

// updateResource updates a resource
func (server *Server) updateResource(ctx *gin.Context) {
	var req updateResourceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resourceID, err := strconv.Atoi(ctx.Param("resource_id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, errorResponse(err))
        return
    }

	// Handle image upload (if included in the request)
	var imageFilename string
	file, header, err := ctx.Request.FormFile("image") 
	if err == nil {
	  imageFilename, err = server.uploadSingleFile(ctx, file, header)
	  if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	  }
	}
  
	// Update request struct with image filename (if uploaded)
	req.ContentUrl = imageFilename 

	resource, err := server.store.UpdateResource(ctx, db.UpdateResourceParam{
		ResourceID: int64(resourceID),
		Title: req.Title,
		Type: db.TypeResource(req.Type),
		ContentUrl: req.ContentUrl,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, resource)
}
