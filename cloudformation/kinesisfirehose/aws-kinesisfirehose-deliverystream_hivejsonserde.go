// Code generated by "go generate". Please don't change this file directly.

package kinesisfirehose

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// DeliveryStream_HiveJsonSerDe AWS CloudFormation Resource (AWS::KinesisFirehose::DeliveryStream.HiveJsonSerDe)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-hivejsonserde.html
type DeliveryStream_HiveJsonSerDe[T any] struct {

	// TimestampFormats AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-hivejsonserde.html#cfn-kinesisfirehose-deliverystream-hivejsonserde-timestampformats
	TimestampFormats []string `json:"TimestampFormats,omitempty"`

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
func (r *DeliveryStream_HiveJsonSerDe[any]) AWSCloudFormationType() string {
	return "AWS::KinesisFirehose::DeliveryStream.HiveJsonSerDe"
}
