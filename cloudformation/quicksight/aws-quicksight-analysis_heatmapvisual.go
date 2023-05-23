// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_HeatMapVisual AWS CloudFormation Resource (AWS::QuickSight::Analysis.HeatMapVisual)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-heatmapvisual.html
type Analysis_HeatMapVisual struct {

	// Actions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-heatmapvisual.html#cfn-quicksight-analysis-heatmapvisual-actions
	Actions []Analysis_VisualCustomAction `json:"Actions,omitempty"`

	// ChartConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-heatmapvisual.html#cfn-quicksight-analysis-heatmapvisual-chartconfiguration
	ChartConfiguration *Analysis_HeatMapConfiguration `json:"ChartConfiguration,omitempty"`

	// ColumnHierarchies AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-heatmapvisual.html#cfn-quicksight-analysis-heatmapvisual-columnhierarchies
	ColumnHierarchies []Analysis_ColumnHierarchy `json:"ColumnHierarchies,omitempty"`

	// Subtitle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-heatmapvisual.html#cfn-quicksight-analysis-heatmapvisual-subtitle
	Subtitle *Analysis_VisualSubtitleLabelOptions `json:"Subtitle,omitempty"`

	// Title AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-heatmapvisual.html#cfn-quicksight-analysis-heatmapvisual-title
	Title *Analysis_VisualTitleLabelOptions `json:"Title,omitempty"`

	// VisualId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-heatmapvisual.html#cfn-quicksight-analysis-heatmapvisual-visualid
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
func (r *Analysis_HeatMapVisual) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.HeatMapVisual"
}
