// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_DateTimePickerControlDisplayOptions AWS CloudFormation Resource (AWS::QuickSight::Dashboard.DateTimePickerControlDisplayOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datetimepickercontroldisplayoptions.html
type Dashboard_DateTimePickerControlDisplayOptions[T any] struct {

	// DateTimeFormat AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datetimepickercontroldisplayoptions.html#cfn-quicksight-dashboard-datetimepickercontroldisplayoptions-datetimeformat
	DateTimeFormat *T `json:"DateTimeFormat,omitempty"`

	// InfoIconLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datetimepickercontroldisplayoptions.html#cfn-quicksight-dashboard-datetimepickercontroldisplayoptions-infoiconlabeloptions
	InfoIconLabelOptions *Dashboard_SheetControlInfoIconLabelOptions[any] `json:"InfoIconLabelOptions,omitempty"`

	// TitleOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datetimepickercontroldisplayoptions.html#cfn-quicksight-dashboard-datetimepickercontroldisplayoptions-titleoptions
	TitleOptions *Dashboard_LabelOptions[any] `json:"TitleOptions,omitempty"`

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
func (r *Dashboard_DateTimePickerControlDisplayOptions[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.DateTimePickerControlDisplayOptions"
}
