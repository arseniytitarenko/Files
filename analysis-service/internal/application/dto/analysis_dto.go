package dto

type FileAnalysisResponse struct {
	Location       string `json:"location"`
	ParagraphCount int    `json:"paragraph_count"`
	WordCount      int    `json:"word_count"`
	CharCount      int    `json:"char_count"`
}

type WordCLoudParams struct {
	Location string `form:"location" binding:"required"`
}
