// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_FormatConfiguration AWS CloudFormation Resource (AWS::QuickSight::Template.FormatConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-formatconfiguration.html
type Template_FormatConfiguration[T any] struct {

	// DateTimeFormatConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-formatconfiguration.html#cfn-quicksight-template-formatconfiguration-datetimeformatconfiguration
	DateTimeFormatConfiguration *Template_DateTimeFormatConfiguration[any] `json:"DateTimeFormatConfiguration,omitempty"`

	// NumberFormatConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-formatconfiguration.html#cfn-quicksight-template-formatconfiguration-numberformatconfiguration
	NumberFormatConfiguration *Template_NumberFormatConfiguration[any] `json:"NumberFormatConfiguration,omitempty"`

	// StringFormatConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-formatconfiguration.html#cfn-quicksight-template-formatconfiguration-stringformatconfiguration
	StringFormatConfiguration *Template_StringFormatConfiguration[any] `json:"StringFormatConfiguration,omitempty"`

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
func (r *Template_FormatConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.FormatConfiguration"
}
