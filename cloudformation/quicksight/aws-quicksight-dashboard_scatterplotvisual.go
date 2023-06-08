// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_ScatterPlotVisual AWS CloudFormation Resource (AWS::QuickSight::Dashboard.ScatterPlotVisual)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotvisual.html
type Dashboard_ScatterPlotVisual[T any] struct {

	// Actions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotvisual.html#cfn-quicksight-dashboard-scatterplotvisual-actions
	Actions []Dashboard_VisualCustomAction[any] `json:"Actions,omitempty"`

	// ChartConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotvisual.html#cfn-quicksight-dashboard-scatterplotvisual-chartconfiguration
	ChartConfiguration *Dashboard_ScatterPlotConfiguration[any] `json:"ChartConfiguration,omitempty"`

	// ColumnHierarchies AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotvisual.html#cfn-quicksight-dashboard-scatterplotvisual-columnhierarchies
	ColumnHierarchies []Dashboard_ColumnHierarchy[any] `json:"ColumnHierarchies,omitempty"`

	// Subtitle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotvisual.html#cfn-quicksight-dashboard-scatterplotvisual-subtitle
	Subtitle *Dashboard_VisualSubtitleLabelOptions[any] `json:"Subtitle,omitempty"`

	// Title AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotvisual.html#cfn-quicksight-dashboard-scatterplotvisual-title
	Title *Dashboard_VisualTitleLabelOptions[any] `json:"Title,omitempty"`

	// VisualId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotvisual.html#cfn-quicksight-dashboard-scatterplotvisual-visualid
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
func (r *Dashboard_ScatterPlotVisual[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.ScatterPlotVisual"
}
