package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	var (
		instanceId string
		err        error
	)
	arg := os.Args

	if len(arg) < 2 {
		fmt.Printf("Use: go run main.go <arg>")
		os.Exit(1)
	}
	ctx := context.Background()
	if instanceId, err = createEC2(arg[1], ctx, "us-east-1"); err != nil {
		fmt.Errorf("CreateEc2 - %s", err)
		os.Exit(1)
	}
	fmt.Printf("Instance id: $s\n", instanceId)
}
func createEC2(bucket string, ctx context.Context, region string) (string, error) {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithSharedConfigProfile("homero"),
		config.WithRegion(region))
	if err != nil {
		return "", fmt.Errorf("Unable to load sdk config, %s", err)
	}
	ec2Client := s3.NewFromConfig(cfg)
	output, err := ec2Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Name bucket: %s\n", *output.Name)
	for _, objetos := range output.Contents {
		fmt.Printf("key: %v\nSize: %d\n", *objetos.Key, *objetos.Size)
	}

	return "", nil
}
