package intrinsics

import (
	"context"
	"fmt"
	acfn "github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/drmmarsunited/goformation/v7/cloudformation/utils"
)

// FnImportValue is now implemented and will return a value if found, and returns "dummyvalue if not found.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-importvalue.html
func FnImportValue(name string, input interface{}, template interface{}) interface{} {

	// { "Fn::ImportValue" : sharedValueToImport }

	// Set up the AWS helper
	awsHelper, err := utils.NewAwsHelper()
	if err != nil {
		return "dummyvalue"
	}

	// Set up CloudFormation client
	cfnClient := acfn.NewFromConfig(awsHelper.Cfg)

	// List CloudFormation exports for region
	exports, err := cfnClient.ListExports(context.TODO(), nil)
	if err != nil {
		fmt.Printf("Could not get list of CFN exports: %s\n", err.Error())
		return "dummyvalue"
	}

	// Parse through exports for matching name
	if name, ok := input.(string); ok {
		for _, e := range exports.Exports {
			if *e.Name == name {
				return *e.Value
			}
		}
	}

	// Return dummy value if no matches were found
	return "dummyvalue"
}
