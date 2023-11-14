// Code generated by "go generate". Please don't change this file directly.

package acmpca

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// CertificateAuthority_OtherName AWS CloudFormation Resource (AWS::ACMPCA::CertificateAuthority.OtherName)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-othername.html
type CertificateAuthority_OtherName[T any] struct {

	// TypeId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-othername.html#cfn-acmpca-certificateauthority-othername-typeid
	TypeId T `json:"TypeId"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-othername.html#cfn-acmpca-certificateauthority-othername-value
	Value T `json:"Value"`

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
func (r *CertificateAuthority_OtherName[any]) AWSCloudFormationType() string {
	return "AWS::ACMPCA::CertificateAuthority.OtherName"
}
