// Code generated by "go generate". Please don't change this file directly.

package networkmanager

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// CoreNetwork_CoreNetworkEdge AWS CloudFormation Resource (AWS::NetworkManager::CoreNetwork.CoreNetworkEdge)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-corenetwork-corenetworkedge.html
type CoreNetwork_CoreNetworkEdge struct {

	// Asn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-corenetwork-corenetworkedge.html#cfn-networkmanager-corenetwork-corenetworkedge-asn
	Asn *float64 `json:"Asn,omitempty"`

	// EdgeLocation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-corenetwork-corenetworkedge.html#cfn-networkmanager-corenetwork-corenetworkedge-edgelocation
	EdgeLocation *string `json:"EdgeLocation,omitempty"`

	// InsideCidrBlocks AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-corenetwork-corenetworkedge.html#cfn-networkmanager-corenetwork-corenetworkedge-insidecidrblocks
	InsideCidrBlocks []string `json:"InsideCidrBlocks,omitempty"`

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
func (r *CoreNetwork_CoreNetworkEdge) AWSCloudFormationType() string {
	return "AWS::NetworkManager::CoreNetwork.CoreNetworkEdge"
}
