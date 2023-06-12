package intrinsics

import (
	"context"
	acfn "github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

// FnImportValue is not implemented, and always returns "dummyvalue.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-importvalue.html
func FnImportValue(name string, input interface{}, template interface{}) interface{} {

	// { "Fn::ImportValue" : sharedValueToImport }

	// Set up the AWS helper
	awsHelper, err := newAwsHelper()
	if err != nil {
		return "dummyvalue"
	}

	// Set up CloudFormation client
	cfnClient := acfn.NewFromConfig(awsHelper.cfg)

	// List CloudFormation exports for region
	exports, err := cfnClient.ListExports(context.TODO(), &acfn.ListExportsInput{})
	if err != nil {
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
