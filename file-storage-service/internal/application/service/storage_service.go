package service

import (
	"files/internal/application/port/out"
	"files/internal/domain"
	"mime/multipart"
)

type StorageService struct {
	fileDataRepo out.FileDataRepo
	fileRepo     out.FileRepo
}

func NewStorageService(fileRepo out.FileRepo, fileDataRepo out.FileDataRepo) *StorageService {
	return &StorageService{fileRepo: fileRepo, fileDataRepo: fileDataRepo}
}

func (s *StorageService) UploadFile(file multipart.File, header *multipart.FileHeader) (*domain.FileData, error) {
	return nil, nil
}
