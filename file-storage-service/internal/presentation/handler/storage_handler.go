package handler

import (
	"files/internal/application/dto"
	"files/internal/application/port/in"
	"files/internal/presentation/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StorageHandler struct {
	storageUseCase in.StorageUseCase
}

func NewStorageHandler(storageUseCase in.StorageUseCase) *StorageHandler {
	return &StorageHandler{storageUseCase: storageUseCase}
}

func (h *StorageHandler) GetFile(c *gin.Context) {}

func (h *StorageHandler) UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		response.HandleError(c, err)
		return
	}
	defer file.Close()

	fileData, err := h.storageUseCase.UploadFile(file, header)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.UploadFileResponse{ID: fileData.ID})
}
