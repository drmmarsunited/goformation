// Code generated by "go generate". Please don't change this file directly.

package verifiedpermissions

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Policy_EntityIdentifier AWS CloudFormation Resource (AWS::VerifiedPermissions::Policy.EntityIdentifier)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policy-entityidentifier.html
type Policy_EntityIdentifier[T any] struct {

	// EntityId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policy-entityidentifier.html#cfn-verifiedpermissions-policy-entityidentifier-entityid
	EntityId T `json:"EntityId"`

	// EntityType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policy-entityidentifier.html#cfn-verifiedpermissions-policy-entityidentifier-entitytype
	EntityType T `json:"EntityType"`

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
func (r *Policy_EntityIdentifier[any]) AWSCloudFormationType() string {
	return "AWS::VerifiedPermissions::Policy.EntityIdentifier"
}
