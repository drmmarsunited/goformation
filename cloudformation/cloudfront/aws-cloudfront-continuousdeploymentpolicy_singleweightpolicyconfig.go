// Code generated by "go generate". Please don't change this file directly.

package cloudfront

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// ContinuousDeploymentPolicy_SingleWeightPolicyConfig AWS CloudFormation Resource (AWS::CloudFront::ContinuousDeploymentPolicy.SingleWeightPolicyConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-singleweightpolicyconfig.html
type ContinuousDeploymentPolicy_SingleWeightPolicyConfig[T any] struct {

	// SessionStickinessConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-singleweightpolicyconfig.html#cfn-cloudfront-continuousdeploymentpolicy-singleweightpolicyconfig-sessionstickinessconfig
	SessionStickinessConfig *ContinuousDeploymentPolicy_SessionStickinessConfig[any] `json:"SessionStickinessConfig,omitempty"`

	// Weight AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-singleweightpolicyconfig.html#cfn-cloudfront-continuousdeploymentpolicy-singleweightpolicyconfig-weight
	Weight T `json:"Weight"`

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
func (r *ContinuousDeploymentPolicy_SingleWeightPolicyConfig[any]) AWSCloudFormationType() string {
	return "AWS::CloudFront::ContinuousDeploymentPolicy.SingleWeightPolicyConfig"
}