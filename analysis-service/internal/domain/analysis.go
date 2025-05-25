package domain

import "github.com/google/uuid"

type Analysis struct {
	FileID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	Location       string
	ParagraphCount int
	WordCount      int
	CharCount      int
}
