// Code generated by "go generate". Please don't change this file directly.

package greengrassv2

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Deployment_DeploymentComponentUpdatePolicy AWS CloudFormation Resource (AWS::GreengrassV2::Deployment.DeploymentComponentUpdatePolicy)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-deployment-deploymentcomponentupdatepolicy.html
type Deployment_DeploymentComponentUpdatePolicy struct {

	// Action AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-deployment-deploymentcomponentupdatepolicy.html#cfn-greengrassv2-deployment-deploymentcomponentupdatepolicy-action
	Action *string `json:"Action,omitempty"`

	// TimeoutInSeconds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-deployment-deploymentcomponentupdatepolicy.html#cfn-greengrassv2-deployment-deploymentcomponentupdatepolicy-timeoutinseconds
	TimeoutInSeconds *int `json:"TimeoutInSeconds,omitempty"`

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
func (r *Deployment_DeploymentComponentUpdatePolicy) AWSCloudFormationType() string {
	return "AWS::GreengrassV2::Deployment.DeploymentComponentUpdatePolicy"
}
