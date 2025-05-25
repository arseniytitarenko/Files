package externalapi

import (
	"bytes"
	"encoding/json"
	"files-analysis/internal/application/errs"
	"io"
	"log"
	"net/http"
)

type WordCloudParams struct {
	Format          string `json:"format"`
	Text            string `json:"text"`
	Width           int    `json:"width"`
	Height          int    `json:"height"`
	FontStyle       string `json:"fontFamily"`
	RemoveStopwords bool   `json:"removeStopwords"`
	MaxNumWords     int    `json:"maxNumWords"`
	MinWordLength   int    `json:"minWordLength"`
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
		FontStyle:       "sans-serif",
		RemoveStopwords: true,
		MaxNumWords:     50,
		MinWordLength:   3,
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
		log.Println(err)
		return nil, 0, errs.ExternalApiError
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(resp.Status)
		return nil, 0, errs.ExternalApiError
	}

	buf := new(bytes.Buffer)
	n, err := io.Copy(buf, resp.Body)
	if err != nil {
		log.Println("Error reading body:", err)
		return nil, 0, errs.ExternalApiError
	}

	return bytes.NewReader(buf.Bytes()), n, nil
}
