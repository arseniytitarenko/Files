package out

import (
	"context"
	"io"
)

type FileRepo interface {
	UploadFile(ctx context.Context, name string, data io.Reader, size int64) error
	GetFile(ctx context.Context, name string) (io.ReadCloser, error)
}
