// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_TotalAggregationOption AWS CloudFormation Resource (AWS::QuickSight::Dashboard.TotalAggregationOption)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-totalaggregationoption.html
type Dashboard_TotalAggregationOption[T any] struct {

	// FieldId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-totalaggregationoption.html#cfn-quicksight-dashboard-totalaggregationoption-fieldid
	FieldId T `json:"FieldId"`

	// TotalAggregationFunction AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-totalaggregationoption.html#cfn-quicksight-dashboard-totalaggregationoption-totalaggregationfunction
	TotalAggregationFunction *Dashboard_TotalAggregationFunction[any] `json:"TotalAggregationFunction"`

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
func (r *Dashboard_TotalAggregationOption[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.TotalAggregationOption"
}