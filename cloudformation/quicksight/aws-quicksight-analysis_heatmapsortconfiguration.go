// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_HeatMapSortConfiguration AWS CloudFormation Resource (AWS::QuickSight::Analysis.HeatMapSortConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-heatmapsortconfiguration.html
type Analysis_HeatMapSortConfiguration struct {

	// HeatMapColumnItemsLimitConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-heatmapsortconfiguration.html#cfn-quicksight-analysis-heatmapsortconfiguration-heatmapcolumnitemslimitconfiguration
	HeatMapColumnItemsLimitConfiguration *Analysis_ItemsLimitConfiguration `json:"HeatMapColumnItemsLimitConfiguration,omitempty"`

	// HeatMapColumnSort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-heatmapsortconfiguration.html#cfn-quicksight-analysis-heatmapsortconfiguration-heatmapcolumnsort
	HeatMapColumnSort []Analysis_FieldSortOptions `json:"HeatMapColumnSort,omitempty"`

	// HeatMapRowItemsLimitConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-heatmapsortconfiguration.html#cfn-quicksight-analysis-heatmapsortconfiguration-heatmaprowitemslimitconfiguration
	HeatMapRowItemsLimitConfiguration *Analysis_ItemsLimitConfiguration `json:"HeatMapRowItemsLimitConfiguration,omitempty"`

	// HeatMapRowSort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-heatmapsortconfiguration.html#cfn-quicksight-analysis-heatmapsortconfiguration-heatmaprowsort
	HeatMapRowSort []Analysis_FieldSortOptions `json:"HeatMapRowSort,omitempty"`

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
func (r *Analysis_HeatMapSortConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.HeatMapSortConfiguration"
}
