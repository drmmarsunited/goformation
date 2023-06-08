// Code generated by "go generate". Please don't change this file directly.

package mediapackage

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// PackagingConfiguration_EncryptionContractConfiguration AWS CloudFormation Resource (AWS::MediaPackage::PackagingConfiguration.EncryptionContractConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-encryptioncontractconfiguration.html
type PackagingConfiguration_EncryptionContractConfiguration[T any] struct {

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
func (r *PackagingConfiguration_EncryptionContractConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::MediaPackage::PackagingConfiguration.EncryptionContractConfiguration"
}
