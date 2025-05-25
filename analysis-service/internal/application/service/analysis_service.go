package service

import (
	"context"
	"errors"
	"files-analysis/internal/application/errs"
	"files-analysis/internal/application/port/out"
	"files-analysis/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"io"
	"log"
	"strconv"
	"strings"
	"time"
	"unicode"
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
	uuidFileID, err := uuid.Parse(fileID)
	if err != nil {

		return nil, errs.InvalidID
	}

	analysis, err := s.analysisRepo.GetAnalysis(uuidFileID)
	if err == nil || !errors.Is(err, gorm.ErrRecordNotFound) {
		return analysis, err
	}

	fileText, fileName, err := s.fileApi.GetFile(fileID)
	if err != nil {
		return nil, err
	}

	log.Println("File got")

	wordCloud, size, err := s.quickChartApi.GetWordCloud(fileText)
	if err != nil {
		return nil, err
	}

	location := strconv.FormatInt(time.Now().UnixMilli(), 10) + "_" + fileName

	err = s.wordcloudRepo.UploadWordCloud(ctx, location, wordCloud, size)
	if err != nil {
		return nil, err
	}

	analysis = &domain.Analysis{
		FileID:   uuidFileID,
		Location: location,
	}
	analyzeText(fileText, analysis)

	err = s.analysisRepo.NewAnalysis(analysis)
	if err != nil {
		return nil, err
	}
	return analysis, nil
}

func (s *AnalysisService) GetWordCloud(ctx context.Context, location string) (io.ReadCloser, error) {
	file, err := s.wordcloudRepo.GetWordCloud(ctx, location)
	if err != nil {
		return nil, errs.ErrLocationNotFound
	}
	return file, nil
}

func analyzeText(text string, analysis *domain.Analysis) {
	words := 0
	inWord := false

	for _, r := range text {
		if unicode.IsSpace(r) {
			inWord = false
		} else {
			if !inWord {
				words++
				inWord = true
			}
		}
	}

	chars := len([]rune(text))
	paragraphs := len(strings.Split(strings.TrimSpace(text), "\n\n"))

	analysis.ParagraphCount = paragraphs
	analysis.WordCount = words
	analysis.CharCount = chars
}
