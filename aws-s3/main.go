package main

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const bucketName = "create-bucket-with-go-324f"

func main() {
	var (
		s3Client *s3.Client
		err      error
	)

	ctx := context.Background()
	if s3Client, err = initS3Client("us-east-1", ctx); err != nil {
		fmt.Printf("S3 error - %s", err)
		os.Exit(1)
	}
	if err = createBucket(ctx, s3Client); err != nil {
		fmt.Printf("S3 error - %s", err)
		os.Exit(1)
	}
	fmt.Printf("init upload file...\n")
	if err = uploadFile(ctx, s3Client); err != nil {
		fmt.Printf("Upload file - %s", err)
		os.Exit(1)
	}
	fmt.Println("upload complete")
}
func initS3Client(region string, ctx context.Context) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("unable to load sdk config, %s", err)
	}
	return s3.NewFromConfig(cfg), nil
}

func createBucket(ctx context.Context, s3Client *s3.Client) error {
	allBuckets, err := s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return fmt.Errorf("list bucket -->  %s", err)
	}
	bu := allBuckets.Buckets
	found := false
	for _, bucket := range allBuckets.Buckets {
		if *bucket.Name == bucketName {
			fmt.Printf("bucket existing -> %v\n", *bu[0].Name)
			found = true
		}
	}
	if !found {
		_, err := s3Client.CreateBucket(ctx, &s3.CreateBucketInput{
			Bucket: aws.String(bucketName),
		})
		if err != nil {
			return fmt.Errorf("error create bucket -->  %s", err)
		}
		fmt.Printf("Bucket created success: %s", bucketName)
	}
	return nil
}
func uploadFile(ctx context.Context, s3Client *s3.Client) error {
	file, err := os.ReadFile("test.txt")
	if err != nil {
		return fmt.Errorf("upload -->  %s", err)
	}
	upload := manager.NewUploader(s3Client)
	_, err = upload.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("test.txt"),
		Body:   bytes.NewReader(file),
	})
	if err != nil {
		return fmt.Errorf("upload -->  %s", err)
	}
	return nil
}
