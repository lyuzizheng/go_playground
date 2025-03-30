package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GeneratePresignedURL(filename string, filesize int64, contentType string, md5 string) (string, error) {
	svc := GetPresignClient()

	// Validate content type
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/bmp":  true,
	}
	if !allowedTypes[contentType] {
		return "", fmt.Errorf("invalid content type: %s", contentType)
	}

	input := &s3.PutObjectInput{
		Bucket:        aws.String(GetBucketName()),
		Key:           aws.String(filename),
		ContentType:   aws.String(contentType),
		ContentLength: aws.Int64(filesize),
		ContentMD5:    aws.String(md5),
	}

	req, err := svc.PresignPutObject(context.TODO(), input)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %+v, %w", input, err)
	}
	fmt.Println(json.Marshal(req))
	return req.URL, nil
}
