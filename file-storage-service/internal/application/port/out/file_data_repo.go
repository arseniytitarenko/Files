package out

import "files/internal/domain"

type FileDataRepo interface {
	NewFileData(fileData *domain.FileData) error
	GetFileData(fileDataID string) (*domain.FileData, error)
}
