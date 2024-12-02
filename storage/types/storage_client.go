package types

import "context"

type StorageType string

const (
	StorageTypeMinio StorageType = "minio"
	StorageTypeS3    StorageType = "s3"
	StorageTypeGCS   StorageType = "gcs"
)

type StorageClient interface {
	CreateBucket(ctx context.Context, bucketName string) error
	PutObject(ctx context.Context, bucket, object, filePath string) error
	GetObject(ctx context.Context, bucket, object string) ([]byte, error)
}
