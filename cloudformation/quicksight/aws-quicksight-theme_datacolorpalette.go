// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Theme_DataColorPalette AWS CloudFormation Resource (AWS::QuickSight::Theme.DataColorPalette)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-datacolorpalette.html
type Theme_DataColorPalette[T any] struct {

	// Colors AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-datacolorpalette.html#cfn-quicksight-theme-datacolorpalette-colors
	Colors []T `json:"Colors,omitempty"`

	// EmptyFillColor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-datacolorpalette.html#cfn-quicksight-theme-datacolorpalette-emptyfillcolor
	EmptyFillColor *T `json:"EmptyFillColor,omitempty"`

	// MinMaxGradient AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-datacolorpalette.html#cfn-quicksight-theme-datacolorpalette-minmaxgradient
	MinMaxGradient []T `json:"MinMaxGradient,omitempty"`

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
func (r *Theme_DataColorPalette[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Theme.DataColorPalette"
}
