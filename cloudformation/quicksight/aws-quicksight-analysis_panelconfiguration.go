// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_PanelConfiguration AWS CloudFormation Resource (AWS::QuickSight::Analysis.PanelConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-panelconfiguration.html
type Analysis_PanelConfiguration[T any] struct {

	// BackgroundColor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-panelconfiguration.html#cfn-quicksight-analysis-panelconfiguration-backgroundcolor
	BackgroundColor *T `json:"BackgroundColor,omitempty"`

	// BackgroundVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-panelconfiguration.html#cfn-quicksight-analysis-panelconfiguration-backgroundvisibility
	BackgroundVisibility *T `json:"BackgroundVisibility,omitempty"`

	// BorderColor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-panelconfiguration.html#cfn-quicksight-analysis-panelconfiguration-bordercolor
	BorderColor *T `json:"BorderColor,omitempty"`

	// BorderStyle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-panelconfiguration.html#cfn-quicksight-analysis-panelconfiguration-borderstyle
	BorderStyle *T `json:"BorderStyle,omitempty"`

	// BorderThickness AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-panelconfiguration.html#cfn-quicksight-analysis-panelconfiguration-borderthickness
	BorderThickness *T `json:"BorderThickness,omitempty"`

	// BorderVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-panelconfiguration.html#cfn-quicksight-analysis-panelconfiguration-bordervisibility
	BorderVisibility *T `json:"BorderVisibility,omitempty"`

	// GutterSpacing AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-panelconfiguration.html#cfn-quicksight-analysis-panelconfiguration-gutterspacing
	GutterSpacing *T `json:"GutterSpacing,omitempty"`

	// GutterVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-panelconfiguration.html#cfn-quicksight-analysis-panelconfiguration-guttervisibility
	GutterVisibility *T `json:"GutterVisibility,omitempty"`

	// Title AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-panelconfiguration.html#cfn-quicksight-analysis-panelconfiguration-title
	Title *Analysis_PanelTitleOptions[any] `json:"Title,omitempty"`

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
func (r *Analysis_PanelConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.PanelConfiguration"
}
