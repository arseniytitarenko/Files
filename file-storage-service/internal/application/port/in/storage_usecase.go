package in

import "mime/multipart"

type StorageUseCase interface {
	UploadFile(file multipart.File, header *multipart.FileHeader) error
	GetFileByID(id string) error
}
