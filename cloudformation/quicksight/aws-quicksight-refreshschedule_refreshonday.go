// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// RefreshSchedule_RefreshOnDay AWS CloudFormation Resource (AWS::QuickSight::RefreshSchedule.RefreshOnDay)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-refreshschedule-refreshonday.html
type RefreshSchedule_RefreshOnDay[T any] struct {

	// DayOfMonth AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-refreshschedule-refreshonday.html#cfn-quicksight-refreshschedule-refreshonday-dayofmonth
	DayOfMonth *string `json:"DayOfMonth,omitempty"`

	// DayOfWeek AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-refreshschedule-refreshonday.html#cfn-quicksight-refreshschedule-refreshonday-dayofweek
	DayOfWeek *string `json:"DayOfWeek,omitempty"`

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
func (r *RefreshSchedule_RefreshOnDay[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::RefreshSchedule.RefreshOnDay"
}
