package repository

import (
	"files-analysis/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PgAnalysisRepo struct {
	db *gorm.DB
}

func NewPgAnalysisRepo(db *gorm.DB) *PgAnalysisRepo {
	return &PgAnalysisRepo{db: db}
}

func (r *PgAnalysisRepo) GetAnalysis(fileID uuid.UUID) (*domain.Analysis, error) {
	var analysis domain.Analysis
	if err := r.db.Where("file_id = ?", fileID).First(&analysis).Error; err != nil {
		return nil, err
	}
	return &analysis, nil
}

func (r *PgAnalysisRepo) NewAnalysis(analysis *domain.Analysis) error {
	return r.db.Create(analysis).Error
}
