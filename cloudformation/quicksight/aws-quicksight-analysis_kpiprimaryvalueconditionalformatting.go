// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_KPIPrimaryValueConditionalFormatting AWS CloudFormation Resource (AWS::QuickSight::Analysis.KPIPrimaryValueConditionalFormatting)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-kpiprimaryvalueconditionalformatting.html
type Analysis_KPIPrimaryValueConditionalFormatting[T any] struct {

	// Icon AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-kpiprimaryvalueconditionalformatting.html#cfn-quicksight-analysis-kpiprimaryvalueconditionalformatting-icon
	Icon *Analysis_ConditionalFormattingIcon[any] `json:"Icon,omitempty"`

	// TextColor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-kpiprimaryvalueconditionalformatting.html#cfn-quicksight-analysis-kpiprimaryvalueconditionalformatting-textcolor
	TextColor *Analysis_ConditionalFormattingColor[any] `json:"TextColor,omitempty"`

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
func (r *Analysis_KPIPrimaryValueConditionalFormatting[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.KPIPrimaryValueConditionalFormatting"
}
