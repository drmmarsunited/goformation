// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_ConditionalFormattingCustomIconCondition AWS CloudFormation Resource (AWS::QuickSight::Template.ConditionalFormattingCustomIconCondition)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-conditionalformattingcustomiconcondition.html
type Template_ConditionalFormattingCustomIconCondition[T any] struct {

	// Color AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-conditionalformattingcustomiconcondition.html#cfn-quicksight-template-conditionalformattingcustomiconcondition-color
	Color *string `json:"Color,omitempty"`

	// DisplayConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-conditionalformattingcustomiconcondition.html#cfn-quicksight-template-conditionalformattingcustomiconcondition-displayconfiguration
	DisplayConfiguration *Template_ConditionalFormattingIconDisplayConfiguration[any] `json:"DisplayConfiguration,omitempty"`

	// Expression AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-conditionalformattingcustomiconcondition.html#cfn-quicksight-template-conditionalformattingcustomiconcondition-expression
	Expression string `json:"Expression"`

	// IconOptions AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-conditionalformattingcustomiconcondition.html#cfn-quicksight-template-conditionalformattingcustomiconcondition-iconoptions
	IconOptions *Template_ConditionalFormattingCustomIconOptions[any] `json:"IconOptions"`

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
func (r *Template_ConditionalFormattingCustomIconCondition[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.ConditionalFormattingCustomIconCondition"
}
