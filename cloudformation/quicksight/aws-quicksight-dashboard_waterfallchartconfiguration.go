// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_WaterfallChartConfiguration AWS CloudFormation Resource (AWS::QuickSight::Dashboard.WaterfallChartConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-waterfallchartconfiguration.html
type Dashboard_WaterfallChartConfiguration struct {

	// CategoryAxisDisplayOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-waterfallchartconfiguration.html#cfn-quicksight-dashboard-waterfallchartconfiguration-categoryaxisdisplayoptions
	CategoryAxisDisplayOptions *Dashboard_AxisDisplayOptions `json:"CategoryAxisDisplayOptions,omitempty"`

	// CategoryAxisLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-waterfallchartconfiguration.html#cfn-quicksight-dashboard-waterfallchartconfiguration-categoryaxislabeloptions
	CategoryAxisLabelOptions *Dashboard_ChartAxisLabelOptions `json:"CategoryAxisLabelOptions,omitempty"`

	// DataLabels AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-waterfallchartconfiguration.html#cfn-quicksight-dashboard-waterfallchartconfiguration-datalabels
	DataLabels *Dashboard_DataLabelOptions `json:"DataLabels,omitempty"`

	// FieldWells AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-waterfallchartconfiguration.html#cfn-quicksight-dashboard-waterfallchartconfiguration-fieldwells
	FieldWells *Dashboard_WaterfallChartFieldWells `json:"FieldWells,omitempty"`

	// Legend AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-waterfallchartconfiguration.html#cfn-quicksight-dashboard-waterfallchartconfiguration-legend
	Legend *Dashboard_LegendOptions `json:"Legend,omitempty"`

	// PrimaryYAxisDisplayOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-waterfallchartconfiguration.html#cfn-quicksight-dashboard-waterfallchartconfiguration-primaryyaxisdisplayoptions
	PrimaryYAxisDisplayOptions *Dashboard_AxisDisplayOptions `json:"PrimaryYAxisDisplayOptions,omitempty"`

	// PrimaryYAxisLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-waterfallchartconfiguration.html#cfn-quicksight-dashboard-waterfallchartconfiguration-primaryyaxislabeloptions
	PrimaryYAxisLabelOptions *Dashboard_ChartAxisLabelOptions `json:"PrimaryYAxisLabelOptions,omitempty"`

	// SortConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-waterfallchartconfiguration.html#cfn-quicksight-dashboard-waterfallchartconfiguration-sortconfiguration
	SortConfiguration *Dashboard_WaterfallChartSortConfiguration `json:"SortConfiguration,omitempty"`

	// VisualPalette AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-waterfallchartconfiguration.html#cfn-quicksight-dashboard-waterfallchartconfiguration-visualpalette
	VisualPalette *Dashboard_VisualPalette `json:"VisualPalette,omitempty"`

	// WaterfallChartOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-waterfallchartconfiguration.html#cfn-quicksight-dashboard-waterfallchartconfiguration-waterfallchartoptions
	WaterfallChartOptions *Dashboard_WaterfallChartOptions `json:"WaterfallChartOptions,omitempty"`

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
func (r *Dashboard_WaterfallChartConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.WaterfallChartConfiguration"
}
