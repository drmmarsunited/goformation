// Code generated by "go generate". Please don't change this file directly.

package ec2

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// NetworkInsightsAnalysis_AnalysisLoadBalancerTarget AWS CloudFormation Resource (AWS::EC2::NetworkInsightsAnalysis.AnalysisLoadBalancerTarget)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisloadbalancertarget.html
type NetworkInsightsAnalysis_AnalysisLoadBalancerTarget[T any] struct {

	// Address AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisloadbalancertarget.html#cfn-ec2-networkinsightsanalysis-analysisloadbalancertarget-address
	Address *string `json:"Address,omitempty"`

	// AvailabilityZone AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisloadbalancertarget.html#cfn-ec2-networkinsightsanalysis-analysisloadbalancertarget-availabilityzone
	AvailabilityZone *string `json:"AvailabilityZone,omitempty"`

	// Instance AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisloadbalancertarget.html#cfn-ec2-networkinsightsanalysis-analysisloadbalancertarget-instance
	Instance *NetworkInsightsAnalysis_AnalysisComponent[any] `json:"Instance,omitempty"`

	// Port AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisloadbalancertarget.html#cfn-ec2-networkinsightsanalysis-analysisloadbalancertarget-port
	Port *T `json:"Port,omitempty"`

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
func (r *NetworkInsightsAnalysis_AnalysisLoadBalancerTarget[any]) AWSCloudFormationType() string {
	return "AWS::EC2::NetworkInsightsAnalysis.AnalysisLoadBalancerTarget"
}
