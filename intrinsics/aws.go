package intrinsics

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var UnableToLoadAwsConfig = errors.New("unable to load shared local AWS config")

type awsHelper struct {
	cfg aws.Config
}

func newAwsHelper() (*awsHelper, error) {
	// Load the AWS shared config
	awsConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Printf("Unable to load AWS config")
		return nil, UnableToLoadAwsConfig
	}

	return &awsHelper{
		cfg: awsConfig,
	}, nil
}
