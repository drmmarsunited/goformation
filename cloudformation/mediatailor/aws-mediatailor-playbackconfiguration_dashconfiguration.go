// Code generated by "go generate". Please don't change this file directly.

package mediatailor

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// PlaybackConfiguration_DashConfiguration AWS CloudFormation Resource (AWS::MediaTailor::PlaybackConfiguration.DashConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-dashconfiguration.html
type PlaybackConfiguration_DashConfiguration[T any] struct {

	// ManifestEndpointPrefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-dashconfiguration.html#cfn-mediatailor-playbackconfiguration-dashconfiguration-manifestendpointprefix
	ManifestEndpointPrefix *T `json:"ManifestEndpointPrefix,omitempty"`

	// MpdLocation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-dashconfiguration.html#cfn-mediatailor-playbackconfiguration-dashconfiguration-mpdlocation
	MpdLocation *T `json:"MpdLocation,omitempty"`

	// OriginManifestType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-dashconfiguration.html#cfn-mediatailor-playbackconfiguration-dashconfiguration-originmanifesttype
	OriginManifestType *T `json:"OriginManifestType,omitempty"`

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
func (r *PlaybackConfiguration_DashConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::MediaTailor::PlaybackConfiguration.DashConfiguration"
}
