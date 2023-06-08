// Code generated by "go generate". Please don't change this file directly.

package sagemaker

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Domain_DefaultSpaceSettings AWS CloudFormation Resource (AWS::SageMaker::Domain.DefaultSpaceSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-defaultspacesettings.html
type Domain_DefaultSpaceSettings[T any] struct {

	// ExecutionRole AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-defaultspacesettings.html#cfn-sagemaker-domain-defaultspacesettings-executionrole
	ExecutionRole string `json:"ExecutionRole"`

	// JupyterServerAppSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-defaultspacesettings.html#cfn-sagemaker-domain-defaultspacesettings-jupyterserverappsettings
	JupyterServerAppSettings *Domain_JupyterServerAppSettings[any] `json:"JupyterServerAppSettings,omitempty"`

	// KernelGatewayAppSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-defaultspacesettings.html#cfn-sagemaker-domain-defaultspacesettings-kernelgatewayappsettings
	KernelGatewayAppSettings *Domain_KernelGatewayAppSettings[any] `json:"KernelGatewayAppSettings,omitempty"`

	// SecurityGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-defaultspacesettings.html#cfn-sagemaker-domain-defaultspacesettings-securitygroups
	SecurityGroups []string `json:"SecurityGroups,omitempty"`

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
func (r *Domain_DefaultSpaceSettings[any]) AWSCloudFormationType() string {
	return "AWS::SageMaker::Domain.DefaultSpaceSettings"
}
