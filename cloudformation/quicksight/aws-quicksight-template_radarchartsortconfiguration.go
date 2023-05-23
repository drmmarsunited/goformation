// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_RadarChartSortConfiguration AWS CloudFormation Resource (AWS::QuickSight::Template.RadarChartSortConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-radarchartsortconfiguration.html
type Template_RadarChartSortConfiguration struct {

	// CategoryItemsLimit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-radarchartsortconfiguration.html#cfn-quicksight-template-radarchartsortconfiguration-categoryitemslimit
	CategoryItemsLimit *Template_ItemsLimitConfiguration `json:"CategoryItemsLimit,omitempty"`

	// CategorySort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-radarchartsortconfiguration.html#cfn-quicksight-template-radarchartsortconfiguration-categorysort
	CategorySort []Template_FieldSortOptions `json:"CategorySort,omitempty"`

	// ColorItemsLimit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-radarchartsortconfiguration.html#cfn-quicksight-template-radarchartsortconfiguration-coloritemslimit
	ColorItemsLimit *Template_ItemsLimitConfiguration `json:"ColorItemsLimit,omitempty"`

	// ColorSort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-radarchartsortconfiguration.html#cfn-quicksight-template-radarchartsortconfiguration-colorsort
	ColorSort []Template_FieldSortOptions `json:"ColorSort,omitempty"`

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
func (r *Template_RadarChartSortConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.RadarChartSortConfiguration"
}
