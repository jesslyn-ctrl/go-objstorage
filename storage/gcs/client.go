package gcs

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"google.golang.org/api/option"
	"io"
	"os"
)

// GcsClient represents a Google Cloud Storage client.
type GcsClient struct {
	client    *storage.Client
	projectID string
}

// NewGcsClient creates a new GCS client.
func NewGcsClient(ctx context.Context, config Config) (*GcsClient, error) {
	// If CredentialsFile is empty, try to use the default credentials (via environment variables)
	var client *storage.Client
	var err error

	if config.CredentialsFile != "" {
		// If credentials file is provided, use it
		// Using GRPC Client transporter
		client, err = storage.NewGRPCClient(ctx, option.WithCredentialsFile(config.CredentialsFile))
	} else {
		// Otherwise, use default application credentials
		client, err = storage.NewGRPCClient(ctx)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create GCS client: %w", err)
	}
	defer client.Close()

	return &GcsClient{client: client, projectID: config.ProjectID}, nil
}

// CreateBucket creates a new bucket in the specified project.
func (c *GcsClient) CreateBucket(ctx context.Context, bucketName string) error {
	bucket := c.client.Bucket(bucketName)
	if err := bucket.Create(ctx, c.projectID, nil); err != nil {
		return fmt.Errorf("failed to create bucket: %w", err)
	}
	return nil
}

// PutObject uploads a file to the specified bucket and object name.
func (c *GcsClient) PutObject(ctx context.Context, bucketName, objectName, filePath string) error {
	// Open the file to upload
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Get a reference to the bucket and object
	bucket := c.client.Bucket(bucketName)
	object := bucket.Object(objectName)

	// Create a writer to upload the object
	writer := object.NewWriter(ctx)
	defer writer.Close()

	// Copy the file content to the object
	if _, err := io.Copy(writer, file); err != nil {
		return fmt.Errorf("failed to upload file to GCS: %w", err)
	}

	return nil
}

// GetObject retrieves an object from the specified bucket and object name.
func (c *GcsClient) GetObject(ctx context.Context, bucketName, objectName string) ([]byte, error) {
	// Get a reference to the bucket and object
	bucket := c.client.Bucket(bucketName)
	object := bucket.Object(objectName)

	// Create a reader to download the object
	reader, err := object.NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to read object from GCS: %w", err)
	}
	defer reader.Close()

	// Read the object content into a byte slice
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read object content: %w", err)
	}

	return data, nil
}
