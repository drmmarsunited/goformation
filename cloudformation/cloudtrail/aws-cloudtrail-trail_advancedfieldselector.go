// Code generated by "go generate". Please don't change this file directly.

package cloudtrail

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Trail_AdvancedFieldSelector AWS CloudFormation Resource (AWS::CloudTrail::Trail.AdvancedFieldSelector)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-advancedfieldselector.html
type Trail_AdvancedFieldSelector[T any] struct {

	// EndsWith AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-advancedfieldselector.html#cfn-cloudtrail-trail-advancedfieldselector-endswith
	EndsWith []T `json:"EndsWith,omitempty"`

	// Equals AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-advancedfieldselector.html#cfn-cloudtrail-trail-advancedfieldselector-equals
	Equals []T `json:"Equals,omitempty"`

	// Field AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-advancedfieldselector.html#cfn-cloudtrail-trail-advancedfieldselector-field
	Field T `json:"Field"`

	// NotEndsWith AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-advancedfieldselector.html#cfn-cloudtrail-trail-advancedfieldselector-notendswith
	NotEndsWith []T `json:"NotEndsWith,omitempty"`

	// NotEquals AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-advancedfieldselector.html#cfn-cloudtrail-trail-advancedfieldselector-notequals
	NotEquals []T `json:"NotEquals,omitempty"`

	// NotStartsWith AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-advancedfieldselector.html#cfn-cloudtrail-trail-advancedfieldselector-notstartswith
	NotStartsWith []T `json:"NotStartsWith,omitempty"`

	// StartsWith AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-advancedfieldselector.html#cfn-cloudtrail-trail-advancedfieldselector-startswith
	StartsWith []T `json:"StartsWith,omitempty"`

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
func (r *Trail_AdvancedFieldSelector[any]) AWSCloudFormationType() string {
	return "AWS::CloudTrail::Trail.AdvancedFieldSelector"
}
