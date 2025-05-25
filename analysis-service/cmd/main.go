package main

import (
	"files-analysis/internal/application/service"
	"files-analysis/internal/domain"
	"files-analysis/internal/infrastructure/database"
	"files-analysis/internal/infrastructure/externalapi"
	"files-analysis/internal/infrastructure/repository"
	"files-analysis/internal/infrastructure/storage"
	"files-analysis/internal/presentation/handler"
	"files-analysis/internal/presentation/router"
	"log"
)

func main() {
	db := database.NewPostgres()
	minioClient := storage.NewMinioClient()
	err := db.AutoMigrate(&domain.Analysis{})
	if err != nil {
		log.Fatal(err)
	}

	analysisRepo := repository.NewPgAnalysisRepo(db)
	wordCloudRepo := repository.NewMinioWordCloudRepository(minioClient, "wordcloud")

	fileApi := externalapi.NewFileApiClient()
	quickChartApi := externalapi.NewQuickChartApiClient()

	analysisService := service.NewAnalysisService(analysisRepo, wordCloudRepo, fileApi, quickChartApi)

	analysisHandler := handler.NewAnalysisHandler(analysisService)

	r := router.SetupRouter(analysisHandler)
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
