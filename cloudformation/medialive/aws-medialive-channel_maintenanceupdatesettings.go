// Code generated by "go generate". Please don't change this file directly.

package medialive

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Channel_MaintenanceUpdateSettings AWS CloudFormation Resource (AWS::MediaLive::Channel.MaintenanceUpdateSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-maintenanceupdatesettings.html
type Channel_MaintenanceUpdateSettings[T any] struct {

	// MaintenanceDay AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-maintenanceupdatesettings.html#cfn-medialive-channel-maintenanceupdatesettings-maintenanceday
	MaintenanceDay *T `json:"MaintenanceDay,omitempty"`

	// MaintenanceScheduledDate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-maintenanceupdatesettings.html#cfn-medialive-channel-maintenanceupdatesettings-maintenancescheduleddate
	MaintenanceScheduledDate *T `json:"MaintenanceScheduledDate,omitempty"`

	// MaintenanceStartTime AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-maintenanceupdatesettings.html#cfn-medialive-channel-maintenanceupdatesettings-maintenancestarttime
	MaintenanceStartTime *T `json:"MaintenanceStartTime,omitempty"`

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
func (r *Channel_MaintenanceUpdateSettings[any]) AWSCloudFormationType() string {
	return "AWS::MediaLive::Channel.MaintenanceUpdateSettings"
}
