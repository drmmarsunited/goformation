package utils

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

type AwsHelper struct {
	Cfg aws.Config
}

func NewAwsHelper() (*AwsHelper, error) {
	// Load the AWS shared config
	awsConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("NewAwsHelper: error trying to load AWS SDK config: %w", err)
	}

	return &AwsHelper{
		Cfg: awsConfig,
	}, nil
}
