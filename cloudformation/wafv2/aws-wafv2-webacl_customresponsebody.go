// Code generated by "go generate". Please don't change this file directly.

package wafv2

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// WebACL_CustomResponseBody AWS CloudFormation Resource (AWS::WAFv2::WebACL.CustomResponseBody)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-customresponsebody.html
type WebACL_CustomResponseBody[T any] struct {

	// Content AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-customresponsebody.html#cfn-wafv2-webacl-customresponsebody-content
	Content T `json:"Content"`

	// ContentType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-customresponsebody.html#cfn-wafv2-webacl-customresponsebody-contenttype
	ContentType T `json:"ContentType"`

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
func (r *WebACL_CustomResponseBody[any]) AWSCloudFormationType() string {
	return "AWS::WAFv2::WebACL.CustomResponseBody"
}
