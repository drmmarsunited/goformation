// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_FontConfiguration AWS CloudFormation Resource (AWS::QuickSight::Dashboard.FontConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-fontconfiguration.html
type Dashboard_FontConfiguration[T any] struct {

	// FontColor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-fontconfiguration.html#cfn-quicksight-dashboard-fontconfiguration-fontcolor
	FontColor *T `json:"FontColor,omitempty"`

	// FontDecoration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-fontconfiguration.html#cfn-quicksight-dashboard-fontconfiguration-fontdecoration
	FontDecoration *T `json:"FontDecoration,omitempty"`

	// FontSize AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-fontconfiguration.html#cfn-quicksight-dashboard-fontconfiguration-fontsize
	FontSize *Dashboard_FontSize[any] `json:"FontSize,omitempty"`

	// FontStyle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-fontconfiguration.html#cfn-quicksight-dashboard-fontconfiguration-fontstyle
	FontStyle *T `json:"FontStyle,omitempty"`

	// FontWeight AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-fontconfiguration.html#cfn-quicksight-dashboard-fontconfiguration-fontweight
	FontWeight *Dashboard_FontWeight[any] `json:"FontWeight,omitempty"`

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
func (r *Dashboard_FontConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.FontConfiguration"
}
