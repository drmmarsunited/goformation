// Code generated by "go generate". Please don't change this file directly.

package acmpca

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Certificate_Extensions AWS CloudFormation Resource (AWS::ACMPCA::Certificate.Extensions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-extensions.html
type Certificate_Extensions[T any] struct {

	// CertificatePolicies AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-extensions.html#cfn-acmpca-certificate-extensions-certificatepolicies
	CertificatePolicies []Certificate_PolicyInformation[any] `json:"CertificatePolicies,omitempty"`

	// CustomExtensions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-extensions.html#cfn-acmpca-certificate-extensions-customextensions
	CustomExtensions []Certificate_CustomExtension[any] `json:"CustomExtensions,omitempty"`

	// ExtendedKeyUsage AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-extensions.html#cfn-acmpca-certificate-extensions-extendedkeyusage
	ExtendedKeyUsage []Certificate_ExtendedKeyUsage[any] `json:"ExtendedKeyUsage,omitempty"`

	// KeyUsage AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-extensions.html#cfn-acmpca-certificate-extensions-keyusage
	KeyUsage *Certificate_KeyUsage[any] `json:"KeyUsage,omitempty"`

	// SubjectAlternativeNames AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-extensions.html#cfn-acmpca-certificate-extensions-subjectalternativenames
	SubjectAlternativeNames []Certificate_GeneralName[any] `json:"SubjectAlternativeNames,omitempty"`

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
func (r *Certificate_Extensions[any]) AWSCloudFormationType() string {
	return "AWS::ACMPCA::Certificate.Extensions"
}
