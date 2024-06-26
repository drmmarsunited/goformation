// Code generated by "go generate". Please don't change this file directly.

package ssmcontacts

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Rotation_WeeklySetting AWS CloudFormation Resource (AWS::SSMContacts::Rotation.WeeklySetting)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-rotation-weeklysetting.html
type Rotation_WeeklySetting[T any] struct {

	// DayOfWeek AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-rotation-weeklysetting.html#cfn-ssmcontacts-rotation-weeklysetting-dayofweek
	DayOfWeek T `json:"DayOfWeek"`

	// HandOffTime AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-rotation-weeklysetting.html#cfn-ssmcontacts-rotation-weeklysetting-handofftime
	HandOffTime T `json:"HandOffTime"`

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
func (r *Rotation_WeeklySetting[any]) AWSCloudFormationType() string {
	return "AWS::SSMContacts::Rotation.WeeklySetting"
}
