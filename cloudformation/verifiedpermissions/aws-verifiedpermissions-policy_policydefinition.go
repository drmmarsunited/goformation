// Code generated by "go generate". Please don't change this file directly.

package verifiedpermissions

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Policy_PolicyDefinition AWS CloudFormation Resource (AWS::VerifiedPermissions::Policy.PolicyDefinition)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policy-policydefinition.html
type Policy_PolicyDefinition[T any] struct {

	// Static AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policy-policydefinition.html#cfn-verifiedpermissions-policy-policydefinition-static
	Static *Policy_StaticPolicyDefinition[any] `json:"Static,omitempty"`

	// TemplateLinked AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policy-policydefinition.html#cfn-verifiedpermissions-policy-policydefinition-templatelinked
	TemplateLinked *Policy_TemplateLinkedPolicyDefinition[any] `json:"TemplateLinked,omitempty"`

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
func (r *Policy_PolicyDefinition[any]) AWSCloudFormationType() string {
	return "AWS::VerifiedPermissions::Policy.PolicyDefinition"
}
