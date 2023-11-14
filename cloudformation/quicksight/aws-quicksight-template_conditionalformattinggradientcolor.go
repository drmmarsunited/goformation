// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_ConditionalFormattingGradientColor AWS CloudFormation Resource (AWS::QuickSight::Template.ConditionalFormattingGradientColor)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-conditionalformattinggradientcolor.html
type Template_ConditionalFormattingGradientColor[T any] struct {

	// Color AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-conditionalformattinggradientcolor.html#cfn-quicksight-template-conditionalformattinggradientcolor-color
	Color *Template_GradientColor[any] `json:"Color"`

	// Expression AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-conditionalformattinggradientcolor.html#cfn-quicksight-template-conditionalformattinggradientcolor-expression
	Expression T `json:"Expression"`

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
func (r *Template_ConditionalFormattingGradientColor[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.ConditionalFormattingGradientColor"
}
