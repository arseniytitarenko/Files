package externalapi

import (
	"bytes"
	"encoding/json"
	"files-analysis/internal/application/errs"
	"io"
	"net/http"
)

type WordCloudParams struct {
	Format          string   `json:"format"`
	Text            string   `json:"text"`
	Width           int      `json:"width"`
	Height          int      `json:"height"`
	Colors          []string `json:"colors"`
	FontStyle       string   `json:"fontFamily"`
	RemoveStopwords bool     `json:"removeStopwords"`
}

type QuickChartApiClient struct {
	httpClient *http.Client
	endpoint   string
}

func NewQuickChartApiClient() *QuickChartApiClient {
	return &QuickChartApiClient{endpoint: "https://quickchart.io/wordcloud", httpClient: &http.Client{}}
}

func (c *QuickChartApiClient) GetWordCloud(text string) (io.Reader, int64, error) {
	params := WordCloudParams{
		Format:          "png",
		Text:            text,
		Width:           1000,
		Height:          1000,
		Colors:          []string{"#FF5733", "#33FF57", "#3357FF"},
		FontStyle:       "sans-serif",
		RemoveStopwords: true,
	}

	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, 0, err
	}

	resp, err := c.httpClient.Post(
		c.endpoint,
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return nil, 0, errs.ExternalApiError
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, 0, errs.ExternalApiError
	}

	return resp.Body, resp.ContentLength, nil
}
