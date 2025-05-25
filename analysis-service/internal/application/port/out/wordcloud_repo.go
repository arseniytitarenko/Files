package out

import (
	"context"
	"io"
)

type WordCloudRepo interface {
	UploadWordCloud(ctx context.Context, name string, data io.Reader, size int64) error
	GetWordCloud(ctx context.Context, name string) (io.ReadCloser, error)
}
