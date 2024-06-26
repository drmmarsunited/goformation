// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_ReferenceLineLabelConfiguration AWS CloudFormation Resource (AWS::QuickSight::Template.ReferenceLineLabelConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referencelinelabelconfiguration.html
type Template_ReferenceLineLabelConfiguration[T any] struct {

	// CustomLabelConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referencelinelabelconfiguration.html#cfn-quicksight-template-referencelinelabelconfiguration-customlabelconfiguration
	CustomLabelConfiguration *Template_ReferenceLineCustomLabelConfiguration[any] `json:"CustomLabelConfiguration,omitempty"`

	// FontColor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referencelinelabelconfiguration.html#cfn-quicksight-template-referencelinelabelconfiguration-fontcolor
	FontColor *T `json:"FontColor,omitempty"`

	// FontConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referencelinelabelconfiguration.html#cfn-quicksight-template-referencelinelabelconfiguration-fontconfiguration
	FontConfiguration *Template_FontConfiguration[any] `json:"FontConfiguration,omitempty"`

	// HorizontalPosition AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referencelinelabelconfiguration.html#cfn-quicksight-template-referencelinelabelconfiguration-horizontalposition
	HorizontalPosition *T `json:"HorizontalPosition,omitempty"`

	// ValueLabelConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referencelinelabelconfiguration.html#cfn-quicksight-template-referencelinelabelconfiguration-valuelabelconfiguration
	ValueLabelConfiguration *Template_ReferenceLineValueLabelConfiguration[any] `json:"ValueLabelConfiguration,omitempty"`

	// VerticalPosition AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referencelinelabelconfiguration.html#cfn-quicksight-template-referencelinelabelconfiguration-verticalposition
	VerticalPosition *T `json:"VerticalPosition,omitempty"`

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
func (r *Template_ReferenceLineLabelConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.ReferenceLineLabelConfiguration"
}
