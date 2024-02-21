package intrinsics

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/drmmarsunited/goformation/v7/cloudformation/utils"
	"strings"
)

type pseudoParamHelper struct {
	awsHelper        *utils.AwsHelper
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

func newPseudoParamHelper() *pseudoParamHelper {
	interimHelper := &pseudoParamHelper{}

	// Load the AWS shared config
	awsHelper, err := utils.NewAwsHelper()
	if err != nil {
		fmt.Println("Unable to load shared AWS SDK credentials, using dummy values")

		interimHelper.rawCallerData = nil
		interimHelper.accountId = "123456789012"
		interimHelper.notificationArns = []string{"arn:aws:sns:us-east-1:123456789012:MyTopic"}
		interimHelper.noValue = nil
		interimHelper.partition = "aws"
		interimHelper.region = "us-east-1"
		interimHelper.stackId = []string{"arn:aws:cloudformation:us-east-1:123456789012:stack/MyStack/1c2fa620-982a-11e3-aff7-50e2416294e0"}
		interimHelper.stackName = "goformation-stack"

		return interimHelper
	}

	// Set up an STS service client
	svc := sts.NewFromConfig(awsHelper.Cfg)
	interimHelper.awsHelper = awsHelper
	interimHelper.stsSvc = svc

	// Call the GetCallerIdentity STS API
	resp, err := svc.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	if err != nil {
		interimHelper.rawCallerData = nil
		interimHelper.accountId = "123456789012"
		interimHelper.notificationArns = []string{"arn:aws:sns:us-east-1:123456789012:MyTopic"}
		interimHelper.noValue = nil
		interimHelper.partition = "aws"
		interimHelper.region = "us-east-1"
		interimHelper.stackId = []string{"arn:aws:cloudformation:us-east-1:123456789012:stack/MyStack/1c2fa620-982a-11e3-aff7-50e2416294e0"}
		interimHelper.stackName = "goformation-stack"
	}

	interimHelper.rawCallerData = resp

	return interimHelper
}

// getCallerIdentityData will attempt to call the GetCallerIdentity STS API and load the raw response data
// into the pseudoParamHelper struct. If any errors occur, load dummy data into the struct
func (h *pseudoParamHelper) getCallerIdentityData() {
	// Call the GetCallerIdentity STS API
	resp, err := h.stsSvc.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	if err != nil {
		h.rawCallerData = nil
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
		h.region = h.awsHelper.Cfg.Region
	}
}
