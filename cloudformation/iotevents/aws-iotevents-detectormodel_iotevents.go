// Code generated by "go generate". Please don't change this file directly.

package iotevents

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// DetectorModel_IotEvents AWS CloudFormation Resource (AWS::IoTEvents::DetectorModel.IotEvents)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-iotevents.html
type DetectorModel_IotEvents[T any] struct {

	// InputName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-iotevents.html#cfn-iotevents-detectormodel-iotevents-inputname
	InputName string `json:"InputName"`

	// Payload AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-iotevents.html#cfn-iotevents-detectormodel-iotevents-payload
	Payload *DetectorModel_Payload[any] `json:"Payload,omitempty"`

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
func (r *DetectorModel_IotEvents[any]) AWSCloudFormationType() string {
	return "AWS::IoTEvents::DetectorModel.IotEvents"
}
