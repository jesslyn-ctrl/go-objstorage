package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"os"
)

// MinioClient represents a Minio client.
type MinioClient struct {
	client *minio.Client
}

// NewMinioClient creates a new Minio client.
func NewMinioClient(config Config) (*MinioClient, error) {
	// Apply default configuration
	config.DefaultConfig()
	// Initialize Minio client
	client, err := minio.New(config.APIUrl, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.SecretKey, ""),
		Secure: *config.UseSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Minio client: %w", err)
	}
	return &MinioClient{client: client}, nil
}

// CreateBucket creates a new bucket in Minio.
func (m *MinioClient) CreateBucket(ctx context.Context, bucketName string) error {
	// Check if the bucket already exists
	exists, err := m.client.BucketExists(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("failed to check bucket existence: %w", err)
	}

	if exists {
		return fmt.Errorf("bucket already exists: %s", bucketName)
	}

	// Create the bucket
	err = m.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		return fmt.Errorf("failed to create bucket: %w", err)
	}

	return nil
}

// PutObject uploads a file to the specified bucket and object name.
func (m *MinioClient) PutObject(ctx context.Context, bucket, object, filePath string) error {
	// Open the file to upload
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Upload the file to Minio
	_, err = m.client.FPutObject(ctx, bucket, object, filePath, minio.PutObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to upload file to MinIO: %w", err)
	}

	return nil
}

// GetObject retrieves an object from the specified bucket and object name.
func (m *MinioClient) GetObject(ctx context.Context, bucket, object string) ([]byte, error) {
	// Retrieve the object from Minio
	obj, err := m.client.GetObject(ctx, bucket, object, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get object from MinIO: %w", err)
	}
	defer obj.Close()

	// Read the object content into a byte slice
	data, err := io.ReadAll(obj)
	if err != nil {
		return nil, fmt.Errorf("failed to read object content: %w", err)
	}

	return data, nil
}
