// Code generated by "go generate". Please don't change this file directly.

package pcaconnectorad

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// TemplateGroupAccessControlEntry_AccessRights AWS CloudFormation Resource (AWS::PCAConnectorAD::TemplateGroupAccessControlEntry.AccessRights)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-templategroupaccesscontrolentry-accessrights.html
type TemplateGroupAccessControlEntry_AccessRights[T any] struct {

	// AutoEnroll AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-templategroupaccesscontrolentry-accessrights.html#cfn-pcaconnectorad-templategroupaccesscontrolentry-accessrights-autoenroll
	AutoEnroll *T `json:"AutoEnroll,omitempty"`

	// Enroll AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-templategroupaccesscontrolentry-accessrights.html#cfn-pcaconnectorad-templategroupaccesscontrolentry-accessrights-enroll
	Enroll *T `json:"Enroll,omitempty"`

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
func (r *TemplateGroupAccessControlEntry_AccessRights[any]) AWSCloudFormationType() string {
	return "AWS::PCAConnectorAD::TemplateGroupAccessControlEntry.AccessRights"
}
