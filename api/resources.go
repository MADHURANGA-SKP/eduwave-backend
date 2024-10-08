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

//getResourceHandler creates an link to located the exact resource file from the server storage
func (server *Server) getResourceHandler(ctx *gin.Context) {
    filename := ctx.Param("filename")
    serverAddress := server.config.FileSource
    Image := fmt.Sprintf("%s/resources/%s", serverAddress, filename)
    ctx.JSON(http.StatusOK, Image)
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


// ListResourceRequest contains the impurt parameters for list rolebased user data
type ListResourceRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=10,max=100"`
}

// @Summary ListResource
// @Description ListResource all avalible resource
// @ID list-teacher
// @Accept  json
// @Produce  json
// @Param request body ListResourceRequest true "teacher list request"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /resources/get [get]
func (server *Server) ListResource(ctx *gin.Context) {
	var req ListResourceRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.ListResourceParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	userlist, err := server.store.ListResource(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, userlist)
}

// ListResourceRequest contains the impurt parameters for list rolebased user data
type ListResourceByMaterialRequest struct {
	MaterialID int64 `form:"material_id"`
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=10,max=100"`
}

// @Summary ListResource
// @Description List all Resources that belongs to material
// @ID list-teacher
// @Accept  json
// @Produce  json
// @Param request body ListResourceRequest true "teacher list request"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /resources/bymaterial [get]
func (server *Server) ListResourceByMaterial(ctx *gin.Context) {
	var req ListResourceByMaterialRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.ListResourceByMaterialParams{
		MaterialID: req.MaterialID,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	userlist, err := server.store.ListResourceByMaterial(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, userlist)
}