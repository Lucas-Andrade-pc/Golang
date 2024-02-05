package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func main() {
	var (
		instance string
		err      error
		region   string = "us-east-1"
	)
	ctx := context.TODO()
	if instance, err = createEC2(region, ctx); err != nil {
		fmt.Printf("Create error: %s", err)
	} else {
		fmt.Printf("Instance: %s", instance)
	}
}

func createEC2(region string, ctx context.Context) (string, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return "", fmt.Errorf("erro loading config -> %s", err)
	}

	ec2Client := ec2.NewFromConfig(cfg)
	// describeKey, err := ec2Client.DescribeKeyPairs(ctx, &ec2.DescribeKeyPairsInput{
	// 	Filters: []types.Filter{
	// 		{
	// 			Name:   aws.String("key-name"),
	// 			Values: []string{"golang"},
	// 		},
	// 	},
	// })
	// if err != nil {
	// 	return "", fmt.Errorf("key pair -> %v", err)

	// }
	// if len(describeKey.KeyPairs) != 0 {
	// 	return "", fmt.Errorf("KeyPair already exist ")
	// }
	// _, err = ec2Client.CreateKeyPair(ctx, &ec2.CreateKeyPairInput{
	// 	KeyName: aws.String("golang"),
	// })
	// if err != nil {
	// 	return "", fmt.Errorf("create keyPair error -> %s", err)
	// }
	imageOutput, err := ec2Client.DescribeImages(ctx, &ec2.DescribeImagesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("name"),
				Values: []string{"ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-*"},
			},
			{
				Name:   aws.String("virtualization-type"),
				Values: []string{"hvm"},
			},
		},
		Owners: []string{"099720109477"},
	})
	if err != nil {
		return "", fmt.Errorf("describe image error -> %v", err)

	}
	if len(imageOutput.Images) == 0 {
		return "", fmt.Errorf("size slice error -> %v", err)
	}
	instance, err := ec2Client.RunInstances(ctx, &ec2.RunInstancesInput{
		ImageId:      imageOutput.Images[0].ImageId,
		KeyName:      aws.String("golang"),
		InstanceType: types.InstanceTypeT2Micro,
		MinCount:     aws.Int32(1),
		MaxCount:     aws.Int32(1),
	})
	if err != nil {
		return "", fmt.Errorf("erro format create instance -> %s", err)
	}
	if len(instance.Instances) == 0 {
		return "", fmt.Errorf("erro instance -> %s", err)
	}
	return *instance.Instances[0].InstanceId, nil
}
