// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_MetricComparisonComputation AWS CloudFormation Resource (AWS::QuickSight::Template.MetricComparisonComputation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-metriccomparisoncomputation.html
type Template_MetricComparisonComputation[T any] struct {

	// ComputationId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-metriccomparisoncomputation.html#cfn-quicksight-template-metriccomparisoncomputation-computationid
	ComputationId string `json:"ComputationId"`

	// FromValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-metriccomparisoncomputation.html#cfn-quicksight-template-metriccomparisoncomputation-fromvalue
	FromValue *Template_MeasureField[any] `json:"FromValue"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-metriccomparisoncomputation.html#cfn-quicksight-template-metriccomparisoncomputation-name
	Name *string `json:"Name,omitempty"`

	// TargetValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-metriccomparisoncomputation.html#cfn-quicksight-template-metriccomparisoncomputation-targetvalue
	TargetValue *Template_MeasureField[any] `json:"TargetValue"`

	// Time AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-metriccomparisoncomputation.html#cfn-quicksight-template-metriccomparisoncomputation-time
	Time *Template_DimensionField[any] `json:"Time"`

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
func (r *Template_MetricComparisonComputation[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.MetricComparisonComputation"
}
