// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_ColumnSort AWS CloudFormation Resource (AWS::QuickSight::Dashboard.ColumnSort)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-columnsort.html
type Dashboard_ColumnSort[T any] struct {

	// AggregationFunction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-columnsort.html#cfn-quicksight-dashboard-columnsort-aggregationfunction
	AggregationFunction *Dashboard_AggregationFunction[any] `json:"AggregationFunction,omitempty"`

	// Direction AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-columnsort.html#cfn-quicksight-dashboard-columnsort-direction
	Direction T `json:"Direction"`

	// SortBy AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-columnsort.html#cfn-quicksight-dashboard-columnsort-sortby
	SortBy *Dashboard_ColumnIdentifier[any] `json:"SortBy"`

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
func (r *Dashboard_ColumnSort[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.ColumnSort"
}
