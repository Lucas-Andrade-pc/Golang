package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go/aws"
)

const region = "us-east-1"

func main() {
	var (
		err       error
		ec2Client *ec2.Client
	)
	ctx := context.Background()
	if ec2Client, err = initSDK(ctx); err != nil {
		fmt.Errorf("Init - %s", err)
		os.Exit(1)
	}
	if err = describeEc2(ctx, ec2Client); err != nil {
		fmt.Errorf("CreateEc2 - %s", err)
		os.Exit(1)
	}
}
func initSDK(ctx context.Context) (*ec2.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithSharedConfigProfile("homero"),
		config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("unable to load sdk config, %s", err)
	}
	return ec2.NewFromConfig(cfg), nil
}
func describeEc2(ctx context.Context, ec2Client *ec2.Client) error {
	result, err := ec2Client.DescribeInstances(ctx, &ec2.DescribeInstancesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("tag:Schedule"),
				Values: []string{"false"},
			},
			{
				Name:   aws.String("instance-state-name"),
				Values: []string{"stopped"},
			},
		},

		//InstanceIds: []string{"i-06d6e14bac477c470"},
	})
	if err != nil {
		return fmt.Errorf("erro filter, %s", err)
	}

	output := result.Reservations
	for _, instance := range output {
		for _, instance2 := range instance.Instances {
			// ec2Client.StopInstances(ctx, &ec2.StopInstancesInput{
			// 	InstanceIds: []string{*instance2.InstanceId},
			// })
			for _, tag := range instance2.Tags {
				aux := instance2
				fmt.Printf("instance: %v\ntag: %v\nvalue: %v\n", *aux.PrivateIpAddress, []string{*tag.Key}, []string{*tag.Value})
			}
		}
	}

	return nil
}
