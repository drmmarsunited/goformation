// Code generated by "go generate". Please don't change this file directly.

package cloudfront

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// ContinuousDeploymentPolicy_SingleHeaderConfig AWS CloudFormation Resource (AWS::CloudFront::ContinuousDeploymentPolicy.SingleHeaderConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-singleheaderconfig.html
type ContinuousDeploymentPolicy_SingleHeaderConfig struct {

	// Header AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-singleheaderconfig.html#cfn-cloudfront-continuousdeploymentpolicy-singleheaderconfig-header
	Header string `json:"Header"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-singleheaderconfig.html#cfn-cloudfront-continuousdeploymentpolicy-singleheaderconfig-value
	Value string `json:"Value"`

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
func (r *ContinuousDeploymentPolicy_SingleHeaderConfig) AWSCloudFormationType() string {
	return "AWS::CloudFront::ContinuousDeploymentPolicy.SingleHeaderConfig"
}
