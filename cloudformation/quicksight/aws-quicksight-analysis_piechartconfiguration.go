// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_PieChartConfiguration AWS CloudFormation Resource (AWS::QuickSight::Analysis.PieChartConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-piechartconfiguration.html
type Analysis_PieChartConfiguration[T any] struct {

	// CategoryLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-piechartconfiguration.html#cfn-quicksight-analysis-piechartconfiguration-categorylabeloptions
	CategoryLabelOptions *Analysis_ChartAxisLabelOptions[any] `json:"CategoryLabelOptions,omitempty"`

	// ContributionAnalysisDefaults AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-piechartconfiguration.html#cfn-quicksight-analysis-piechartconfiguration-contributionanalysisdefaults
	ContributionAnalysisDefaults []Analysis_ContributionAnalysisDefault[any] `json:"ContributionAnalysisDefaults,omitempty"`

	// DataLabels AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-piechartconfiguration.html#cfn-quicksight-analysis-piechartconfiguration-datalabels
	DataLabels *Analysis_DataLabelOptions[any] `json:"DataLabels,omitempty"`

	// DonutOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-piechartconfiguration.html#cfn-quicksight-analysis-piechartconfiguration-donutoptions
	DonutOptions *Analysis_DonutOptions[any] `json:"DonutOptions,omitempty"`

	// FieldWells AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-piechartconfiguration.html#cfn-quicksight-analysis-piechartconfiguration-fieldwells
	FieldWells *Analysis_PieChartFieldWells[any] `json:"FieldWells,omitempty"`

	// Legend AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-piechartconfiguration.html#cfn-quicksight-analysis-piechartconfiguration-legend
	Legend *Analysis_LegendOptions[any] `json:"Legend,omitempty"`

	// SmallMultiplesOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-piechartconfiguration.html#cfn-quicksight-analysis-piechartconfiguration-smallmultiplesoptions
	SmallMultiplesOptions *Analysis_SmallMultiplesOptions[any] `json:"SmallMultiplesOptions,omitempty"`

	// SortConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-piechartconfiguration.html#cfn-quicksight-analysis-piechartconfiguration-sortconfiguration
	SortConfiguration *Analysis_PieChartSortConfiguration[any] `json:"SortConfiguration,omitempty"`

	// Tooltip AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-piechartconfiguration.html#cfn-quicksight-analysis-piechartconfiguration-tooltip
	Tooltip *Analysis_TooltipOptions[any] `json:"Tooltip,omitempty"`

	// ValueLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-piechartconfiguration.html#cfn-quicksight-analysis-piechartconfiguration-valuelabeloptions
	ValueLabelOptions *Analysis_ChartAxisLabelOptions[any] `json:"ValueLabelOptions,omitempty"`

	// VisualPalette AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-piechartconfiguration.html#cfn-quicksight-analysis-piechartconfiguration-visualpalette
	VisualPalette *Analysis_VisualPalette[any] `json:"VisualPalette,omitempty"`

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
func (r *Analysis_PieChartConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.PieChartConfiguration"
}
