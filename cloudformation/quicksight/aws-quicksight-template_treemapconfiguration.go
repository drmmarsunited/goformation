// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_TreeMapConfiguration AWS CloudFormation Resource (AWS::QuickSight::Template.TreeMapConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-treemapconfiguration.html
type Template_TreeMapConfiguration[T any] struct {

	// ColorLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-treemapconfiguration.html#cfn-quicksight-template-treemapconfiguration-colorlabeloptions
	ColorLabelOptions *Template_ChartAxisLabelOptions[any] `json:"ColorLabelOptions,omitempty"`

	// ColorScale AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-treemapconfiguration.html#cfn-quicksight-template-treemapconfiguration-colorscale
	ColorScale *Template_ColorScale[any] `json:"ColorScale,omitempty"`

	// DataLabels AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-treemapconfiguration.html#cfn-quicksight-template-treemapconfiguration-datalabels
	DataLabels *Template_DataLabelOptions[any] `json:"DataLabels,omitempty"`

	// FieldWells AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-treemapconfiguration.html#cfn-quicksight-template-treemapconfiguration-fieldwells
	FieldWells *Template_TreeMapFieldWells[any] `json:"FieldWells,omitempty"`

	// GroupLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-treemapconfiguration.html#cfn-quicksight-template-treemapconfiguration-grouplabeloptions
	GroupLabelOptions *Template_ChartAxisLabelOptions[any] `json:"GroupLabelOptions,omitempty"`

	// Legend AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-treemapconfiguration.html#cfn-quicksight-template-treemapconfiguration-legend
	Legend *Template_LegendOptions[any] `json:"Legend,omitempty"`

	// SizeLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-treemapconfiguration.html#cfn-quicksight-template-treemapconfiguration-sizelabeloptions
	SizeLabelOptions *Template_ChartAxisLabelOptions[any] `json:"SizeLabelOptions,omitempty"`

	// SortConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-treemapconfiguration.html#cfn-quicksight-template-treemapconfiguration-sortconfiguration
	SortConfiguration *Template_TreeMapSortConfiguration[any] `json:"SortConfiguration,omitempty"`

	// Tooltip AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-treemapconfiguration.html#cfn-quicksight-template-treemapconfiguration-tooltip
	Tooltip *Template_TooltipOptions[any] `json:"Tooltip,omitempty"`

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
func (r *Template_TreeMapConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.TreeMapConfiguration"
}
