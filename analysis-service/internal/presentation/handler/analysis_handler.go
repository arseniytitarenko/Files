package handler

import (
	"files-analysis/internal/application/dto"
	"files-analysis/internal/application/errs"
	"files-analysis/internal/application/port/in"
	"files-analysis/internal/presentation/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type AnalysisHandler struct {
	analysisUseCase in.AnalysisUseCase
}

func NewAnalysisHandler(analysisUseCase in.AnalysisUseCase) *AnalysisHandler {
	return &AnalysisHandler{analysisUseCase: analysisUseCase}
}

func (h *AnalysisHandler) AnalyzeFile(c *gin.Context) {
	id := c.Param("id")

	analysis, err := h.analysisUseCase.AnalyzeFile(c.Request.Context(), id)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.FileAnalysisResponse{
		Location:       analysis.Location,
		ParagraphCount: analysis.ParagraphCount,
		WordCount:      analysis.WordCount,
		CharCount:      analysis.CharCount,
	})
}

func (h *AnalysisHandler) GetWordCloud(c *gin.Context) {
	var queryParams dto.WordCLoudParams
	if err := c.ShouldBindQuery(&queryParams); err != nil {
		response.HandleError(c, errs.InvalidRequest)
		return
	}

	file, err := h.analysisUseCase.GetWordCloud(c.Request.Context(), queryParams.Location)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	defer func(file io.ReadCloser) {
		err = file.Close()
		if err != nil {
			response.HandleError(c, err)
			return
		}
	}(file)

	c.Header("Content-Type", "image/png")
	c.Header("Content-Disposition", fmt.Sprintf(`inline; filename="%s"`, queryParams.Location))
	c.Status(http.StatusOK)
	_, _ = io.Copy(c.Writer, file)
}
