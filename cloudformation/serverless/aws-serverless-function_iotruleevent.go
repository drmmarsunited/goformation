// Code generated by "go generate". Please don't change this file directly.

package serverless

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Function_IoTRuleEvent AWS CloudFormation Resource (AWS::Serverless::Function.IoTRuleEvent)
// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#iotrule
type Function_IoTRuleEvent[T any] struct {

	// AwsIotSqlVersion AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#iotrule
	AwsIotSqlVersion *T `json:"AwsIotSqlVersion,omitempty"`

	// Sql AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#iotrule
	Sql T `json:"Sql"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Function_IoTRuleEvent[any]) AWSCloudFormationType() string {
	return "AWS::Serverless::Function.IoTRuleEvent"
}
