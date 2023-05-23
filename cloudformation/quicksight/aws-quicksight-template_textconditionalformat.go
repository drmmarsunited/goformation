// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_TextConditionalFormat AWS CloudFormation Resource (AWS::QuickSight::Template.TextConditionalFormat)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-textconditionalformat.html
type Template_TextConditionalFormat struct {

	// BackgroundColor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-textconditionalformat.html#cfn-quicksight-template-textconditionalformat-backgroundcolor
	BackgroundColor *Template_ConditionalFormattingColor `json:"BackgroundColor,omitempty"`

	// Icon AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-textconditionalformat.html#cfn-quicksight-template-textconditionalformat-icon
	Icon *Template_ConditionalFormattingIcon `json:"Icon,omitempty"`

	// TextColor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-textconditionalformat.html#cfn-quicksight-template-textconditionalformat-textcolor
	TextColor *Template_ConditionalFormattingColor `json:"TextColor,omitempty"`

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
func (r *Template_TextConditionalFormat) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.TextConditionalFormat"
}
