// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_ParameterSliderControl AWS CloudFormation Resource (AWS::QuickSight::Analysis.ParameterSliderControl)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterslidercontrol.html
type Analysis_ParameterSliderControl[T any] struct {

	// DisplayOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterslidercontrol.html#cfn-quicksight-analysis-parameterslidercontrol-displayoptions
	DisplayOptions *Analysis_SliderControlDisplayOptions[any] `json:"DisplayOptions,omitempty"`

	// MaximumValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterslidercontrol.html#cfn-quicksight-analysis-parameterslidercontrol-maximumvalue
	MaximumValue T `json:"MaximumValue"`

	// MinimumValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterslidercontrol.html#cfn-quicksight-analysis-parameterslidercontrol-minimumvalue
	MinimumValue T `json:"MinimumValue"`

	// ParameterControlId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterslidercontrol.html#cfn-quicksight-analysis-parameterslidercontrol-parametercontrolid
	ParameterControlId T `json:"ParameterControlId"`

	// SourceParameterName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterslidercontrol.html#cfn-quicksight-analysis-parameterslidercontrol-sourceparametername
	SourceParameterName T `json:"SourceParameterName"`

	// StepSize AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterslidercontrol.html#cfn-quicksight-analysis-parameterslidercontrol-stepsize
	StepSize T `json:"StepSize"`

	// Title AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterslidercontrol.html#cfn-quicksight-analysis-parameterslidercontrol-title
	Title T `json:"Title"`

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
func (r *Analysis_ParameterSliderControl[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.ParameterSliderControl"
}
