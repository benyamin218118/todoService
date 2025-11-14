package controller

import (
	"net/http"

	"github.com/benyamin218118/todoService/usecase"
	"github.com/gin-gonic/gin"
)

type StorageController struct {
	uc *usecase.StorageUseCase
}

type UploadFileRes struct {
	ID string `json:"id"`
}

func NewStorageController(storageUC *usecase.StorageUseCase) *StorageController {
	return &StorageController{
		uc: storageUC,
	}
}

// UploadFile godoc
// @Summary Upload a file to S3 / LocalStack
// @Description Uploads a file and returns its generated ID
// @Tags storage
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to upload"
// @Success 201 {object} UploadFileRes "ID of uploaded file"
// @Failure 400 {object} map[string]string "File is required"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /upload [post]
func (ctrl *StorageController) UploadFile(ctx *gin.Context) {
	const MaxFileSize = 10 << 20 // 5 MB
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		ResponseBadRequest(ctx)
		return
	}
	if fileHeader.Size > MaxFileSize {
		ctx.JSON(http.StatusBadRequest, map[string]any{"message": "file too large, max: 10MB"})
		return
	}

	allowedTypes := map[string]bool{
		"image/png":  true,
		"image/jpeg": true,
		"text/plain": true,
	}
	// TODO: check the magic code instead
	if _, ok := allowedTypes[fileHeader.Header.Get("Content-Type")]; !ok {
		ctx.JSON(http.StatusBadRequest, map[string]any{"message": "Invalid File Type, jpg/png/text only."})
		return
	}

	file, err := fileHeader.Open()
	if HandleIfError(ctx, err) {
		return
	}

	defer file.Close()

	id, err := ctrl.uc.Upload(file, fileHeader.Filename)
	if HandleIfError(ctx, err) {
		return
	}

	Response(ctx, http.StatusCreated, UploadFileRes{
		ID: *id,
	})
}
