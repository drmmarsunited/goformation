// Code generated by "go generate". Please don't change this file directly.

package mediaconnect

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// FlowSource_GatewayBridgeSource AWS CloudFormation Resource (AWS::MediaConnect::FlowSource.GatewayBridgeSource)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-gatewaybridgesource.html
type FlowSource_GatewayBridgeSource[T any] struct {

	// BridgeArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-gatewaybridgesource.html#cfn-mediaconnect-flowsource-gatewaybridgesource-bridgearn
	BridgeArn T `json:"BridgeArn"`

	// VpcInterfaceAttachment AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-gatewaybridgesource.html#cfn-mediaconnect-flowsource-gatewaybridgesource-vpcinterfaceattachment
	VpcInterfaceAttachment *FlowSource_VpcInterfaceAttachment[any] `json:"VpcInterfaceAttachment,omitempty"`

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
func (r *FlowSource_GatewayBridgeSource[any]) AWSCloudFormationType() string {
	return "AWS::MediaConnect::FlowSource.GatewayBridgeSource"
}
