package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func generateURL(bucket string, key string) (string, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return "", fmt.Errorf("failed to create session. %w", err)
	}

	svc := s3.New(sess)
	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	presignedURL, err := req.Presign(15 * time.Minute)
	if err != nil {
		return "", fmt.Errorf("failed to create presigned URL. %w", err)
	}

	return presignedURL, nil
}

func main() {
	var bucket string
	var key string
	flag.StringVar(&bucket, "bucket", "", "S3 bucket to generate URL for")
	flag.StringVar(&key, "key", "", "Key to store the file at in the bucket")
	flag.Parse()

	if bucket == "" {
		log.Fatal("Missing required --bucket parameter.")
	}

	if key == "" {
		log.Fatal("Missing required --key parameter.")
	}

	bucketURL, err := generateURL(bucket, key)
	if err != nil {
		log.Fatalf("failed to generate URL. %s", err)
	}

	log.Printf("Generated URL: %s", bucketURL)
}
