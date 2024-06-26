// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_LineChartConfiguration AWS CloudFormation Resource (AWS::QuickSight::Dashboard.LineChartConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html
type Dashboard_LineChartConfiguration[T any] struct {

	// ContributionAnalysisDefaults AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-contributionanalysisdefaults
	ContributionAnalysisDefaults []Dashboard_ContributionAnalysisDefault[any] `json:"ContributionAnalysisDefaults,omitempty"`

	// DataLabels AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-datalabels
	DataLabels *Dashboard_DataLabelOptions[any] `json:"DataLabels,omitempty"`

	// DefaultSeriesSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-defaultseriessettings
	DefaultSeriesSettings *Dashboard_LineChartDefaultSeriesSettings[any] `json:"DefaultSeriesSettings,omitempty"`

	// FieldWells AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-fieldwells
	FieldWells *Dashboard_LineChartFieldWells[any] `json:"FieldWells,omitempty"`

	// ForecastConfigurations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-forecastconfigurations
	ForecastConfigurations []Dashboard_ForecastConfiguration[any] `json:"ForecastConfigurations,omitempty"`

	// Legend AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-legend
	Legend *Dashboard_LegendOptions[any] `json:"Legend,omitempty"`

	// PrimaryYAxisDisplayOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-primaryyaxisdisplayoptions
	PrimaryYAxisDisplayOptions *Dashboard_LineSeriesAxisDisplayOptions[any] `json:"PrimaryYAxisDisplayOptions,omitempty"`

	// PrimaryYAxisLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-primaryyaxislabeloptions
	PrimaryYAxisLabelOptions *Dashboard_ChartAxisLabelOptions[any] `json:"PrimaryYAxisLabelOptions,omitempty"`

	// ReferenceLines AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-referencelines
	ReferenceLines []Dashboard_ReferenceLine[any] `json:"ReferenceLines,omitempty"`

	// SecondaryYAxisDisplayOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-secondaryyaxisdisplayoptions
	SecondaryYAxisDisplayOptions *Dashboard_LineSeriesAxisDisplayOptions[any] `json:"SecondaryYAxisDisplayOptions,omitempty"`

	// SecondaryYAxisLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-secondaryyaxislabeloptions
	SecondaryYAxisLabelOptions *Dashboard_ChartAxisLabelOptions[any] `json:"SecondaryYAxisLabelOptions,omitempty"`

	// Series AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-series
	Series []Dashboard_SeriesItem[any] `json:"Series,omitempty"`

	// SmallMultiplesOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-smallmultiplesoptions
	SmallMultiplesOptions *Dashboard_SmallMultiplesOptions[any] `json:"SmallMultiplesOptions,omitempty"`

	// SortConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-sortconfiguration
	SortConfiguration *Dashboard_LineChartSortConfiguration[any] `json:"SortConfiguration,omitempty"`

	// Tooltip AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-tooltip
	Tooltip *Dashboard_TooltipOptions[any] `json:"Tooltip,omitempty"`

	// Type AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-type
	Type *T `json:"Type,omitempty"`

	// VisualPalette AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-visualpalette
	VisualPalette *Dashboard_VisualPalette[any] `json:"VisualPalette,omitempty"`

	// XAxisDisplayOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-xaxisdisplayoptions
	XAxisDisplayOptions *Dashboard_AxisDisplayOptions[any] `json:"XAxisDisplayOptions,omitempty"`

	// XAxisLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-xaxislabeloptions
	XAxisLabelOptions *Dashboard_ChartAxisLabelOptions[any] `json:"XAxisLabelOptions,omitempty"`

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
func (r *Dashboard_LineChartConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.LineChartConfiguration"
}
