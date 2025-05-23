package in

import (
	"files/internal/domain"
	"mime/multipart"
)

type StorageUseCase interface {
	UploadFile(file multipart.File, header *multipart.FileHeader) (*domain.FileData, error)
	//GetFileByID(id string) *dto.FileResponse
}
