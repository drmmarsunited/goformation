// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_FreeFormLayoutElementBackgroundStyle AWS CloudFormation Resource (AWS::QuickSight::Dashboard.FreeFormLayoutElementBackgroundStyle)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutelementbackgroundstyle.html
type Dashboard_FreeFormLayoutElementBackgroundStyle[T any] struct {

	// Color AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutelementbackgroundstyle.html#cfn-quicksight-dashboard-freeformlayoutelementbackgroundstyle-color
	Color *T `json:"Color,omitempty"`

	// Visibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutelementbackgroundstyle.html#cfn-quicksight-dashboard-freeformlayoutelementbackgroundstyle-visibility
	Visibility *T `json:"Visibility,omitempty"`

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
func (r *Dashboard_FreeFormLayoutElementBackgroundStyle[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.FreeFormLayoutElementBackgroundStyle"
}
