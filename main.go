package main

import (
	"context"
	"fmt"
	"v1/storage/gcs"
	"v1/storage/minio"
	_types "v1/storage/types"
)

// NewStorageClient initializes and returns either a GCS or MinIO client based on the config.
// params: context, storageType (minio/gcs), config
func NewStorageClient(ctx context.Context, storageType _types.StorageType, config interface{}) (_types.StorageClient, error) {
	switch storageType {
	case _types.StorageTypeGCS:
		gcsConfig, ok := config.(gcs.Config)
		if !ok {
			return nil, fmt.Errorf("invalid config for GCS client")
		}
		return gcs.NewGcsClient(ctx, gcsConfig)
	case _types.StorageTypeMinio:
		minioConfig, ok := config.(minio.Config)
		if !ok {
			return nil, fmt.Errorf("invalid config for MinIO client")
		}
		return minio.NewMinioClient(minioConfig)
	default:
		return nil, fmt.Errorf("unsupported storage type: %s", storageType)
	}
}
