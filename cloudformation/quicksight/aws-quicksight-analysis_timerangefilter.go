// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_TimeRangeFilter AWS CloudFormation Resource (AWS::QuickSight::Analysis.TimeRangeFilter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-timerangefilter.html
type Analysis_TimeRangeFilter[T any] struct {

	// Column AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-timerangefilter.html#cfn-quicksight-analysis-timerangefilter-column
	Column *Analysis_ColumnIdentifier[any] `json:"Column"`

	// ExcludePeriodConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-timerangefilter.html#cfn-quicksight-analysis-timerangefilter-excludeperiodconfiguration
	ExcludePeriodConfiguration *Analysis_ExcludePeriodConfiguration[any] `json:"ExcludePeriodConfiguration,omitempty"`

	// FilterId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-timerangefilter.html#cfn-quicksight-analysis-timerangefilter-filterid
	FilterId T `json:"FilterId"`

	// IncludeMaximum AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-timerangefilter.html#cfn-quicksight-analysis-timerangefilter-includemaximum
	IncludeMaximum *T `json:"IncludeMaximum,omitempty"`

	// IncludeMinimum AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-timerangefilter.html#cfn-quicksight-analysis-timerangefilter-includeminimum
	IncludeMinimum *T `json:"IncludeMinimum,omitempty"`

	// NullOption AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-timerangefilter.html#cfn-quicksight-analysis-timerangefilter-nulloption
	NullOption T `json:"NullOption"`

	// RangeMaximumValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-timerangefilter.html#cfn-quicksight-analysis-timerangefilter-rangemaximumvalue
	RangeMaximumValue *Analysis_TimeRangeFilterValue[any] `json:"RangeMaximumValue,omitempty"`

	// RangeMinimumValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-timerangefilter.html#cfn-quicksight-analysis-timerangefilter-rangeminimumvalue
	RangeMinimumValue *Analysis_TimeRangeFilterValue[any] `json:"RangeMinimumValue,omitempty"`

	// TimeGranularity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-timerangefilter.html#cfn-quicksight-analysis-timerangefilter-timegranularity
	TimeGranularity *T `json:"TimeGranularity,omitempty"`

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
func (r *Analysis_TimeRangeFilter[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.TimeRangeFilter"
}
