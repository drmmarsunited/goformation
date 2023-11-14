// Code generated by "go generate". Please don't change this file directly.

package sagemaker

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// AppImageConfig_FileSystemConfig AWS CloudFormation Resource (AWS::SageMaker::AppImageConfig.FileSystemConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-filesystemconfig.html
type AppImageConfig_FileSystemConfig[T any] struct {

	// DefaultGid AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-filesystemconfig.html#cfn-sagemaker-appimageconfig-filesystemconfig-defaultgid
	DefaultGid *T `json:"DefaultGid,omitempty"`

	// DefaultUid AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-filesystemconfig.html#cfn-sagemaker-appimageconfig-filesystemconfig-defaultuid
	DefaultUid *T `json:"DefaultUid,omitempty"`

	// MountPath AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-filesystemconfig.html#cfn-sagemaker-appimageconfig-filesystemconfig-mountpath
	MountPath *T `json:"MountPath,omitempty"`

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
func (r *AppImageConfig_FileSystemConfig[any]) AWSCloudFormationType() string {
	return "AWS::SageMaker::AppImageConfig.FileSystemConfig"
}
