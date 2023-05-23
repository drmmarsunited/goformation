// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_BarChartConfiguration AWS CloudFormation Resource (AWS::QuickSight::Template.BarChartConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartconfiguration.html
type Template_BarChartConfiguration struct {

	// BarsArrangement AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartconfiguration.html#cfn-quicksight-template-barchartconfiguration-barsarrangement
	BarsArrangement *string `json:"BarsArrangement,omitempty"`

	// CategoryAxis AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartconfiguration.html#cfn-quicksight-template-barchartconfiguration-categoryaxis
	CategoryAxis *Template_AxisDisplayOptions `json:"CategoryAxis,omitempty"`

	// CategoryLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartconfiguration.html#cfn-quicksight-template-barchartconfiguration-categorylabeloptions
	CategoryLabelOptions *Template_ChartAxisLabelOptions `json:"CategoryLabelOptions,omitempty"`

	// ColorLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartconfiguration.html#cfn-quicksight-template-barchartconfiguration-colorlabeloptions
	ColorLabelOptions *Template_ChartAxisLabelOptions `json:"ColorLabelOptions,omitempty"`

	// ContributionAnalysisDefaults AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartconfiguration.html#cfn-quicksight-template-barchartconfiguration-contributionanalysisdefaults
	ContributionAnalysisDefaults []Template_ContributionAnalysisDefault `json:"ContributionAnalysisDefaults,omitempty"`

	// DataLabels AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartconfiguration.html#cfn-quicksight-template-barchartconfiguration-datalabels
	DataLabels *Template_DataLabelOptions `json:"DataLabels,omitempty"`

	// FieldWells AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartconfiguration.html#cfn-quicksight-template-barchartconfiguration-fieldwells
	FieldWells *Template_BarChartFieldWells `json:"FieldWells,omitempty"`

	// Legend AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartconfiguration.html#cfn-quicksight-template-barchartconfiguration-legend
	Legend *Template_LegendOptions `json:"Legend,omitempty"`

	// Orientation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartconfiguration.html#cfn-quicksight-template-barchartconfiguration-orientation
	Orientation *string `json:"Orientation,omitempty"`

	// ReferenceLines AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartconfiguration.html#cfn-quicksight-template-barchartconfiguration-referencelines
	ReferenceLines []Template_ReferenceLine `json:"ReferenceLines,omitempty"`

	// SmallMultiplesOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartconfiguration.html#cfn-quicksight-template-barchartconfiguration-smallmultiplesoptions
	SmallMultiplesOptions *Template_SmallMultiplesOptions `json:"SmallMultiplesOptions,omitempty"`

	// SortConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartconfiguration.html#cfn-quicksight-template-barchartconfiguration-sortconfiguration
	SortConfiguration *Template_BarChartSortConfiguration `json:"SortConfiguration,omitempty"`

	// Tooltip AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartconfiguration.html#cfn-quicksight-template-barchartconfiguration-tooltip
	Tooltip *Template_TooltipOptions `json:"Tooltip,omitempty"`

	// ValueAxis AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartconfiguration.html#cfn-quicksight-template-barchartconfiguration-valueaxis
	ValueAxis *Template_AxisDisplayOptions `json:"ValueAxis,omitempty"`

	// ValueLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartconfiguration.html#cfn-quicksight-template-barchartconfiguration-valuelabeloptions
	ValueLabelOptions *Template_ChartAxisLabelOptions `json:"ValueLabelOptions,omitempty"`

	// VisualPalette AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartconfiguration.html#cfn-quicksight-template-barchartconfiguration-visualpalette
	VisualPalette *Template_VisualPalette `json:"VisualPalette,omitempty"`

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
func (r *Template_BarChartConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.BarChartConfiguration"
}
