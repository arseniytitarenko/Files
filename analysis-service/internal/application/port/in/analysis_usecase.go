package in

import (
	"context"
	"files-analysis/internal/domain"
	"io"
)

type AnalysisUseCase interface {
	AnalyzeFile(ctx context.Context, fileID string) (*domain.Analysis, error)
	GetWordCloud(ctx context.Context, location string) (io.ReadCloser, error)
}
