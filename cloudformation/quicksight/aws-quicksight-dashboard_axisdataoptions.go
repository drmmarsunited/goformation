// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_AxisDataOptions AWS CloudFormation Resource (AWS::QuickSight::Dashboard.AxisDataOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axisdataoptions.html
type Dashboard_AxisDataOptions struct {

	// DateAxisOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axisdataoptions.html#cfn-quicksight-dashboard-axisdataoptions-dateaxisoptions
	DateAxisOptions *Dashboard_DateAxisOptions `json:"DateAxisOptions,omitempty"`

	// NumericAxisOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axisdataoptions.html#cfn-quicksight-dashboard-axisdataoptions-numericaxisoptions
	NumericAxisOptions *Dashboard_NumericAxisOptions `json:"NumericAxisOptions,omitempty"`

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
func (r *Dashboard_AxisDataOptions) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.AxisDataOptions"
}
