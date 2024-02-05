package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	var (
		s3Client *s3.Client
		err      error
	)
	ctx := context.Background()
	if s3Client, err = initS3Client("us-east-1", ctx); err != nil {
		fmt.Printf("Init sdk failed - %s", err)
		os.Exit(1)
	}

	if err = downloadFile(ctx, s3Client); err != nil {
		fmt.Printf("Download File Error - %s", err)
		os.Exit(1)
	}
	fmt.Println("Download complete")
}

func initS3Client(region string, ctx context.Context) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("unable to load sdk config, %s", err)
	}
	return s3.NewFromConfig(cfg), nil
}

func downloadFile(ctx context.Context, s3Client *s3.Client) error {

	return nil
}
