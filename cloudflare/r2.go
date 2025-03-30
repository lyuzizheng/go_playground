package cloudflare

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Config holds the configuration for R2
type Config struct {
	AccessKey  string
	SecretKey  string
	AccountID  string
	BucketName string
}

// R2Client is a wrapper for the R2 session and client
type R2Client struct {
	client        *s3.Client
	presignClient *s3.PresignClient
	bucketName    string
}

var (
	once     sync.Once
	r2Client *R2Client
)

// Initialize sets up the R2 session and client
func Initialize(cfg Config) error {
	var initError error
	once.Do(func() {
		if err := validateConfig(cfg); err != nil {
			initError = err
			log.Println(initError)
			return
		}

		awsCfg, err := config.LoadDefaultConfig(context.TODO(),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.AccessKey, cfg.SecretKey, "")),
			config.WithRegion("auto"),
		)
		if err != nil {
			initError = fmt.Errorf("failed to load AWS config: %w", err)
			log.Println(initError)
			return
		}

		client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
			o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", cfg.AccountID))
		})

		r2Client = &R2Client{
			client:        client,
			presignClient: s3.NewPresignClient(client),
			bucketName:    cfg.BucketName,
		}
	})
	return initError
}

// validateConfig checks if the provided configuration is valid
func validateConfig(cfg Config) error {
	if cfg.AccessKey == "" || cfg.SecretKey == "" || cfg.AccountID == "" || cfg.BucketName == "" {
		return fmt.Errorf("invalid configuration: all fields must be non-empty")
	}
	return nil
}

// GetClient returns the global R2 client
func GetClient() *s3.Client {
	if r2Client == nil {
		log.Println("R2 client is not initialized")
		return nil
	}
	return r2Client.client
}

// GetPresignClient returns the global R2 presign client
func GetPresignClient() *s3.PresignClient {
	if r2Client == nil {
		log.Println("R2 presign client is not initialized")
		return nil
	}
	return r2Client.presignClient
}

// GetBucketName returns the bucket name
func GetBucketName() string {
	if r2Client == nil {
		log.Println("R2 client is not initialized")
		return ""
	}
	return r2Client.bucketName
}

// PresignPutObject generates a presigned URL for a PutObject request
func (c *R2Client) PresignPutObject(key string) (string, error) {
	presignResult, err := c.presignClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(c.bucketName),
		Key:    aws.String(key),
	})

	if err != nil {
		return "", fmt.Errorf("couldn't get presigned URL for PutObject: %w", err)
	}

	return presignResult.URL, nil
}
