package out

import (
	"files/internal/domain"
	"github.com/google/uuid"
)

type FileDataRepo interface {
	NewFileData(fileData *domain.FileData) error
	GetFileData(fileDataID uuid.UUID) (*domain.FileData, error)
}
