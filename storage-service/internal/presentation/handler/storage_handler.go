package handler

import (
	"files/internal/application/dto"
	"files/internal/application/port/in"
	"files/internal/presentation/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
)

type StorageHandler struct {
	storageUseCase in.StorageUseCase
}

func NewStorageHandler(storageUseCase in.StorageUseCase) *StorageHandler {
	return &StorageHandler{storageUseCase: storageUseCase}
}

func (h *StorageHandler) GetFile(c *gin.Context) {
	id := c.Param("id")

	file, fileData, err := h.storageUseCase.GetFile(c.Request.Context(), id)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	defer func(file io.ReadCloser) {
		err = file.Close()
		if err != nil {
			response.HandleError(c, err)
			return
		}
	}(file)

	c.Header("Content-Type", "text/plain; charset=utf-8")
	c.Header("Content-Disposition", fmt.Sprintf(`inline; filename="%s"`, fileData.Name))
	c.Status(http.StatusOK)
	_, _ = io.Copy(c.Writer, file)
}

func (h *StorageHandler) UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		response.HandleError(c, err)
		return
	}

	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			response.HandleError(c, err)
			return
		}
	}(file)

	fileData, err := h.storageUseCase.UploadFile(c.Request.Context(), file, header)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.UploadFileResponse{ID: fileData.ID})
}
