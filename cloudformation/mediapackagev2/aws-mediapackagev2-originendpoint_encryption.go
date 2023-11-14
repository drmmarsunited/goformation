// Code generated by "go generate". Please don't change this file directly.

package mediapackagev2

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// OriginEndpoint_Encryption AWS CloudFormation Resource (AWS::MediaPackageV2::OriginEndpoint.Encryption)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-encryption.html
type OriginEndpoint_Encryption[T any] struct {

	// ConstantInitializationVector AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-encryption.html#cfn-mediapackagev2-originendpoint-encryption-constantinitializationvector
	ConstantInitializationVector *T `json:"ConstantInitializationVector,omitempty"`

	// EncryptionMethod AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-encryption.html#cfn-mediapackagev2-originendpoint-encryption-encryptionmethod
	EncryptionMethod *OriginEndpoint_EncryptionMethod[any] `json:"EncryptionMethod"`

	// KeyRotationIntervalSeconds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-encryption.html#cfn-mediapackagev2-originendpoint-encryption-keyrotationintervalseconds
	KeyRotationIntervalSeconds *T `json:"KeyRotationIntervalSeconds,omitempty"`

	// SpekeKeyProvider AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-encryption.html#cfn-mediapackagev2-originendpoint-encryption-spekekeyprovider
	SpekeKeyProvider *OriginEndpoint_SpekeKeyProvider[any] `json:"SpekeKeyProvider"`

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
func (r *OriginEndpoint_Encryption[any]) AWSCloudFormationType() string {
	return "AWS::MediaPackageV2::OriginEndpoint.Encryption"
}
