// Code generated by "go generate". Please don't change this file directly.

package iotsitewise

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// AccessPolicy_User AWS CloudFormation Resource (AWS::IoTSiteWise::AccessPolicy.User)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-accesspolicy-user.html
type AccessPolicy_User[T any] struct {

	// id AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-accesspolicy-user.html#cfn-iotsitewise-accesspolicy-user-id
	id *T `json:"id,omitempty"`

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
func (r *AccessPolicy_User[any]) AWSCloudFormationType() string {
	return "AWS::IoTSiteWise::AccessPolicy.User"
}
