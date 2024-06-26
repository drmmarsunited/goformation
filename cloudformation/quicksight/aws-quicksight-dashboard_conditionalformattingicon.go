// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_ConditionalFormattingIcon AWS CloudFormation Resource (AWS::QuickSight::Dashboard.ConditionalFormattingIcon)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-conditionalformattingicon.html
type Dashboard_ConditionalFormattingIcon[T any] struct {

	// CustomCondition AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-conditionalformattingicon.html#cfn-quicksight-dashboard-conditionalformattingicon-customcondition
	CustomCondition *Dashboard_ConditionalFormattingCustomIconCondition[any] `json:"CustomCondition,omitempty"`

	// IconSet AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-conditionalformattingicon.html#cfn-quicksight-dashboard-conditionalformattingicon-iconset
	IconSet *Dashboard_ConditionalFormattingIconSet[any] `json:"IconSet,omitempty"`

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
func (r *Dashboard_ConditionalFormattingIcon[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.ConditionalFormattingIcon"
}
