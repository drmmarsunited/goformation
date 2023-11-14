// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Topic_DatasetMetadata AWS CloudFormation Resource (AWS::QuickSight::Topic.DatasetMetadata)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-datasetmetadata.html
type Topic_DatasetMetadata[T any] struct {

	// CalculatedFields AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-datasetmetadata.html#cfn-quicksight-topic-datasetmetadata-calculatedfields
	CalculatedFields []Topic_TopicCalculatedField[any] `json:"CalculatedFields,omitempty"`

	// Columns AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-datasetmetadata.html#cfn-quicksight-topic-datasetmetadata-columns
	Columns []Topic_TopicColumn[any] `json:"Columns,omitempty"`

	// DataAggregation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-datasetmetadata.html#cfn-quicksight-topic-datasetmetadata-dataaggregation
	DataAggregation *Topic_DataAggregation[any] `json:"DataAggregation,omitempty"`

	// DatasetArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-datasetmetadata.html#cfn-quicksight-topic-datasetmetadata-datasetarn
	DatasetArn T `json:"DatasetArn"`

	// DatasetDescription AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-datasetmetadata.html#cfn-quicksight-topic-datasetmetadata-datasetdescription
	DatasetDescription *T `json:"DatasetDescription,omitempty"`

	// DatasetName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-datasetmetadata.html#cfn-quicksight-topic-datasetmetadata-datasetname
	DatasetName *T `json:"DatasetName,omitempty"`

	// Filters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-datasetmetadata.html#cfn-quicksight-topic-datasetmetadata-filters
	Filters []Topic_TopicFilter[any] `json:"Filters,omitempty"`

	// NamedEntities AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-datasetmetadata.html#cfn-quicksight-topic-datasetmetadata-namedentities
	NamedEntities []Topic_TopicNamedEntity[any] `json:"NamedEntities,omitempty"`

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
func (r *Topic_DatasetMetadata[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Topic.DatasetMetadata"
}
