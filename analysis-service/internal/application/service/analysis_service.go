package service

import (
	"context"
	"files-analysis/internal/application/errs"
	"files-analysis/internal/application/port/out"
	"files-analysis/internal/domain"
	"io"
)

type AnalysisService struct {
	analysisRepo  out.AnalysisRepo
	wordcloudRepo out.WordCloudRepo
	fileApi       out.FileApi
	quickChartApi out.QuickChartApi
}

func NewAnalysisService(
	analysisRepo out.AnalysisRepo,
	wordcloudRepo out.WordCloudRepo,
	fileApi out.FileApi,
	quickChartApi out.QuickChartApi) *AnalysisService {
	return &AnalysisService{analysisRepo, wordcloudRepo, fileApi, quickChartApi}
}

func (s *AnalysisService) AnalyzeFile(ctx context.Context, fileID string) (*domain.Analysis, error) {

}

func (s *AnalysisService) GetWordCloud(ctx context.Context, location string) (io.ReadCloser, error) {
	file, err := s.wordcloudRepo.GetWordCloud(ctx, location)
	if err != nil {
		return nil, errs.ErrLocationNotFound
	}
	return file, nil
}
