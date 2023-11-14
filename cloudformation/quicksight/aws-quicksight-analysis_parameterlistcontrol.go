// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_ParameterListControl AWS CloudFormation Resource (AWS::QuickSight::Analysis.ParameterListControl)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterlistcontrol.html
type Analysis_ParameterListControl[T any] struct {

	// CascadingControlConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterlistcontrol.html#cfn-quicksight-analysis-parameterlistcontrol-cascadingcontrolconfiguration
	CascadingControlConfiguration *Analysis_CascadingControlConfiguration[any] `json:"CascadingControlConfiguration,omitempty"`

	// DisplayOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterlistcontrol.html#cfn-quicksight-analysis-parameterlistcontrol-displayoptions
	DisplayOptions *Analysis_ListControlDisplayOptions[any] `json:"DisplayOptions,omitempty"`

	// ParameterControlId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterlistcontrol.html#cfn-quicksight-analysis-parameterlistcontrol-parametercontrolid
	ParameterControlId T `json:"ParameterControlId"`

	// SelectableValues AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterlistcontrol.html#cfn-quicksight-analysis-parameterlistcontrol-selectablevalues
	SelectableValues *Analysis_ParameterSelectableValues[any] `json:"SelectableValues,omitempty"`

	// SourceParameterName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterlistcontrol.html#cfn-quicksight-analysis-parameterlistcontrol-sourceparametername
	SourceParameterName T `json:"SourceParameterName"`

	// Title AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterlistcontrol.html#cfn-quicksight-analysis-parameterlistcontrol-title
	Title T `json:"Title"`

	// Type AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterlistcontrol.html#cfn-quicksight-analysis-parameterlistcontrol-type
	Type *T `json:"Type,omitempty"`

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
func (r *Analysis_ParameterListControl[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.ParameterListControl"
}
