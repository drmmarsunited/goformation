// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_ParameterDropDownControl AWS CloudFormation Resource (AWS::QuickSight::Analysis.ParameterDropDownControl)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterdropdowncontrol.html
type Analysis_ParameterDropDownControl struct {

	// CascadingControlConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterdropdowncontrol.html#cfn-quicksight-analysis-parameterdropdowncontrol-cascadingcontrolconfiguration
	CascadingControlConfiguration *Analysis_CascadingControlConfiguration `json:"CascadingControlConfiguration,omitempty"`

	// DisplayOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterdropdowncontrol.html#cfn-quicksight-analysis-parameterdropdowncontrol-displayoptions
	DisplayOptions *Analysis_DropDownControlDisplayOptions `json:"DisplayOptions,omitempty"`

	// ParameterControlId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterdropdowncontrol.html#cfn-quicksight-analysis-parameterdropdowncontrol-parametercontrolid
	ParameterControlId string `json:"ParameterControlId"`

	// SelectableValues AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterdropdowncontrol.html#cfn-quicksight-analysis-parameterdropdowncontrol-selectablevalues
	SelectableValues *Analysis_ParameterSelectableValues `json:"SelectableValues,omitempty"`

	// SourceParameterName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterdropdowncontrol.html#cfn-quicksight-analysis-parameterdropdowncontrol-sourceparametername
	SourceParameterName string `json:"SourceParameterName"`

	// Title AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterdropdowncontrol.html#cfn-quicksight-analysis-parameterdropdowncontrol-title
	Title string `json:"Title"`

	// Type AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterdropdowncontrol.html#cfn-quicksight-analysis-parameterdropdowncontrol-type
	Type *string `json:"Type,omitempty"`

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
func (r *Analysis_ParameterDropDownControl) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.ParameterDropDownControl"
}
