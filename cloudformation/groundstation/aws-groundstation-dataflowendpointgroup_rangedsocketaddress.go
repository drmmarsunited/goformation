// Code generated by "go generate". Please don't change this file directly.

package groundstation

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// DataflowEndpointGroup_RangedSocketAddress AWS CloudFormation Resource (AWS::GroundStation::DataflowEndpointGroup.RangedSocketAddress)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroup-rangedsocketaddress.html
type DataflowEndpointGroup_RangedSocketAddress struct {

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroup-rangedsocketaddress.html#cfn-groundstation-dataflowendpointgroup-rangedsocketaddress-name
	Name *string `json:"Name,omitempty"`

	// PortRange AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroup-rangedsocketaddress.html#cfn-groundstation-dataflowendpointgroup-rangedsocketaddress-portrange
	PortRange *DataflowEndpointGroup_IntegerRange `json:"PortRange,omitempty"`

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
func (r *DataflowEndpointGroup_RangedSocketAddress) AWSCloudFormationType() string {
	return "AWS::GroundStation::DataflowEndpointGroup.RangedSocketAddress"
}
