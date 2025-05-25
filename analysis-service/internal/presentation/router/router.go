package router

import (
	"files-analysis/internal/presentation/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(storageHandler *handler.AnalysisHandler) *gin.Engine {
	r := gin.Default()
	r.GET("/file/:id", storageHandler.AnalyzeFile)
	r.GET("/wordcloud", storageHandler.GetWordCloud)
	return r
}
