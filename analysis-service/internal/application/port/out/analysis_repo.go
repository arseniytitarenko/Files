package out

import (
	"files-analysis/internal/domain"
	"github.com/google/uuid"
)

type AnalysisRepo interface {
	NewAnalysis(analysis *domain.Analysis) error
	GetAnalysis(fileID uuid.UUID) (*domain.Analysis, error)
}
