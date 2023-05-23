// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_FunnelChartConfiguration AWS CloudFormation Resource (AWS::QuickSight::Dashboard.FunnelChartConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-funnelchartconfiguration.html
type Dashboard_FunnelChartConfiguration struct {

	// CategoryLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-funnelchartconfiguration.html#cfn-quicksight-dashboard-funnelchartconfiguration-categorylabeloptions
	CategoryLabelOptions *Dashboard_ChartAxisLabelOptions `json:"CategoryLabelOptions,omitempty"`

	// DataLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-funnelchartconfiguration.html#cfn-quicksight-dashboard-funnelchartconfiguration-datalabeloptions
	DataLabelOptions *Dashboard_FunnelChartDataLabelOptions `json:"DataLabelOptions,omitempty"`

	// FieldWells AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-funnelchartconfiguration.html#cfn-quicksight-dashboard-funnelchartconfiguration-fieldwells
	FieldWells *Dashboard_FunnelChartFieldWells `json:"FieldWells,omitempty"`

	// SortConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-funnelchartconfiguration.html#cfn-quicksight-dashboard-funnelchartconfiguration-sortconfiguration
	SortConfiguration *Dashboard_FunnelChartSortConfiguration `json:"SortConfiguration,omitempty"`

	// Tooltip AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-funnelchartconfiguration.html#cfn-quicksight-dashboard-funnelchartconfiguration-tooltip
	Tooltip *Dashboard_TooltipOptions `json:"Tooltip,omitempty"`

	// ValueLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-funnelchartconfiguration.html#cfn-quicksight-dashboard-funnelchartconfiguration-valuelabeloptions
	ValueLabelOptions *Dashboard_ChartAxisLabelOptions `json:"ValueLabelOptions,omitempty"`

	// VisualPalette AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-funnelchartconfiguration.html#cfn-quicksight-dashboard-funnelchartconfiguration-visualpalette
	VisualPalette *Dashboard_VisualPalette `json:"VisualPalette,omitempty"`

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
func (r *Dashboard_FunnelChartConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.FunnelChartConfiguration"
}
