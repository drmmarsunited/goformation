// Code generated by "go generate". Please don't change this file directly.

package medialive

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Channel_MaintenanceCreateSettings AWS CloudFormation Resource (AWS::MediaLive::Channel.MaintenanceCreateSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-maintenancecreatesettings.html
type Channel_MaintenanceCreateSettings[T any] struct {

	// MaintenanceDay AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-maintenancecreatesettings.html#cfn-medialive-channel-maintenancecreatesettings-maintenanceday
	MaintenanceDay *T `json:"MaintenanceDay,omitempty"`

	// MaintenanceStartTime AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-maintenancecreatesettings.html#cfn-medialive-channel-maintenancecreatesettings-maintenancestarttime
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
func (r *Channel_MaintenanceCreateSettings[any]) AWSCloudFormationType() string {
	return "AWS::MediaLive::Channel.MaintenanceCreateSettings"
}
