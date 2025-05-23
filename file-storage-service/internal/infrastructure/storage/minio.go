package storage

import (
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinioClient() *minio.Client {
	client, err := minio.New(os.Getenv("MINIO_ENDPOINT"), &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("MINIO_USER"), os.Getenv("MINIO_PASSWORD"), ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalf("failed to init minio client: %v", err)
	}
	return client
}
