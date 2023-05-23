// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_SankeyDiagramVisual AWS CloudFormation Resource (AWS::QuickSight::Template.SankeyDiagramVisual)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sankeydiagramvisual.html
type Template_SankeyDiagramVisual struct {

	// Actions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sankeydiagramvisual.html#cfn-quicksight-template-sankeydiagramvisual-actions
	Actions []Template_VisualCustomAction `json:"Actions,omitempty"`

	// ChartConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sankeydiagramvisual.html#cfn-quicksight-template-sankeydiagramvisual-chartconfiguration
	ChartConfiguration *Template_SankeyDiagramChartConfiguration `json:"ChartConfiguration,omitempty"`

	// Subtitle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sankeydiagramvisual.html#cfn-quicksight-template-sankeydiagramvisual-subtitle
	Subtitle *Template_VisualSubtitleLabelOptions `json:"Subtitle,omitempty"`

	// Title AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sankeydiagramvisual.html#cfn-quicksight-template-sankeydiagramvisual-title
	Title *Template_VisualTitleLabelOptions `json:"Title,omitempty"`

	// VisualId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sankeydiagramvisual.html#cfn-quicksight-template-sankeydiagramvisual-visualid
	VisualId string `json:"VisualId"`

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
func (r *Template_SankeyDiagramVisual) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.SankeyDiagramVisual"
}
