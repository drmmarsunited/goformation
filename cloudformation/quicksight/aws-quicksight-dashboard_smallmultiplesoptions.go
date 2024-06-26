// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_SmallMultiplesOptions AWS CloudFormation Resource (AWS::QuickSight::Dashboard.SmallMultiplesOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-smallmultiplesoptions.html
type Dashboard_SmallMultiplesOptions[T any] struct {

	// MaxVisibleColumns AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-smallmultiplesoptions.html#cfn-quicksight-dashboard-smallmultiplesoptions-maxvisiblecolumns
	MaxVisibleColumns *T `json:"MaxVisibleColumns,omitempty"`

	// MaxVisibleRows AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-smallmultiplesoptions.html#cfn-quicksight-dashboard-smallmultiplesoptions-maxvisiblerows
	MaxVisibleRows *T `json:"MaxVisibleRows,omitempty"`

	// PanelConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-smallmultiplesoptions.html#cfn-quicksight-dashboard-smallmultiplesoptions-panelconfiguration
	PanelConfiguration *Dashboard_PanelConfiguration[any] `json:"PanelConfiguration,omitempty"`

	// XAxis AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-smallmultiplesoptions.html#cfn-quicksight-dashboard-smallmultiplesoptions-xaxis
	XAxis *Dashboard_SmallMultiplesAxisProperties[any] `json:"XAxis,omitempty"`

	// YAxis AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-smallmultiplesoptions.html#cfn-quicksight-dashboard-smallmultiplesoptions-yaxis
	YAxis *Dashboard_SmallMultiplesAxisProperties[any] `json:"YAxis,omitempty"`

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
func (r *Dashboard_SmallMultiplesOptions[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.SmallMultiplesOptions"
}
