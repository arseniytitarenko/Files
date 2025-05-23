package in

import (
	"context"
	"files/internal/domain"
	"io"
	"mime/multipart"
)

type StorageUseCase interface {
	UploadFile(ctx context.Context, file multipart.File, header *multipart.FileHeader) (*domain.FileData, error)
	GetFile(ctx context.Context, id string) (io.ReadCloser, *domain.FileData, error)
}
