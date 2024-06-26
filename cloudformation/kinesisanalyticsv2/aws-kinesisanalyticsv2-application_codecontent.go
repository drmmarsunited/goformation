// Code generated by "go generate". Please don't change this file directly.

package kinesisanalyticsv2

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Application_CodeContent AWS CloudFormation Resource (AWS::KinesisAnalyticsV2::Application.CodeContent)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-codecontent.html
type Application_CodeContent[T any] struct {

	// S3ContentLocation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-codecontent.html#cfn-kinesisanalyticsv2-application-codecontent-s3contentlocation
	S3ContentLocation *Application_S3ContentLocation[any] `json:"S3ContentLocation,omitempty"`

	// TextContent AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-codecontent.html#cfn-kinesisanalyticsv2-application-codecontent-textcontent
	TextContent *T `json:"TextContent,omitempty"`

	// ZipFileContent AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-codecontent.html#cfn-kinesisanalyticsv2-application-codecontent-zipfilecontent
	ZipFileContent *T `json:"ZipFileContent,omitempty"`

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
func (r *Application_CodeContent[any]) AWSCloudFormationType() string {
	return "AWS::KinesisAnalyticsV2::Application.CodeContent"
}
