package storage

import (
	"context"
	"fmt"
	_types "v1/types"
	_gcs "v1/v1/gcs"
	_minio "v1/v1/minio"
)

// NewStorageClient initializes and returns either a GCS or MinIO client based on the config.
// params: context, storageType (minio/gcs), config
func NewStorageClient(ctx context.Context, storageType _types.StorageType, config interface{}) (_types.StorageClient, error) {
	switch storageType {
	case _types.StorageTypeGCS:
		gcsConfig, ok := config.(_gcs.Config)
		if !ok {
			return nil, fmt.Errorf("invalid config for GCS client")
		}
		return _gcs.NewGcsClient(ctx, gcsConfig)
	case _types.StorageTypeMinio:
		minioConfig, ok := config.(_minio.Config)
		if !ok {
			return nil, fmt.Errorf("invalid config for MinIO client")
		}
		return _minio.NewMinioClient(minioConfig)
	default:
		return nil, fmt.Errorf("unsupported storage type: %s", storageType)
	}
}
