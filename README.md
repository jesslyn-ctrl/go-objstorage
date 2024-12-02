# Go EDTS Object Storage Library
This library provides a unified client for connecting to different storage systems, interacting with: 
- Google Cloud Storage (GCS)
- MinIO

It abstracts away the specific details of each storage system, allowing you to use a common interface to interact with either system.

## Features

---
- Supports both **Google Cloud Storage (GCS)** and **MinIO**.
- Uses a common `StorageClient` interface for CRUD operations.
- Allows flexible configuration using maps or structs.

## Supported Operations

---
- **Create Bucket:** Create a new storage bucket.
- **Put Object:** Upload a file (object) to the storage.
- **Get Object:** Download a file (object) from the storage.

## Prerequisites

---
- For **GCS:** Google Cloud service account credentials (.json file) and project ID.
- For **MinIO:** Access credentials (Access Key and Secret Key) and MinIO API URL.

## Installation

---
```
go get github.com/yourrepo/go-edts-object-storage
```

## Usage Guide

---
### Step 1: Initialize the Storage Client
The client initialization is abstracted by the `NewStorageClient` function, which accepts a `storageType` and a configuration map or struct for the respective storage system.

**Example: Google Cloud Storage (GCS)**
```
func main() {
	ctx := context.Background()

	// GCS configuration
	gcsConfig := map[string]interface{}{
		"ProjectID":       "your-gcp-project-id",
		"CredentialsFile": "/path/to/credentials.json", // Service account file
	}

	// Initialize GCS client
	client, err := NewStorageClient(ctx, _types.StorageTypeGCS, gcsConfig)
	if err != nil {
		log.Fatalf("Error initializing GCS client: %v", err)
	}
	fmt.Println("GCS Client initialized:", client)

	// Use client to interact with GCS (example: create bucket)
	err = client.CreateBucket(ctx, "your-bucket-name")
	if err != nil {
		log.Fatalf("Failed to create bucket: %v", err)
	}
	fmt.Println("Bucket created successfully!")
}
```

**Example: MinIO**
```
func main() {
	ctx := context.Background()

	// MinIO configuration
	minioConfig := map[string]interface{}{
		"APIUrl":    "https://minio.example.com",
		"AccessKey": "your-access-key",
		"SecretKey": "your-secret-key",
		"UseSSL":    true,
	}

	// Initialize MinIO client
	client, err := NewStorageClient(ctx, _types.StorageTypeMinio, minioConfig)
	if err != nil {
		log.Fatalf("Error initializing MinIO client: %v", err)
	}
	fmt.Println("MinIO Client initialized:", client)

	// Use client to interact with MinIO (example: put/get object)
	err = client.PutObject(ctx, "your-bucket-name", "your-object-name", "/path/to/file")
	if err != nil {
		log.Fatalf("Failed to upload object: %v", err)
	}
	fmt.Println("Object uploaded successfully!")
}
```

### Step 2: Define the Configuration Structs
You can define configurations for GCS and MinIO as structs or maps.

**GCS Configuration (`GcsConfig`)**
```aiexclude
type GcsConfig struct {
	ProjectID       string `json:"project_id"`
	CredentialsFile string `json:"credentials_file"`
}
```

**MinIO Configuration (`MinioConfig`)**
```aiexclude
type MinioConfig struct {
	APIUrl    string // MinIO API endpoint
	AccessKey string // MinIO access key
	SecretKey string // MinIO secret key
	UseSSL    *bool  // MinIO SSL usage
}
```

### Step 3: Use StorageClient Functions
After initializing the client using `NewStorageClient`, you can use the `StorageClient` methods to interact with the storage system.

#### Create a Bucket
```aiexclude
err := client.CreateBucket(ctx, "your-bucket-name")
if err != nil {
	log.Fatalf("failed to create bucket: %v", err)
}
```

#### Put an Object (Upload File)
```aiexclude
err := client.PutObject(ctx, "your-bucket-name", "your-object-name", "/path/to/file")
if err != nil {
	log.Fatalf("failed to upload object: %v", err)
}
```

#### Get an Object (Download File)
```aiexclude
data, err := client.GetObject(ctx, "your-bucket-name", "your-object-name")
if err != nil {
	log.Fatalf("failed to get object: %v", err)
}
fmt.Printf("downloaded object data: %s\n", string(data))
```

### Step 4: Customizing SSL or Credentials
- For **MinIO**, you can specify whether to use SSL by setting the `UseSSL` field in the configuration map (`true` or `false`).
- For **GCS**, you can either use default application credentials or specify a custom credentials file by providing the path to your service account's `.json` file.

## Configuration Details

---
#### GCS Config
- **ProjectID:** The GCP project ID.
- **CredentialsFile:** Path to the service account credentials file.

#### MinIO Config
- **APIUrl:** The endpoint URL for MinIO (e.g., https://minio.example.com).
- **AccessKey:** Your MinIO access key.
- **SecretKey:** Your MinIO secret key.
- **UseSSL:** A boolean flag (true or false) indicating whether to use SSL.

## Error Handling

---
Make sure to handle errors gracefully when interacting with the client, especially when dealing with network operations and storage-related tasks like uploading and downloading objects.
