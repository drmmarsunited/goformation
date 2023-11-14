// Code generated by "go generate". Please don't change this file directly.

package nimblestudio

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// LaunchProfile_StreamConfiguration AWS CloudFormation Resource (AWS::NimbleStudio::LaunchProfile.StreamConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html
type LaunchProfile_StreamConfiguration[T any] struct {

	// AutomaticTerminationMode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-automaticterminationmode
	AutomaticTerminationMode *T `json:"AutomaticTerminationMode,omitempty"`

	// ClipboardMode AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-clipboardmode
	ClipboardMode T `json:"ClipboardMode"`

	// Ec2InstanceTypes AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-ec2instancetypes
	Ec2InstanceTypes []T `json:"Ec2InstanceTypes"`

	// MaxSessionLengthInMinutes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-maxsessionlengthinminutes
	MaxSessionLengthInMinutes *T `json:"MaxSessionLengthInMinutes,omitempty"`

	// MaxStoppedSessionLengthInMinutes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-maxstoppedsessionlengthinminutes
	MaxStoppedSessionLengthInMinutes *T `json:"MaxStoppedSessionLengthInMinutes,omitempty"`

	// SessionBackup AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-sessionbackup
	SessionBackup *LaunchProfile_StreamConfigurationSessionBackup[any] `json:"SessionBackup,omitempty"`

	// SessionPersistenceMode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-sessionpersistencemode
	SessionPersistenceMode *T `json:"SessionPersistenceMode,omitempty"`

	// SessionStorage AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-sessionstorage
	SessionStorage *LaunchProfile_StreamConfigurationSessionStorage[any] `json:"SessionStorage,omitempty"`

	// StreamingImageIds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-streamingimageids
	StreamingImageIds []T `json:"StreamingImageIds"`

	// VolumeConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-volumeconfiguration
	VolumeConfiguration *LaunchProfile_VolumeConfiguration[any] `json:"VolumeConfiguration,omitempty"`

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
func (r *LaunchProfile_StreamConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::NimbleStudio::LaunchProfile.StreamConfiguration"
}
