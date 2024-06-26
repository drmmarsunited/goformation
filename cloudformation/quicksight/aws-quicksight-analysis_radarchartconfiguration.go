// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_RadarChartConfiguration AWS CloudFormation Resource (AWS::QuickSight::Analysis.RadarChartConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html
type Analysis_RadarChartConfiguration[T any] struct {

	// AlternateBandColorsVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-alternatebandcolorsvisibility
	AlternateBandColorsVisibility *T `json:"AlternateBandColorsVisibility,omitempty"`

	// AlternateBandEvenColor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-alternatebandevencolor
	AlternateBandEvenColor *T `json:"AlternateBandEvenColor,omitempty"`

	// AlternateBandOddColor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-alternatebandoddcolor
	AlternateBandOddColor *T `json:"AlternateBandOddColor,omitempty"`

	// AxesRangeScale AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-axesrangescale
	AxesRangeScale *T `json:"AxesRangeScale,omitempty"`

	// BaseSeriesSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-baseseriessettings
	BaseSeriesSettings *Analysis_RadarChartSeriesSettings[any] `json:"BaseSeriesSettings,omitempty"`

	// CategoryAxis AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-categoryaxis
	CategoryAxis *Analysis_AxisDisplayOptions[any] `json:"CategoryAxis,omitempty"`

	// CategoryLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-categorylabeloptions
	CategoryLabelOptions *Analysis_ChartAxisLabelOptions[any] `json:"CategoryLabelOptions,omitempty"`

	// ColorAxis AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-coloraxis
	ColorAxis *Analysis_AxisDisplayOptions[any] `json:"ColorAxis,omitempty"`

	// ColorLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-colorlabeloptions
	ColorLabelOptions *Analysis_ChartAxisLabelOptions[any] `json:"ColorLabelOptions,omitempty"`

	// FieldWells AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-fieldwells
	FieldWells *Analysis_RadarChartFieldWells[any] `json:"FieldWells,omitempty"`

	// Legend AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-legend
	Legend *Analysis_LegendOptions[any] `json:"Legend,omitempty"`

	// Shape AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-shape
	Shape *T `json:"Shape,omitempty"`

	// SortConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-sortconfiguration
	SortConfiguration *Analysis_RadarChartSortConfiguration[any] `json:"SortConfiguration,omitempty"`

	// StartAngle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-startangle
	StartAngle *T `json:"StartAngle,omitempty"`

	// VisualPalette AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-visualpalette
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
func (r *Analysis_RadarChartConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.RadarChartConfiguration"
}
