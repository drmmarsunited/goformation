// Code generated by "go generate". Please don't change this file directly.

package pcaconnectorad

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_EnrollmentFlagsV3 AWS CloudFormation Resource (AWS::PCAConnectorAD::Template.EnrollmentFlagsV3)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-enrollmentflagsv3.html
type Template_EnrollmentFlagsV3[T any] struct {

	// EnableKeyReuseOnNtTokenKeysetStorageFull AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-enrollmentflagsv3.html#cfn-pcaconnectorad-template-enrollmentflagsv3-enablekeyreuseonnttokenkeysetstoragefull
	EnableKeyReuseOnNtTokenKeysetStorageFull *T `json:"EnableKeyReuseOnNtTokenKeysetStorageFull,omitempty"`

	// IncludeSymmetricAlgorithms AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-enrollmentflagsv3.html#cfn-pcaconnectorad-template-enrollmentflagsv3-includesymmetricalgorithms
	IncludeSymmetricAlgorithms *T `json:"IncludeSymmetricAlgorithms,omitempty"`

	// NoSecurityExtension AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-enrollmentflagsv3.html#cfn-pcaconnectorad-template-enrollmentflagsv3-nosecurityextension
	NoSecurityExtension *T `json:"NoSecurityExtension,omitempty"`

	// RemoveInvalidCertificateFromPersonalStore AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-enrollmentflagsv3.html#cfn-pcaconnectorad-template-enrollmentflagsv3-removeinvalidcertificatefrompersonalstore
	RemoveInvalidCertificateFromPersonalStore *T `json:"RemoveInvalidCertificateFromPersonalStore,omitempty"`

	// UserInteractionRequired AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-enrollmentflagsv3.html#cfn-pcaconnectorad-template-enrollmentflagsv3-userinteractionrequired
	UserInteractionRequired *T `json:"UserInteractionRequired,omitempty"`

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
func (r *Template_EnrollmentFlagsV3[any]) AWSCloudFormationType() string {
	return "AWS::PCAConnectorAD::Template.EnrollmentFlagsV3"
}
