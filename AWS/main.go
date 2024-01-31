package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	var (
		instanceId string
		err        error
	)
	ctx := context.Background()
	if instanceId, err = createEC2(ctx, "us-west2"); err != nil {
		fmt.Errorf("CreateEc2 - %s", err)
		os.Exit(1)
	}
	fmt.Printf("Instance id: $s\n", instanceId)
}
func createEC2(ctx context.Context, region string) (string, error) {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithSharedConfigProfile("homero"),
		config.WithRegion(region))
	if err != nil {
		return "", fmt.Errorf("Unable to load sdk config, %s", err)
	}
	ec2Client := ec2.NewFromConfig(cfg)

	filter := ec2Client.DescribeInstances(cfg)

	return "", nil
}
