// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_PivotTableFieldCollapseStateOption AWS CloudFormation Resource (AWS::QuickSight::Dashboard.PivotTableFieldCollapseStateOption)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablefieldcollapsestateoption.html
type Dashboard_PivotTableFieldCollapseStateOption[T any] struct {

	// State AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablefieldcollapsestateoption.html#cfn-quicksight-dashboard-pivottablefieldcollapsestateoption-state
	State *string `json:"State,omitempty"`

	// Target AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablefieldcollapsestateoption.html#cfn-quicksight-dashboard-pivottablefieldcollapsestateoption-target
	Target *Dashboard_PivotTableFieldCollapseStateTarget[any] `json:"Target"`

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
func (r *Dashboard_PivotTableFieldCollapseStateOption[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.PivotTableFieldCollapseStateOption"
}
