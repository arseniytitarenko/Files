package service

import (
	"context"
	"files-analysis/internal/application/port/out"
	"files-analysis/internal/domain"
	"io"
)

type AnalysisService struct {
	analysisRepo  out.AnalysisRepo
	wordcloudRepo out.WordCloudRepo
}

func NewAnalysisService(analysisRepo out.AnalysisRepo, wordcloudRepo out.WordCloudRepo) *AnalysisService {
	return &AnalysisService{analysisRepo: analysisRepo, wordcloudRepo: wordcloudRepo}
}

func (s *AnalysisService) AnalyzeFile(ctx context.Context, fileID string) (*domain.Analysis, error) {
	return nil, nil
}
func (s *AnalysisService) GetWordCloud(ctx context.Context, location string) (io.ReadCloser, error) {
	return nil, nil
}
