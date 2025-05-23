package repository

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"io"
	"log"
)

type MinioFileRepo struct {
	client *minio.Client
	bucket string
}

func NewMinioFileRepository(client *minio.Client, bucket string) *MinioFileRepo {
	exists, err := client.BucketExists(context.Background(), bucket)
	if err != nil {
		log.Fatalf("failed to check bucket: %v", err)
	}
	if !exists {
		err = client.MakeBucket(context.Background(), bucket, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalf("failed to create bucket: %v", err)
		}
		log.Printf("Bucket %s created", bucket)
	}

	return &MinioFileRepo{
		client: client,
		bucket: bucket,
	}
}

func (r *MinioFileRepo) UploadFile(ctx context.Context, name string, data io.Reader, size int64) error {
	_, err := r.client.PutObject(ctx, r.bucket, name, data, size, minio.PutObjectOptions{
		ContentType: "text/plain; charset=utf-8",
	})
	if err != nil {
		return fmt.Errorf("failed to save file to minio: %w", err)
	}
	return nil
}

func (r *MinioFileRepo) GetFile(ctx context.Context, name string) (io.ReadCloser, error) {
	obj, err := r.client.GetObject(ctx, r.bucket, name, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get file from minio: %w", err)
	}
	return obj, nil
}
