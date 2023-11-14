// Code generated by "go generate". Please don't change this file directly.

package datasync

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// LocationAzureBlob_AzureBlobSasConfiguration AWS CloudFormation Resource (AWS::DataSync::LocationAzureBlob.AzureBlobSasConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationazureblob-azureblobsasconfiguration.html
type LocationAzureBlob_AzureBlobSasConfiguration[T any] struct {

	// AzureBlobSasToken AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationazureblob-azureblobsasconfiguration.html#cfn-datasync-locationazureblob-azureblobsasconfiguration-azureblobsastoken
	AzureBlobSasToken T `json:"AzureBlobSasToken"`

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
func (r *LocationAzureBlob_AzureBlobSasConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::DataSync::LocationAzureBlob.AzureBlobSasConfiguration"
}
