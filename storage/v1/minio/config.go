package minio

// Config configurations for Minio.
type Config struct {
	// Minio
	APIUrl    string // MinIO API endpoint
	AccessKey string // MinIO access key
	SecretKey string // MinIO secret key
	UseSSL    *bool  // MinIO SSL usage
}

// DefaultConfig ensures defaults are applied.
func (c *Config) DefaultConfig() {
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
