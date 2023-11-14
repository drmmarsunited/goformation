// Code generated by "go generate". Please don't change this file directly.

package pcaconnectorad

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_PrivateKeyFlagsV4 AWS CloudFormation Resource (AWS::PCAConnectorAD::Template.PrivateKeyFlagsV4)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyflagsv4.html
type Template_PrivateKeyFlagsV4[T any] struct {

	// ClientVersion AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyflagsv4.html#cfn-pcaconnectorad-template-privatekeyflagsv4-clientversion
	ClientVersion T `json:"ClientVersion"`

	// ExportableKey AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyflagsv4.html#cfn-pcaconnectorad-template-privatekeyflagsv4-exportablekey
	ExportableKey *T `json:"ExportableKey,omitempty"`

	// RequireAlternateSignatureAlgorithm AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyflagsv4.html#cfn-pcaconnectorad-template-privatekeyflagsv4-requirealternatesignaturealgorithm
	RequireAlternateSignatureAlgorithm *T `json:"RequireAlternateSignatureAlgorithm,omitempty"`

	// RequireSameKeyRenewal AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyflagsv4.html#cfn-pcaconnectorad-template-privatekeyflagsv4-requiresamekeyrenewal
	RequireSameKeyRenewal *T `json:"RequireSameKeyRenewal,omitempty"`

	// StrongKeyProtectionRequired AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyflagsv4.html#cfn-pcaconnectorad-template-privatekeyflagsv4-strongkeyprotectionrequired
	StrongKeyProtectionRequired *T `json:"StrongKeyProtectionRequired,omitempty"`

	// UseLegacyProvider AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyflagsv4.html#cfn-pcaconnectorad-template-privatekeyflagsv4-uselegacyprovider
	UseLegacyProvider *T `json:"UseLegacyProvider,omitempty"`

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
func (r *Template_PrivateKeyFlagsV4[any]) AWSCloudFormationType() string {
	return "AWS::PCAConnectorAD::Template.PrivateKeyFlagsV4"
}
