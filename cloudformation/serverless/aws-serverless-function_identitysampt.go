// Code generated by "go generate". Please don't change this file directly.

package serverless

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Function_IdentitySAMPT AWS CloudFormation Resource (AWS::Serverless::Function.IdentitySAMPT)
// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
type Function_IdentitySAMPT[T any] struct {

	// IdentityName AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	IdentityName T `json:"IdentityName"`

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
func (r *Function_IdentitySAMPT[any]) AWSCloudFormationType() string {
	return "AWS::Serverless::Function.IdentitySAMPT"
}
