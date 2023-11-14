// Code generated by "go generate". Please don't change this file directly.

package medialive

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Channel_UdpGroupSettings AWS CloudFormation Resource (AWS::MediaLive::Channel.UdpGroupSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-udpgroupsettings.html
type Channel_UdpGroupSettings[T any] struct {

	// InputLossAction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-udpgroupsettings.html#cfn-medialive-channel-udpgroupsettings-inputlossaction
	InputLossAction *T `json:"InputLossAction,omitempty"`

	// TimedMetadataId3Frame AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-udpgroupsettings.html#cfn-medialive-channel-udpgroupsettings-timedmetadataid3frame
	TimedMetadataId3Frame *T `json:"TimedMetadataId3Frame,omitempty"`

	// TimedMetadataId3Period AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-udpgroupsettings.html#cfn-medialive-channel-udpgroupsettings-timedmetadataid3period
	TimedMetadataId3Period *T `json:"TimedMetadataId3Period,omitempty"`

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
func (r *Channel_UdpGroupSettings[any]) AWSCloudFormationType() string {
	return "AWS::MediaLive::Channel.UdpGroupSettings"
}
