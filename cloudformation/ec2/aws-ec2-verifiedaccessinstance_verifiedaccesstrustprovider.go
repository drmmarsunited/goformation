// Code generated by "go generate". Please don't change this file directly.

package ec2

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// VerifiedAccessInstance_VerifiedAccessTrustProvider AWS CloudFormation Resource (AWS::EC2::VerifiedAccessInstance.VerifiedAccessTrustProvider)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessinstance-verifiedaccesstrustprovider.html
type VerifiedAccessInstance_VerifiedAccessTrustProvider[T any] struct {

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessinstance-verifiedaccesstrustprovider.html#cfn-ec2-verifiedaccessinstance-verifiedaccesstrustprovider-description
	Description *T `json:"Description,omitempty"`

	// DeviceTrustProviderType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessinstance-verifiedaccesstrustprovider.html#cfn-ec2-verifiedaccessinstance-verifiedaccesstrustprovider-devicetrustprovidertype
	DeviceTrustProviderType *T `json:"DeviceTrustProviderType,omitempty"`

	// TrustProviderType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessinstance-verifiedaccesstrustprovider.html#cfn-ec2-verifiedaccessinstance-verifiedaccesstrustprovider-trustprovidertype
	TrustProviderType *T `json:"TrustProviderType,omitempty"`

	// UserTrustProviderType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessinstance-verifiedaccesstrustprovider.html#cfn-ec2-verifiedaccessinstance-verifiedaccesstrustprovider-usertrustprovidertype
	UserTrustProviderType *T `json:"UserTrustProviderType,omitempty"`

	// VerifiedAccessTrustProviderId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessinstance-verifiedaccesstrustprovider.html#cfn-ec2-verifiedaccessinstance-verifiedaccesstrustprovider-verifiedaccesstrustproviderid
	VerifiedAccessTrustProviderId *T `json:"VerifiedAccessTrustProviderId,omitempty"`

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
func (r *VerifiedAccessInstance_VerifiedAccessTrustProvider[any]) AWSCloudFormationType() string {
	return "AWS::EC2::VerifiedAccessInstance.VerifiedAccessTrustProvider"
}
