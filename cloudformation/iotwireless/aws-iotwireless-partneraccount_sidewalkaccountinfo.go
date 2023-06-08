// Code generated by "go generate". Please don't change this file directly.

package iotwireless

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// PartnerAccount_SidewalkAccountInfo AWS CloudFormation Resource (AWS::IoTWireless::PartnerAccount.SidewalkAccountInfo)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-partneraccount-sidewalkaccountinfo.html
type PartnerAccount_SidewalkAccountInfo[T any] struct {

	// AppServerPrivateKey AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-partneraccount-sidewalkaccountinfo.html#cfn-iotwireless-partneraccount-sidewalkaccountinfo-appserverprivatekey
	AppServerPrivateKey string `json:"AppServerPrivateKey"`

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
func (r *PartnerAccount_SidewalkAccountInfo[any]) AWSCloudFormationType() string {
	return "AWS::IoTWireless::PartnerAccount.SidewalkAccountInfo"
}
