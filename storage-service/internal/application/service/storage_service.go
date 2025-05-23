package service

import (
	"context"
	"files/internal/application/errs"
	"files/internal/application/port/out"
	"files/internal/domain"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

type StorageService struct {
	fileDataRepo out.FileDataRepo
	fileRepo     out.FileRepo
}

func NewStorageService(fileRepo out.FileRepo, fileDataRepo out.FileDataRepo) *StorageService {
	return &StorageService{fileRepo: fileRepo, fileDataRepo: fileDataRepo}
}

func (s *StorageService) UploadFile(ctx context.Context, file multipart.File, header *multipart.FileHeader) (*domain.FileData, error) {
	if !strings.HasSuffix(header.Filename, ".txt") {
		return nil, errs.InvalidFileFormat
	}

	location := strconv.FormatInt(time.Now().UnixMilli(), 10) + "_" + header.Filename

	err := s.fileRepo.UploadFile(ctx, location, file, header.Size)
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

func (s *StorageService) GetFile(ctx context.Context, id string) (io.ReadCloser, *domain.FileData, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return nil, nil, errs.InvalidID
	}

	fileData, err := s.fileDataRepo.GetFileData(uuidID)
	if err != nil {
		return nil, nil, errs.FileNotFound
	}

	file, err := s.fileRepo.GetFile(ctx, fileData.Location)
	return file, fileData, nil
}
