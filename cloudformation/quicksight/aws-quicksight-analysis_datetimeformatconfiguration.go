// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_DateTimeFormatConfiguration AWS CloudFormation Resource (AWS::QuickSight::Analysis.DateTimeFormatConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datetimeformatconfiguration.html
type Analysis_DateTimeFormatConfiguration[T any] struct {

	// DateTimeFormat AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datetimeformatconfiguration.html#cfn-quicksight-analysis-datetimeformatconfiguration-datetimeformat
	DateTimeFormat *T `json:"DateTimeFormat,omitempty"`

	// NullValueFormatConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datetimeformatconfiguration.html#cfn-quicksight-analysis-datetimeformatconfiguration-nullvalueformatconfiguration
	NullValueFormatConfiguration *Analysis_NullValueFormatConfiguration[any] `json:"NullValueFormatConfiguration,omitempty"`

	// NumericFormatConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datetimeformatconfiguration.html#cfn-quicksight-analysis-datetimeformatconfiguration-numericformatconfiguration
	NumericFormatConfiguration *Analysis_NumericFormatConfiguration[any] `json:"NumericFormatConfiguration,omitempty"`

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
func (r *Analysis_DateTimeFormatConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.DateTimeFormatConfiguration"
}
