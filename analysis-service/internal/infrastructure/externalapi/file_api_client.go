package externalapi

import (
	"files-analysis/internal/application/errs"
	"io"
	"mime"
	"net/http"
)

type FileApiClient struct {
	httpClient *http.Client
	endpoint   string
}

func NewFileApiClient() *FileApiClient {
	return &FileApiClient{endpoint: "http://storage-service:8080/file", httpClient: &http.Client{}}
}

func (f *FileApiClient) GetFile(id string) (string, string, error) {
	resp, err := f.httpClient.Get(f.endpoint + "/" + id)
	if err != nil {
		return "", "", errs.ExternalApiError
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", errs.ExternalApiError
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	cd := resp.Header.Get("Content-Disposition")
	_, params, err := mime.ParseMediaType(cd)
	if err != nil {
		return "", "", err
	}

	filename := params["filename"]
	return string(body), filename, nil
}
