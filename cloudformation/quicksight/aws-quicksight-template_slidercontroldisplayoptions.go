// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_SliderControlDisplayOptions AWS CloudFormation Resource (AWS::QuickSight::Template.SliderControlDisplayOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-slidercontroldisplayoptions.html
type Template_SliderControlDisplayOptions[T any] struct {

	// InfoIconLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-slidercontroldisplayoptions.html#cfn-quicksight-template-slidercontroldisplayoptions-infoiconlabeloptions
	InfoIconLabelOptions *Template_SheetControlInfoIconLabelOptions[any] `json:"InfoIconLabelOptions,omitempty"`

	// TitleOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-slidercontroldisplayoptions.html#cfn-quicksight-template-slidercontroldisplayoptions-titleoptions
	TitleOptions *Template_LabelOptions[any] `json:"TitleOptions,omitempty"`

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
func (r *Template_SliderControlDisplayOptions[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.SliderControlDisplayOptions"
}
