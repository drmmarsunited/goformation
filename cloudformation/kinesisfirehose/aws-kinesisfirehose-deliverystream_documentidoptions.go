// Code generated by "go generate". Please don't change this file directly.

package kinesisfirehose

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// DeliveryStream_DocumentIdOptions AWS CloudFormation Resource (AWS::KinesisFirehose::DeliveryStream.DocumentIdOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-documentidoptions.html
type DeliveryStream_DocumentIdOptions[T any] struct {

	// DefaultDocumentIdFormat AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-documentidoptions.html#cfn-kinesisfirehose-deliverystream-documentidoptions-defaultdocumentidformat
	DefaultDocumentIdFormat T `json:"DefaultDocumentIdFormat"`

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
func (r *DeliveryStream_DocumentIdOptions[any]) AWSCloudFormationType() string {
	return "AWS::KinesisFirehose::DeliveryStream.DocumentIdOptions"
}
