// Code generated by "go generate". Please don't change this file directly.

package acmpca

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// CertificateAuthority_CsrExtensions AWS CloudFormation Resource (AWS::ACMPCA::CertificateAuthority.CsrExtensions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-csrextensions.html
type CertificateAuthority_CsrExtensions[T any] struct {

	// KeyUsage AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-csrextensions.html#cfn-acmpca-certificateauthority-csrextensions-keyusage
	KeyUsage *CertificateAuthority_KeyUsage[any] `json:"KeyUsage,omitempty"`

	// SubjectInformationAccess AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-csrextensions.html#cfn-acmpca-certificateauthority-csrextensions-subjectinformationaccess
	SubjectInformationAccess []CertificateAuthority_AccessDescription[any] `json:"SubjectInformationAccess,omitempty"`

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
func (r *CertificateAuthority_CsrExtensions[any]) AWSCloudFormationType() string {
	return "AWS::ACMPCA::CertificateAuthority.CsrExtensions"
}
