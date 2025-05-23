package service

import (
	"context"
	"files/internal/application/port/out"
	"files/internal/domain"
	"github.com/google/uuid"
	"mime/multipart"
	"strconv"
	"time"
)

type StorageService struct {
	fileDataRepo out.FileDataRepo
	fileRepo     out.FileRepo
}

func NewStorageService(fileRepo out.FileRepo, fileDataRepo out.FileDataRepo) *StorageService {
	return &StorageService{fileRepo: fileRepo, fileDataRepo: fileDataRepo}
}

func (s *StorageService) UploadFile(file multipart.File, header *multipart.FileHeader) (*domain.FileData, error) {
	location := header.Filename + "_" + strconv.FormatInt(time.Now().UnixMilli(), 10)

	err := s.fileRepo.UploadFile(context.Background(), location, file, header.Size)
	if err != nil {
		return nil, err
	}

	fileData := &domain.FileData{
		Location: location,
		ID:       uuid.New(),
		Name:     header.Filename,
	}

	err = s.fileDataRepo.NewFileData(fileData)
	if err != nil {
		return nil, err
	}

	return fileData, nil
}
