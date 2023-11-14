// Code generated by "go generate". Please don't change this file directly.

package mediatailor

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// SourceLocation_HttpConfiguration AWS CloudFormation Resource (AWS::MediaTailor::SourceLocation.HttpConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-sourcelocation-httpconfiguration.html
type SourceLocation_HttpConfiguration[T any] struct {

	// BaseUrl AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-sourcelocation-httpconfiguration.html#cfn-mediatailor-sourcelocation-httpconfiguration-baseurl
	BaseUrl T `json:"BaseUrl"`

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
func (r *SourceLocation_HttpConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::MediaTailor::SourceLocation.HttpConfiguration"
}
