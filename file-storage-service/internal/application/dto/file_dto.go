package dto

import "github.com/google/uuid"

type UploadFileResponse struct {
	ID uuid.UUID `json:"id"`
}
