// Code generated by "go generate". Please don't change this file directly.

package connect

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// InstanceStorageConfig_EncryptionConfig AWS CloudFormation Resource (AWS::Connect::InstanceStorageConfig.EncryptionConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instancestorageconfig-encryptionconfig.html
type InstanceStorageConfig_EncryptionConfig[T any] struct {

	// EncryptionType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instancestorageconfig-encryptionconfig.html#cfn-connect-instancestorageconfig-encryptionconfig-encryptiontype
	EncryptionType T `json:"EncryptionType"`

	// KeyId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instancestorageconfig-encryptionconfig.html#cfn-connect-instancestorageconfig-encryptionconfig-keyid
	KeyId T `json:"KeyId"`

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
func (r *InstanceStorageConfig_EncryptionConfig[any]) AWSCloudFormationType() string {
	return "AWS::Connect::InstanceStorageConfig.EncryptionConfig"
}
