// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_RadarChartConfiguration AWS CloudFormation Resource (AWS::QuickSight::Dashboard.RadarChartConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-radarchartconfiguration.html
type Dashboard_RadarChartConfiguration[T any] struct {

	// AlternateBandColorsVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-radarchartconfiguration.html#cfn-quicksight-dashboard-radarchartconfiguration-alternatebandcolorsvisibility
	AlternateBandColorsVisibility *T `json:"AlternateBandColorsVisibility,omitempty"`

	// AlternateBandEvenColor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-radarchartconfiguration.html#cfn-quicksight-dashboard-radarchartconfiguration-alternatebandevencolor
	AlternateBandEvenColor *T `json:"AlternateBandEvenColor,omitempty"`

	// AlternateBandOddColor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-radarchartconfiguration.html#cfn-quicksight-dashboard-radarchartconfiguration-alternatebandoddcolor
	AlternateBandOddColor *T `json:"AlternateBandOddColor,omitempty"`

	// AxesRangeScale AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-radarchartconfiguration.html#cfn-quicksight-dashboard-radarchartconfiguration-axesrangescale
	AxesRangeScale *T `json:"AxesRangeScale,omitempty"`

	// BaseSeriesSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-radarchartconfiguration.html#cfn-quicksight-dashboard-radarchartconfiguration-baseseriessettings
	BaseSeriesSettings *Dashboard_RadarChartSeriesSettings[any] `json:"BaseSeriesSettings,omitempty"`

	// CategoryAxis AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-radarchartconfiguration.html#cfn-quicksight-dashboard-radarchartconfiguration-categoryaxis
	CategoryAxis *Dashboard_AxisDisplayOptions[any] `json:"CategoryAxis,omitempty"`

	// CategoryLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-radarchartconfiguration.html#cfn-quicksight-dashboard-radarchartconfiguration-categorylabeloptions
	CategoryLabelOptions *Dashboard_ChartAxisLabelOptions[any] `json:"CategoryLabelOptions,omitempty"`

	// ColorAxis AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-radarchartconfiguration.html#cfn-quicksight-dashboard-radarchartconfiguration-coloraxis
	ColorAxis *Dashboard_AxisDisplayOptions[any] `json:"ColorAxis,omitempty"`

	// ColorLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-radarchartconfiguration.html#cfn-quicksight-dashboard-radarchartconfiguration-colorlabeloptions
	ColorLabelOptions *Dashboard_ChartAxisLabelOptions[any] `json:"ColorLabelOptions,omitempty"`

	// FieldWells AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-radarchartconfiguration.html#cfn-quicksight-dashboard-radarchartconfiguration-fieldwells
	FieldWells *Dashboard_RadarChartFieldWells[any] `json:"FieldWells,omitempty"`

	// Legend AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-radarchartconfiguration.html#cfn-quicksight-dashboard-radarchartconfiguration-legend
	Legend *Dashboard_LegendOptions[any] `json:"Legend,omitempty"`

	// Shape AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-radarchartconfiguration.html#cfn-quicksight-dashboard-radarchartconfiguration-shape
	Shape *T `json:"Shape,omitempty"`

	// SortConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-radarchartconfiguration.html#cfn-quicksight-dashboard-radarchartconfiguration-sortconfiguration
	SortConfiguration *Dashboard_RadarChartSortConfiguration[any] `json:"SortConfiguration,omitempty"`

	// StartAngle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-radarchartconfiguration.html#cfn-quicksight-dashboard-radarchartconfiguration-startangle
	StartAngle *T `json:"StartAngle,omitempty"`

	// VisualPalette AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-radarchartconfiguration.html#cfn-quicksight-dashboard-radarchartconfiguration-visualpalette
	VisualPalette *Dashboard_VisualPalette[any] `json:"VisualPalette,omitempty"`

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
func (r *Dashboard_RadarChartConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.RadarChartConfiguration"
}
