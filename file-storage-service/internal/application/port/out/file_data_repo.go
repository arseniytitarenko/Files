package out

import "files/internal/domain"

type FileDataRepo interface {
	NewFileData(benefit *domain.FileData) error
	GetFileData(fileDataID string) (*domain.FileData, error)
}
