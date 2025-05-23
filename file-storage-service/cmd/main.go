package main

import (
	"files/internal/application/service"
	"files/internal/domain"
	"files/internal/infrastructure/database"
	"files/internal/infrastructure/repository"
	"files/internal/infrastructure/storage"
	"files/internal/presentation/handler"
	"files/internal/presentation/router"
	"log"
)

func main() {
	db := database.NewPostgres()
	client := storage.NewMinioClient()
	err := db.AutoMigrate(&domain.FileData{})
	if err != nil {
		log.Fatal(err)
	}

	fileDataRepo := repository.NewPgFileDataRepo(db)
	fileRepo := repository.NewMinioFileRepository(client, "files")

	storageService := service.NewStorageService(fileRepo, fileDataRepo)

	storageHandler := handler.NewStorageHandler(storageService)

	r := router.SetupRouter(storageHandler)
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
