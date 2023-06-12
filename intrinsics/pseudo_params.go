package intrinsics

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"strings"
)

var UnableToLoadAwsConfig = errors.New("unable to load shared local AWS config")

type pseudoParamHelper struct {
	cfg              aws.Config
	stsSvc           *sts.Client
	rawCallerData    *sts.GetCallerIdentityOutput
	accountId        string
	notificationArns []string
	noValue          interface{}
	partition        string
	region           string
	stackId          []string
	stackName        string
}

func newPseudoParamHelper() (*pseudoParamHelper, error) {
	// Load the AWS shared config
	awsConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Printf("Unable to load AWS config")
		return nil, UnableToLoadAwsConfig
	}

	// Set up an STS service client
	svc := sts.NewFromConfig(awsConfig)

	return &pseudoParamHelper{
		cfg:    awsConfig,
		stsSvc: svc,
	}, nil
}

// getCallerIdentityData will attempt to call the GetCallerIdentity STS API and load the raw response data
// into the pseudoParamHelper struct. If any errors occur, load dummy data into the struct
func (h *pseudoParamHelper) getCallerIdentityData() {
	// Call the GetCallerIdentity STS API
	resp, err := h.stsSvc.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	if err != nil {
		h.rawCallerData = nil

		h.accountId = "123456789012"
		h.notificationArns = []string{"arn:aws:sns:us-east-1:123456789012:MyTopic"}
		h.noValue = nil
		h.partition = "aws"
		h.region = "us-east-1"
		h.stackId = []string{"arn:aws:cloudformation:us-east-1:123456789012:stack/MyStack/1c2fa620-982a-11e3-aff7-50e2416294e0"}
		h.stackName = "goformation-stack"
	}

	h.rawCallerData = resp
}

// parseAwsAccountId is a helper function to extract the AWS account ID from the raw caller data
func (h *pseudoParamHelper) parseAwsAccountId() {
	if h.rawCallerData != nil {
		h.accountId = *h.rawCallerData.Account
	}
}

// parseAwsPartition is a helper function to extract the AWS partition from the caller ID ARN
func (h *pseudoParamHelper) parseAwsPartition() {
	if h.rawCallerData != nil {
		h.partition = strings.Split(*h.rawCallerData.Arn, ":")[1]
	}
}

// parseAwsRegion is a helper function to extract the AWS region from the caller ID ARN
func (h *pseudoParamHelper) parseAwsRegion() {
	if h.rawCallerData != nil {
		h.region = h.cfg.Region
	}
}
