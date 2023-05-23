// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_TimeRangeDrillDownFilter AWS CloudFormation Resource (AWS::QuickSight::Analysis.TimeRangeDrillDownFilter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-timerangedrilldownfilter.html
type Analysis_TimeRangeDrillDownFilter struct {

	// Column AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-timerangedrilldownfilter.html#cfn-quicksight-analysis-timerangedrilldownfilter-column
	Column *Analysis_ColumnIdentifier `json:"Column"`

	// RangeMaximum AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-timerangedrilldownfilter.html#cfn-quicksight-analysis-timerangedrilldownfilter-rangemaximum
	RangeMaximum string `json:"RangeMaximum"`

	// RangeMinimum AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-timerangedrilldownfilter.html#cfn-quicksight-analysis-timerangedrilldownfilter-rangeminimum
	RangeMinimum string `json:"RangeMinimum"`

	// TimeGranularity AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-timerangedrilldownfilter.html#cfn-quicksight-analysis-timerangedrilldownfilter-timegranularity
	TimeGranularity string `json:"TimeGranularity"`

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
func (r *Analysis_TimeRangeDrillDownFilter) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.TimeRangeDrillDownFilter"
}
