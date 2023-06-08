// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_WaterfallChartSortConfiguration AWS CloudFormation Resource (AWS::QuickSight::Analysis.WaterfallChartSortConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-waterfallchartsortconfiguration.html
type Analysis_WaterfallChartSortConfiguration[T any] struct {

	// BreakdownItemsLimit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-waterfallchartsortconfiguration.html#cfn-quicksight-analysis-waterfallchartsortconfiguration-breakdownitemslimit
	BreakdownItemsLimit *Analysis_ItemsLimitConfiguration[any] `json:"BreakdownItemsLimit,omitempty"`

	// CategorySort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-waterfallchartsortconfiguration.html#cfn-quicksight-analysis-waterfallchartsortconfiguration-categorysort
	CategorySort []Analysis_FieldSortOptions[any] `json:"CategorySort,omitempty"`

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
func (r *Analysis_WaterfallChartSortConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.WaterfallChartSortConfiguration"
}
