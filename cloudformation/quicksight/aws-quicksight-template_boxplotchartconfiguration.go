// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_BoxPlotChartConfiguration AWS CloudFormation Resource (AWS::QuickSight::Template.BoxPlotChartConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-boxplotchartconfiguration.html
type Template_BoxPlotChartConfiguration[T any] struct {

	// BoxPlotOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-boxplotchartconfiguration.html#cfn-quicksight-template-boxplotchartconfiguration-boxplotoptions
	BoxPlotOptions *Template_BoxPlotOptions[any] `json:"BoxPlotOptions,omitempty"`

	// CategoryAxis AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-boxplotchartconfiguration.html#cfn-quicksight-template-boxplotchartconfiguration-categoryaxis
	CategoryAxis *Template_AxisDisplayOptions[any] `json:"CategoryAxis,omitempty"`

	// CategoryLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-boxplotchartconfiguration.html#cfn-quicksight-template-boxplotchartconfiguration-categorylabeloptions
	CategoryLabelOptions *Template_ChartAxisLabelOptions[any] `json:"CategoryLabelOptions,omitempty"`

	// FieldWells AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-boxplotchartconfiguration.html#cfn-quicksight-template-boxplotchartconfiguration-fieldwells
	FieldWells *Template_BoxPlotFieldWells[any] `json:"FieldWells,omitempty"`

	// Legend AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-boxplotchartconfiguration.html#cfn-quicksight-template-boxplotchartconfiguration-legend
	Legend *Template_LegendOptions[any] `json:"Legend,omitempty"`

	// PrimaryYAxisDisplayOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-boxplotchartconfiguration.html#cfn-quicksight-template-boxplotchartconfiguration-primaryyaxisdisplayoptions
	PrimaryYAxisDisplayOptions *Template_AxisDisplayOptions[any] `json:"PrimaryYAxisDisplayOptions,omitempty"`

	// PrimaryYAxisLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-boxplotchartconfiguration.html#cfn-quicksight-template-boxplotchartconfiguration-primaryyaxislabeloptions
	PrimaryYAxisLabelOptions *Template_ChartAxisLabelOptions[any] `json:"PrimaryYAxisLabelOptions,omitempty"`

	// ReferenceLines AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-boxplotchartconfiguration.html#cfn-quicksight-template-boxplotchartconfiguration-referencelines
	ReferenceLines []Template_ReferenceLine[any] `json:"ReferenceLines,omitempty"`

	// SortConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-boxplotchartconfiguration.html#cfn-quicksight-template-boxplotchartconfiguration-sortconfiguration
	SortConfiguration *Template_BoxPlotSortConfiguration[any] `json:"SortConfiguration,omitempty"`

	// Tooltip AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-boxplotchartconfiguration.html#cfn-quicksight-template-boxplotchartconfiguration-tooltip
	Tooltip *Template_TooltipOptions[any] `json:"Tooltip,omitempty"`

	// VisualPalette AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-boxplotchartconfiguration.html#cfn-quicksight-template-boxplotchartconfiguration-visualpalette
	VisualPalette *Template_VisualPalette[any] `json:"VisualPalette,omitempty"`

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
func (r *Template_BoxPlotChartConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.BoxPlotChartConfiguration"
}
