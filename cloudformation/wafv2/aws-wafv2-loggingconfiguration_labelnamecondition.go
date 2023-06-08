// Code generated by "go generate". Please don't change this file directly.

package wafv2

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// LoggingConfiguration_LabelNameCondition AWS CloudFormation Resource (AWS::WAFv2::LoggingConfiguration.LabelNameCondition)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-loggingconfiguration-labelnamecondition.html
type LoggingConfiguration_LabelNameCondition[T any] struct {

	// LabelName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-loggingconfiguration-labelnamecondition.html#cfn-wafv2-loggingconfiguration-labelnamecondition-labelname
	LabelName string `json:"LabelName"`

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
func (r *LoggingConfiguration_LabelNameCondition[any]) AWSCloudFormationType() string {
	return "AWS::WAFv2::LoggingConfiguration.LabelNameCondition"
}
