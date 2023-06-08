// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_ConditionalFormattingCustomIconCondition AWS CloudFormation Resource (AWS::QuickSight::Dashboard.ConditionalFormattingCustomIconCondition)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-conditionalformattingcustomiconcondition.html
type Dashboard_ConditionalFormattingCustomIconCondition[T any] struct {

	// Color AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-conditionalformattingcustomiconcondition.html#cfn-quicksight-dashboard-conditionalformattingcustomiconcondition-color
	Color *string `json:"Color,omitempty"`

	// DisplayConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-conditionalformattingcustomiconcondition.html#cfn-quicksight-dashboard-conditionalformattingcustomiconcondition-displayconfiguration
	DisplayConfiguration *Dashboard_ConditionalFormattingIconDisplayConfiguration[any] `json:"DisplayConfiguration,omitempty"`

	// Expression AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-conditionalformattingcustomiconcondition.html#cfn-quicksight-dashboard-conditionalformattingcustomiconcondition-expression
	Expression string `json:"Expression"`

	// IconOptions AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-conditionalformattingcustomiconcondition.html#cfn-quicksight-dashboard-conditionalformattingcustomiconcondition-iconoptions
	IconOptions *Dashboard_ConditionalFormattingCustomIconOptions[any] `json:"IconOptions"`

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
func (r *Dashboard_ConditionalFormattingCustomIconCondition[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.ConditionalFormattingCustomIconCondition"
}
