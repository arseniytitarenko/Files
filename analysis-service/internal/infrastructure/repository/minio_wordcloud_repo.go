package repository

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"io"
	"log"
)

type MinioWordCloudRepo struct {
	client *minio.Client
	bucket string
}

func NewMinioWordCloudRepository(client *minio.Client, bucket string) *MinioWordCloudRepo {
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

	return &MinioWordCloudRepo{
		client: client,
		bucket: bucket,
	}
}

func (r *MinioWordCloudRepo) UploadWordCloud(ctx context.Context, name string, data io.Reader, size int64) error {
	_, err := r.client.PutObject(ctx, r.bucket, name, data, size, minio.PutObjectOptions{
		ContentType: "image/png",
	})
	if err != nil {
		return fmt.Errorf("failed to save file to minio: %w", err)
	}
	return nil
}

func (r *MinioWordCloudRepo) GetWordCloud(ctx context.Context, name string) (io.ReadCloser, error) {
	_, err := r.client.StatObject(ctx, r.bucket, name, minio.StatObjectOptions{})
	if err != nil {
		return nil, err
	}
	obj, err := r.client.GetObject(ctx, r.bucket, name, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return obj, nil
}
