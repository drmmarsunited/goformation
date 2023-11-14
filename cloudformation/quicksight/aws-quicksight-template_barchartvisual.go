// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_BarChartVisual AWS CloudFormation Resource (AWS::QuickSight::Template.BarChartVisual)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartvisual.html
type Template_BarChartVisual[T any] struct {

	// Actions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartvisual.html#cfn-quicksight-template-barchartvisual-actions
	Actions []Template_VisualCustomAction[any] `json:"Actions,omitempty"`

	// ChartConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartvisual.html#cfn-quicksight-template-barchartvisual-chartconfiguration
	ChartConfiguration *Template_BarChartConfiguration[any] `json:"ChartConfiguration,omitempty"`

	// ColumnHierarchies AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartvisual.html#cfn-quicksight-template-barchartvisual-columnhierarchies
	ColumnHierarchies []Template_ColumnHierarchy[any] `json:"ColumnHierarchies,omitempty"`

	// Subtitle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartvisual.html#cfn-quicksight-template-barchartvisual-subtitle
	Subtitle *Template_VisualSubtitleLabelOptions[any] `json:"Subtitle,omitempty"`

	// Title AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartvisual.html#cfn-quicksight-template-barchartvisual-title
	Title *Template_VisualTitleLabelOptions[any] `json:"Title,omitempty"`

	// VisualId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartvisual.html#cfn-quicksight-template-barchartvisual-visualid
	VisualId T `json:"VisualId"`

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
func (r *Template_BarChartVisual[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.BarChartVisual"
}
