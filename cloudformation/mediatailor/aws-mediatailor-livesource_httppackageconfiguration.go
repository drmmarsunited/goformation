// Code generated by "go generate". Please don't change this file directly.

package mediatailor

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// LiveSource_HttpPackageConfiguration AWS CloudFormation Resource (AWS::MediaTailor::LiveSource.HttpPackageConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-livesource-httppackageconfiguration.html
type LiveSource_HttpPackageConfiguration[T any] struct {

	// Path AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-livesource-httppackageconfiguration.html#cfn-mediatailor-livesource-httppackageconfiguration-path
	Path T `json:"Path"`

	// SourceGroup AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-livesource-httppackageconfiguration.html#cfn-mediatailor-livesource-httppackageconfiguration-sourcegroup
	SourceGroup T `json:"SourceGroup"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-livesource-httppackageconfiguration.html#cfn-mediatailor-livesource-httppackageconfiguration-type
	Type T `json:"Type"`

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
func (r *LiveSource_HttpPackageConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::MediaTailor::LiveSource.HttpPackageConfiguration"
}
