// Code generated by "go generate". Please don't change this file directly.

package glue

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// DataCatalogEncryptionSettings_ConnectionPasswordEncryption AWS CloudFormation Resource (AWS::Glue::DataCatalogEncryptionSettings.ConnectionPasswordEncryption)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-datacatalogencryptionsettings-connectionpasswordencryption.html
type DataCatalogEncryptionSettings_ConnectionPasswordEncryption[T any] struct {

	// KmsKeyId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-datacatalogencryptionsettings-connectionpasswordencryption.html#cfn-glue-datacatalogencryptionsettings-connectionpasswordencryption-kmskeyid
	KmsKeyId *string `json:"KmsKeyId,omitempty"`

	// ReturnConnectionPasswordEncrypted AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-datacatalogencryptionsettings-connectionpasswordencryption.html#cfn-glue-datacatalogencryptionsettings-connectionpasswordencryption-returnconnectionpasswordencrypted
	ReturnConnectionPasswordEncrypted *T `json:"ReturnConnectionPasswordEncrypted,omitempty"`

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
func (r *DataCatalogEncryptionSettings_ConnectionPasswordEncryption[any]) AWSCloudFormationType() string {
	return "AWS::Glue::DataCatalogEncryptionSettings.ConnectionPasswordEncryption"
}
