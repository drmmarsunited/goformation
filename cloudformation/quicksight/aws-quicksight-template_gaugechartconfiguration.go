// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_GaugeChartConfiguration AWS CloudFormation Resource (AWS::QuickSight::Template.GaugeChartConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gaugechartconfiguration.html
type Template_GaugeChartConfiguration[T any] struct {

	// DataLabels AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gaugechartconfiguration.html#cfn-quicksight-template-gaugechartconfiguration-datalabels
	DataLabels *Template_DataLabelOptions[any] `json:"DataLabels,omitempty"`

	// FieldWells AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gaugechartconfiguration.html#cfn-quicksight-template-gaugechartconfiguration-fieldwells
	FieldWells *Template_GaugeChartFieldWells[any] `json:"FieldWells,omitempty"`

	// GaugeChartOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gaugechartconfiguration.html#cfn-quicksight-template-gaugechartconfiguration-gaugechartoptions
	GaugeChartOptions *Template_GaugeChartOptions[any] `json:"GaugeChartOptions,omitempty"`

	// TooltipOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gaugechartconfiguration.html#cfn-quicksight-template-gaugechartconfiguration-tooltipoptions
	TooltipOptions *Template_TooltipOptions[any] `json:"TooltipOptions,omitempty"`

	// VisualPalette AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gaugechartconfiguration.html#cfn-quicksight-template-gaugechartconfiguration-visualpalette
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
func (r *Template_GaugeChartConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.GaugeChartConfiguration"
}
