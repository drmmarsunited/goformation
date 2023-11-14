// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_TextAreaControlDisplayOptions AWS CloudFormation Resource (AWS::QuickSight::Dashboard.TextAreaControlDisplayOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-textareacontroldisplayoptions.html
type Dashboard_TextAreaControlDisplayOptions[T any] struct {

	// InfoIconLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-textareacontroldisplayoptions.html#cfn-quicksight-dashboard-textareacontroldisplayoptions-infoiconlabeloptions
	InfoIconLabelOptions *Dashboard_SheetControlInfoIconLabelOptions[any] `json:"InfoIconLabelOptions,omitempty"`

	// PlaceholderOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-textareacontroldisplayoptions.html#cfn-quicksight-dashboard-textareacontroldisplayoptions-placeholderoptions
	PlaceholderOptions *Dashboard_TextControlPlaceholderOptions[any] `json:"PlaceholderOptions,omitempty"`

	// TitleOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-textareacontroldisplayoptions.html#cfn-quicksight-dashboard-textareacontroldisplayoptions-titleoptions
	TitleOptions *Dashboard_LabelOptions[any] `json:"TitleOptions,omitempty"`

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
func (r *Dashboard_TextAreaControlDisplayOptions[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.TextAreaControlDisplayOptions"
}
