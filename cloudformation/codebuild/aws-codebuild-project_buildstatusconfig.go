// Code generated by "go generate". Please don't change this file directly.

package codebuild

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Project_BuildStatusConfig AWS CloudFormation Resource (AWS::CodeBuild::Project.BuildStatusConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-buildstatusconfig.html
type Project_BuildStatusConfig[T any] struct {

	// Context AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-buildstatusconfig.html#cfn-codebuild-project-buildstatusconfig-context
	Context *T `json:"Context,omitempty"`

	// TargetUrl AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-buildstatusconfig.html#cfn-codebuild-project-buildstatusconfig-targeturl
	TargetUrl *T `json:"TargetUrl,omitempty"`

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
func (r *Project_BuildStatusConfig[any]) AWSCloudFormationType() string {
	return "AWS::CodeBuild::Project.BuildStatusConfig"
}
