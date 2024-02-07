package main

import (
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
		out      []byte
	)
	ctx := context.Background()
	if s3Client, err = initS3Client("us-east-1", ctx); err != nil {
		fmt.Printf("Init sdk failed - %s", err)
		os.Exit(1)
	}

	if out, err = downloadFile(ctx, s3Client); err != nil {
		fmt.Printf("Download File Error - %s", err)
		os.Exit(1)
	}
	fmt.Printf("Download complete %s", out)
}

func initS3Client(region string, ctx context.Context) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("unable to load sdk config, %s", err)
	}
	return s3.NewFromConfig(cfg), nil
}

func downloadFile(ctx context.Context, s3Client *s3.Client) ([]byte, error) {
	download := manager.NewDownloader(s3Client)
	buffer := manager.NewWriteAtBuffer([]byte{})
	numBytes, err := download.Download(ctx, buffer, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("test.txt"),
	})
	if err != nil {
		return nil, fmt.Errorf("download error, %s", err)
	}
	if numBytesReceived := len(buffer.Bytes()); numBytes != int64(numBytesReceived) {
		return nil, fmt.Errorf("numbytes received doents match %d %d", numBytes, numBytesReceived)
	}
	return buffer.Bytes(), nil
}
