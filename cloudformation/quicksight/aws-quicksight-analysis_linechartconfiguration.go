// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_LineChartConfiguration AWS CloudFormation Resource (AWS::QuickSight::Analysis.LineChartConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html
type Analysis_LineChartConfiguration[T any] struct {

	// ContributionAnalysisDefaults AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-contributionanalysisdefaults
	ContributionAnalysisDefaults []Analysis_ContributionAnalysisDefault[any] `json:"ContributionAnalysisDefaults,omitempty"`

	// DataLabels AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-datalabels
	DataLabels *Analysis_DataLabelOptions[any] `json:"DataLabels,omitempty"`

	// DefaultSeriesSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-defaultseriessettings
	DefaultSeriesSettings *Analysis_LineChartDefaultSeriesSettings[any] `json:"DefaultSeriesSettings,omitempty"`

	// FieldWells AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-fieldwells
	FieldWells *Analysis_LineChartFieldWells[any] `json:"FieldWells,omitempty"`

	// ForecastConfigurations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-forecastconfigurations
	ForecastConfigurations []Analysis_ForecastConfiguration[any] `json:"ForecastConfigurations,omitempty"`

	// Legend AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-legend
	Legend *Analysis_LegendOptions[any] `json:"Legend,omitempty"`

	// PrimaryYAxisDisplayOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-primaryyaxisdisplayoptions
	PrimaryYAxisDisplayOptions *Analysis_LineSeriesAxisDisplayOptions[any] `json:"PrimaryYAxisDisplayOptions,omitempty"`

	// PrimaryYAxisLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-primaryyaxislabeloptions
	PrimaryYAxisLabelOptions *Analysis_ChartAxisLabelOptions[any] `json:"PrimaryYAxisLabelOptions,omitempty"`

	// ReferenceLines AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-referencelines
	ReferenceLines []Analysis_ReferenceLine[any] `json:"ReferenceLines,omitempty"`

	// SecondaryYAxisDisplayOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-secondaryyaxisdisplayoptions
	SecondaryYAxisDisplayOptions *Analysis_LineSeriesAxisDisplayOptions[any] `json:"SecondaryYAxisDisplayOptions,omitempty"`

	// SecondaryYAxisLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-secondaryyaxislabeloptions
	SecondaryYAxisLabelOptions *Analysis_ChartAxisLabelOptions[any] `json:"SecondaryYAxisLabelOptions,omitempty"`

	// Series AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-series
	Series []Analysis_SeriesItem[any] `json:"Series,omitempty"`

	// SmallMultiplesOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-smallmultiplesoptions
	SmallMultiplesOptions *Analysis_SmallMultiplesOptions[any] `json:"SmallMultiplesOptions,omitempty"`

	// SortConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-sortconfiguration
	SortConfiguration *Analysis_LineChartSortConfiguration[any] `json:"SortConfiguration,omitempty"`

	// Tooltip AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-tooltip
	Tooltip *Analysis_TooltipOptions[any] `json:"Tooltip,omitempty"`

	// Type AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-type
	Type *T `json:"Type,omitempty"`

	// VisualPalette AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-visualpalette
	VisualPalette *Analysis_VisualPalette[any] `json:"VisualPalette,omitempty"`

	// XAxisDisplayOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-xaxisdisplayoptions
	XAxisDisplayOptions *Analysis_AxisDisplayOptions[any] `json:"XAxisDisplayOptions,omitempty"`

	// XAxisLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-xaxislabeloptions
	XAxisLabelOptions *Analysis_ChartAxisLabelOptions[any] `json:"XAxisLabelOptions,omitempty"`

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
func (r *Analysis_LineChartConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.LineChartConfiguration"
}
