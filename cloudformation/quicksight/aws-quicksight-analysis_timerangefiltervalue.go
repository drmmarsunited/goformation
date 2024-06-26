// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_TimeRangeFilterValue AWS CloudFormation Resource (AWS::QuickSight::Analysis.TimeRangeFilterValue)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-timerangefiltervalue.html
type Analysis_TimeRangeFilterValue[T any] struct {

	// Parameter AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-timerangefiltervalue.html#cfn-quicksight-analysis-timerangefiltervalue-parameter
	Parameter *T `json:"Parameter,omitempty"`

	// RollingDate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-timerangefiltervalue.html#cfn-quicksight-analysis-timerangefiltervalue-rollingdate
	RollingDate *Analysis_RollingDateConfiguration[any] `json:"RollingDate,omitempty"`

	// StaticValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-timerangefiltervalue.html#cfn-quicksight-analysis-timerangefiltervalue-staticvalue
	StaticValue *T `json:"StaticValue,omitempty"`

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
func (r *Analysis_TimeRangeFilterValue[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.TimeRangeFilterValue"
}
