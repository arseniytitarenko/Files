package domain

import "github.com/google/uuid"

type FileData struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name     string
	Location string
}
