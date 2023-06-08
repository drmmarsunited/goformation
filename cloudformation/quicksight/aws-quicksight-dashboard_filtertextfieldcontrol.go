// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_FilterTextFieldControl AWS CloudFormation Resource (AWS::QuickSight::Dashboard.FilterTextFieldControl)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filtertextfieldcontrol.html
type Dashboard_FilterTextFieldControl[T any] struct {

	// DisplayOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filtertextfieldcontrol.html#cfn-quicksight-dashboard-filtertextfieldcontrol-displayoptions
	DisplayOptions *Dashboard_TextFieldControlDisplayOptions[any] `json:"DisplayOptions,omitempty"`

	// FilterControlId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filtertextfieldcontrol.html#cfn-quicksight-dashboard-filtertextfieldcontrol-filtercontrolid
	FilterControlId string `json:"FilterControlId"`

	// SourceFilterId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filtertextfieldcontrol.html#cfn-quicksight-dashboard-filtertextfieldcontrol-sourcefilterid
	SourceFilterId string `json:"SourceFilterId"`

	// Title AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filtertextfieldcontrol.html#cfn-quicksight-dashboard-filtertextfieldcontrol-title
	Title string `json:"Title"`

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
func (r *Dashboard_FilterTextFieldControl[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.FilterTextFieldControl"
}
