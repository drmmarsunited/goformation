// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_LineChartConfiguration AWS CloudFormation Resource (AWS::QuickSight::Template.LineChartConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html
type Template_LineChartConfiguration struct {

	// ContributionAnalysisDefaults AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-contributionanalysisdefaults
	ContributionAnalysisDefaults []Template_ContributionAnalysisDefault `json:"ContributionAnalysisDefaults,omitempty"`

	// DataLabels AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-datalabels
	DataLabels *Template_DataLabelOptions `json:"DataLabels,omitempty"`

	// DefaultSeriesSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-defaultseriessettings
	DefaultSeriesSettings *Template_LineChartDefaultSeriesSettings `json:"DefaultSeriesSettings,omitempty"`

	// FieldWells AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-fieldwells
	FieldWells *Template_LineChartFieldWells `json:"FieldWells,omitempty"`

	// ForecastConfigurations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-forecastconfigurations
	ForecastConfigurations []Template_ForecastConfiguration `json:"ForecastConfigurations,omitempty"`

	// Legend AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-legend
	Legend *Template_LegendOptions `json:"Legend,omitempty"`

	// PrimaryYAxisDisplayOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-primaryyaxisdisplayoptions
	PrimaryYAxisDisplayOptions *Template_LineSeriesAxisDisplayOptions `json:"PrimaryYAxisDisplayOptions,omitempty"`

	// PrimaryYAxisLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-primaryyaxislabeloptions
	PrimaryYAxisLabelOptions *Template_ChartAxisLabelOptions `json:"PrimaryYAxisLabelOptions,omitempty"`

	// ReferenceLines AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-referencelines
	ReferenceLines []Template_ReferenceLine `json:"ReferenceLines,omitempty"`

	// SecondaryYAxisDisplayOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-secondaryyaxisdisplayoptions
	SecondaryYAxisDisplayOptions *Template_LineSeriesAxisDisplayOptions `json:"SecondaryYAxisDisplayOptions,omitempty"`

	// SecondaryYAxisLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-secondaryyaxislabeloptions
	SecondaryYAxisLabelOptions *Template_ChartAxisLabelOptions `json:"SecondaryYAxisLabelOptions,omitempty"`

	// Series AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-series
	Series []Template_SeriesItem `json:"Series,omitempty"`

	// SmallMultiplesOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-smallmultiplesoptions
	SmallMultiplesOptions *Template_SmallMultiplesOptions `json:"SmallMultiplesOptions,omitempty"`

	// SortConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-sortconfiguration
	SortConfiguration *Template_LineChartSortConfiguration `json:"SortConfiguration,omitempty"`

	// Tooltip AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-tooltip
	Tooltip *Template_TooltipOptions `json:"Tooltip,omitempty"`

	// Type AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-type
	Type *string `json:"Type,omitempty"`

	// VisualPalette AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-visualpalette
	VisualPalette *Template_VisualPalette `json:"VisualPalette,omitempty"`

	// XAxisDisplayOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-xaxisdisplayoptions
	XAxisDisplayOptions *Template_AxisDisplayOptions `json:"XAxisDisplayOptions,omitempty"`

	// XAxisLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-xaxislabeloptions
	XAxisLabelOptions *Template_ChartAxisLabelOptions `json:"XAxisLabelOptions,omitempty"`

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
func (r *Template_LineChartConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.LineChartConfiguration"
}
