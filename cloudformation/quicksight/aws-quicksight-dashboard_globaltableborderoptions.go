// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_GlobalTableBorderOptions AWS CloudFormation Resource (AWS::QuickSight::Dashboard.GlobalTableBorderOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-globaltableborderoptions.html
type Dashboard_GlobalTableBorderOptions[T any] struct {

	// SideSpecificBorder AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-globaltableborderoptions.html#cfn-quicksight-dashboard-globaltableborderoptions-sidespecificborder
	SideSpecificBorder *Dashboard_TableSideBorderOptions[any] `json:"SideSpecificBorder,omitempty"`

	// UniformBorder AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-globaltableborderoptions.html#cfn-quicksight-dashboard-globaltableborderoptions-uniformborder
	UniformBorder *Dashboard_TableBorderOptions[any] `json:"UniformBorder,omitempty"`

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
func (r *Dashboard_GlobalTableBorderOptions[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.GlobalTableBorderOptions"
}
