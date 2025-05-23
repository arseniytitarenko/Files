package router

import (
	"files/internal/presentation/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(storageHandler *handler.StorageHandler) *gin.Engine {
	r := gin.Default()
	r.POST("/file", storageHandler.UploadFile)
	r.GET("/file/:id", storageHandler.GetFile)
	return r
}
