// Code generated by "go generate". Please don't change this file directly.

package wafv2

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// RuleGroup_JsonBody AWS CloudFormation Resource (AWS::WAFv2::RuleGroup.JsonBody)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-jsonbody.html
type RuleGroup_JsonBody[T any] struct {

	// InvalidFallbackBehavior AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-jsonbody.html#cfn-wafv2-rulegroup-jsonbody-invalidfallbackbehavior
	InvalidFallbackBehavior *string `json:"InvalidFallbackBehavior,omitempty"`

	// MatchPattern AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-jsonbody.html#cfn-wafv2-rulegroup-jsonbody-matchpattern
	MatchPattern *RuleGroup_JsonMatchPattern[any] `json:"MatchPattern"`

	// MatchScope AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-jsonbody.html#cfn-wafv2-rulegroup-jsonbody-matchscope
	MatchScope string `json:"MatchScope"`

	// OversizeHandling AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-jsonbody.html#cfn-wafv2-rulegroup-jsonbody-oversizehandling
	OversizeHandling *string `json:"OversizeHandling,omitempty"`

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
func (r *RuleGroup_JsonBody[any]) AWSCloudFormationType() string {
	return "AWS::WAFv2::RuleGroup.JsonBody"
}
