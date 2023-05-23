// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_HistogramConfiguration AWS CloudFormation Resource (AWS::QuickSight::Analysis.HistogramConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-histogramconfiguration.html
type Analysis_HistogramConfiguration struct {

	// BinOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-histogramconfiguration.html#cfn-quicksight-analysis-histogramconfiguration-binoptions
	BinOptions *Analysis_HistogramBinOptions `json:"BinOptions,omitempty"`

	// DataLabels AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-histogramconfiguration.html#cfn-quicksight-analysis-histogramconfiguration-datalabels
	DataLabels *Analysis_DataLabelOptions `json:"DataLabels,omitempty"`

	// FieldWells AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-histogramconfiguration.html#cfn-quicksight-analysis-histogramconfiguration-fieldwells
	FieldWells *Analysis_HistogramFieldWells `json:"FieldWells,omitempty"`

	// Tooltip AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-histogramconfiguration.html#cfn-quicksight-analysis-histogramconfiguration-tooltip
	Tooltip *Analysis_TooltipOptions `json:"Tooltip,omitempty"`

	// VisualPalette AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-histogramconfiguration.html#cfn-quicksight-analysis-histogramconfiguration-visualpalette
	VisualPalette *Analysis_VisualPalette `json:"VisualPalette,omitempty"`

	// XAxisDisplayOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-histogramconfiguration.html#cfn-quicksight-analysis-histogramconfiguration-xaxisdisplayoptions
	XAxisDisplayOptions *Analysis_AxisDisplayOptions `json:"XAxisDisplayOptions,omitempty"`

	// XAxisLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-histogramconfiguration.html#cfn-quicksight-analysis-histogramconfiguration-xaxislabeloptions
	XAxisLabelOptions *Analysis_ChartAxisLabelOptions `json:"XAxisLabelOptions,omitempty"`

	// YAxisDisplayOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-histogramconfiguration.html#cfn-quicksight-analysis-histogramconfiguration-yaxisdisplayoptions
	YAxisDisplayOptions *Analysis_AxisDisplayOptions `json:"YAxisDisplayOptions,omitempty"`

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
func (r *Analysis_HistogramConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.HistogramConfiguration"
}
