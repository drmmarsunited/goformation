// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_PanelConfiguration AWS CloudFormation Resource (AWS::QuickSight::Template.PanelConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-panelconfiguration.html
type Template_PanelConfiguration[T any] struct {

	// BackgroundColor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-panelconfiguration.html#cfn-quicksight-template-panelconfiguration-backgroundcolor
	BackgroundColor *string `json:"BackgroundColor,omitempty"`

	// BackgroundVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-panelconfiguration.html#cfn-quicksight-template-panelconfiguration-backgroundvisibility
	BackgroundVisibility *string `json:"BackgroundVisibility,omitempty"`

	// BorderColor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-panelconfiguration.html#cfn-quicksight-template-panelconfiguration-bordercolor
	BorderColor *string `json:"BorderColor,omitempty"`

	// BorderStyle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-panelconfiguration.html#cfn-quicksight-template-panelconfiguration-borderstyle
	BorderStyle *string `json:"BorderStyle,omitempty"`

	// BorderThickness AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-panelconfiguration.html#cfn-quicksight-template-panelconfiguration-borderthickness
	BorderThickness *string `json:"BorderThickness,omitempty"`

	// BorderVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-panelconfiguration.html#cfn-quicksight-template-panelconfiguration-bordervisibility
	BorderVisibility *string `json:"BorderVisibility,omitempty"`

	// GutterSpacing AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-panelconfiguration.html#cfn-quicksight-template-panelconfiguration-gutterspacing
	GutterSpacing *string `json:"GutterSpacing,omitempty"`

	// GutterVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-panelconfiguration.html#cfn-quicksight-template-panelconfiguration-guttervisibility
	GutterVisibility *string `json:"GutterVisibility,omitempty"`

	// Title AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-panelconfiguration.html#cfn-quicksight-template-panelconfiguration-title
	Title *Template_PanelTitleOptions[any] `json:"Title,omitempty"`

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
func (r *Template_PanelConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.PanelConfiguration"
}
