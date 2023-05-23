// Code generated by "go generate". Please don't change this file directly.

package s3

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// AccessPoint_PolicyStatus AWS CloudFormation Resource (AWS::S3::AccessPoint.PolicyStatus)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-accesspoint-policystatus.html
type AccessPoint_PolicyStatus struct {

	// IsPublic AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-accesspoint-policystatus.html#cfn-s3-accesspoint-policystatus-ispublic
	IsPublic *string `json:"IsPublic,omitempty"`

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
func (r *AccessPoint_PolicyStatus) AWSCloudFormationType() string {
	return "AWS::S3::AccessPoint.PolicyStatus"
}
