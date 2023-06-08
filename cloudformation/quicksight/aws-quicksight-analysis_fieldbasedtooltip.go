// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_FieldBasedTooltip AWS CloudFormation Resource (AWS::QuickSight::Analysis.FieldBasedTooltip)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-fieldbasedtooltip.html
type Analysis_FieldBasedTooltip[T any] struct {

	// AggregationVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-fieldbasedtooltip.html#cfn-quicksight-analysis-fieldbasedtooltip-aggregationvisibility
	AggregationVisibility *string `json:"AggregationVisibility,omitempty"`

	// TooltipFields AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-fieldbasedtooltip.html#cfn-quicksight-analysis-fieldbasedtooltip-tooltipfields
	TooltipFields []Analysis_TooltipItem[any] `json:"TooltipFields,omitempty"`

	// TooltipTitleType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-fieldbasedtooltip.html#cfn-quicksight-analysis-fieldbasedtooltip-tooltiptitletype
	TooltipTitleType *string `json:"TooltipTitleType,omitempty"`

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
func (r *Analysis_FieldBasedTooltip[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.FieldBasedTooltip"
}
