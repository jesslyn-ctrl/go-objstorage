package types

type Config struct {
	// GCS
	ProjectID       string // GCS project ID
	CredentialsFile string // GCS path to the service account credentials file

	// Minio
	APIUrl    string // MinIO API endpoint
	AccessKey string // MinIO access key
	SecretKey string // MinIO secret key
	UseSSL    *bool  // MinIO SSL usage
}

// MinioDefaultConfig ensures defaults are applied.
func (c *Config) MinioDefaultConfig() {
	apiUrl := c.APIUrl
	useSSL := c.UseSSL
	if apiUrl == "" {
		apiUrl = "localhost:9000"
	}
	if useSSL == nil {
		defaultSSL := true
		useSSL = &defaultSSL
	}
}
