// Code generated by "go generate". Please don't change this file directly.

package wafregional

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Rule_Predicate AWS CloudFormation Resource (AWS::WAFRegional::Rule.Predicate)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-rule-predicate.html
type Rule_Predicate[T any] struct {

	// DataId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-rule-predicate.html#cfn-wafregional-rule-predicate-dataid
	DataId string `json:"DataId"`

	// Negated AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-rule-predicate.html#cfn-wafregional-rule-predicate-negated
	Negated T `json:"Negated"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-rule-predicate.html#cfn-wafregional-rule-predicate-type
	Type string `json:"Type"`

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
func (r *Rule_Predicate[any]) AWSCloudFormationType() string {
	return "AWS::WAFRegional::Rule.Predicate"
}
