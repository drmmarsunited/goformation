// Code generated by "go generate". Please don't change this file directly.

package mediapackagev2

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// OriginEndpoint_EncryptionContractConfiguration AWS CloudFormation Resource (AWS::MediaPackageV2::OriginEndpoint.EncryptionContractConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-encryptioncontractconfiguration.html
type OriginEndpoint_EncryptionContractConfiguration[T any] struct {

	// PresetSpeke20Audio AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-encryptioncontractconfiguration.html#cfn-mediapackagev2-originendpoint-encryptioncontractconfiguration-presetspeke20audio
	PresetSpeke20Audio T `json:"PresetSpeke20Audio"`

	// PresetSpeke20Video AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-encryptioncontractconfiguration.html#cfn-mediapackagev2-originendpoint-encryptioncontractconfiguration-presetspeke20video
	PresetSpeke20Video T `json:"PresetSpeke20Video"`

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
func (r *OriginEndpoint_EncryptionContractConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::MediaPackageV2::OriginEndpoint.EncryptionContractConfiguration"
}