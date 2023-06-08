// Code generated by "go generate". Please don't change this file directly.

package databrew

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dataset_Input AWS CloudFormation Resource (AWS::DataBrew::Dataset.Input)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-input.html
type Dataset_Input[T any] struct {

	// DataCatalogInputDefinition AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-input.html#cfn-databrew-dataset-input-datacataloginputdefinition
	DataCatalogInputDefinition *Dataset_DataCatalogInputDefinition[any] `json:"DataCatalogInputDefinition,omitempty"`

	// DatabaseInputDefinition AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-input.html#cfn-databrew-dataset-input-databaseinputdefinition
	DatabaseInputDefinition *Dataset_DatabaseInputDefinition[any] `json:"DatabaseInputDefinition,omitempty"`

	// Metadata AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-input.html#cfn-databrew-dataset-input-metadata
	Metadata *Dataset_Metadata[any] `json:"Metadata,omitempty"`

	// S3InputDefinition AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-input.html#cfn-databrew-dataset-input-s3inputdefinition
	S3InputDefinition *Dataset_S3Location[any] `json:"S3InputDefinition,omitempty"`

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
func (r *Dataset_Input[any]) AWSCloudFormationType() string {
	return "AWS::DataBrew::Dataset.Input"
}
