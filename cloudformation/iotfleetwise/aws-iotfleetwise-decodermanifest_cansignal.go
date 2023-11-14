// Code generated by "go generate". Please don't change this file directly.

package iotfleetwise

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// DecoderManifest_CanSignal AWS CloudFormation Resource (AWS::IoTFleetWise::DecoderManifest.CanSignal)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cansignal.html
type DecoderManifest_CanSignal[T any] struct {

	// Factor AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cansignal.html#cfn-iotfleetwise-decodermanifest-cansignal-factor
	Factor T `json:"Factor"`

	// IsBigEndian AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cansignal.html#cfn-iotfleetwise-decodermanifest-cansignal-isbigendian
	IsBigEndian T `json:"IsBigEndian"`

	// IsSigned AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cansignal.html#cfn-iotfleetwise-decodermanifest-cansignal-issigned
	IsSigned T `json:"IsSigned"`

	// Length AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cansignal.html#cfn-iotfleetwise-decodermanifest-cansignal-length
	Length T `json:"Length"`

	// MessageId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cansignal.html#cfn-iotfleetwise-decodermanifest-cansignal-messageid
	MessageId T `json:"MessageId"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cansignal.html#cfn-iotfleetwise-decodermanifest-cansignal-name
	Name *T `json:"Name,omitempty"`

	// Offset AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cansignal.html#cfn-iotfleetwise-decodermanifest-cansignal-offset
	Offset T `json:"Offset"`

	// StartBit AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cansignal.html#cfn-iotfleetwise-decodermanifest-cansignal-startbit
	StartBit T `json:"StartBit"`

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
func (r *DecoderManifest_CanSignal[any]) AWSCloudFormationType() string {
	return "AWS::IoTFleetWise::DecoderManifest.CanSignal"
}
