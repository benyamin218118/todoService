package controller

import (
	"io"
	"net/http"

	"github.com/benyamin218118/todoService/usecase"
	"github.com/gin-gonic/gin"
)

type StorageController struct {
	uc *usecase.StorageUseCase
}

type UploadFileReq struct {
	File io.Reader
	Name string
}

type UploadFileRes struct {
	ID string `json:"id"`
}

func NewStorageController(storageUC *usecase.StorageUseCase) *StorageController {
	return &StorageController{
		uc: storageUC,
	}
}

func (ctrl *StorageController) UploadFile(ctx *gin.Context) {
	var input UploadFileReq

	id, err := ctrl.uc.Upload(input.File, input.Name)
	if HandleIfError(ctx, err) {
		return
	}

	Response(ctx, http.StatusCreated, UploadFileRes{
		ID: *id,
	})
}
