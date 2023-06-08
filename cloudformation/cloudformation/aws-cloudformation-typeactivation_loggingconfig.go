// Code generated by "go generate". Please don't change this file directly.

package cloudformation

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// TypeActivation_LoggingConfig AWS CloudFormation Resource (AWS::CloudFormation::TypeActivation.LoggingConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-typeactivation-loggingconfig.html
type TypeActivation_LoggingConfig[T any] struct {

	// LogGroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-typeactivation-loggingconfig.html#cfn-cloudformation-typeactivation-loggingconfig-loggroupname
	LogGroupName *string `json:"LogGroupName,omitempty"`

	// LogRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-typeactivation-loggingconfig.html#cfn-cloudformation-typeactivation-loggingconfig-logrolearn
	LogRoleArn *string `json:"LogRoleArn,omitempty"`

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
func (r *TypeActivation_LoggingConfig[any]) AWSCloudFormationType() string {
	return "AWS::CloudFormation::TypeActivation.LoggingConfig"
}
