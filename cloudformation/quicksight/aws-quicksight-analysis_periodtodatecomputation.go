// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_PeriodToDateComputation AWS CloudFormation Resource (AWS::QuickSight::Analysis.PeriodToDateComputation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-periodtodatecomputation.html
type Analysis_PeriodToDateComputation struct {

	// ComputationId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-periodtodatecomputation.html#cfn-quicksight-analysis-periodtodatecomputation-computationid
	ComputationId string `json:"ComputationId"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-periodtodatecomputation.html#cfn-quicksight-analysis-periodtodatecomputation-name
	Name *string `json:"Name,omitempty"`

	// PeriodTimeGranularity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-periodtodatecomputation.html#cfn-quicksight-analysis-periodtodatecomputation-periodtimegranularity
	PeriodTimeGranularity *string `json:"PeriodTimeGranularity,omitempty"`

	// Time AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-periodtodatecomputation.html#cfn-quicksight-analysis-periodtodatecomputation-time
	Time *Analysis_DimensionField `json:"Time"`

	// Value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-periodtodatecomputation.html#cfn-quicksight-analysis-periodtodatecomputation-value
	Value *Analysis_MeasureField `json:"Value,omitempty"`

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
func (r *Analysis_PeriodToDateComputation) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.PeriodToDateComputation"
}
