// Code generated by "go generate". Please don't change this file directly.

package wafv2

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// RuleGroup_Rule AWS CloudFormation Resource (AWS::WAFv2::RuleGroup.Rule)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-rule.html
type RuleGroup_Rule[T any] struct {

	// Action AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-rule.html#cfn-wafv2-rulegroup-rule-action
	Action *RuleGroup_RuleAction[any] `json:"Action,omitempty"`

	// CaptchaConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-rule.html#cfn-wafv2-rulegroup-rule-captchaconfig
	CaptchaConfig *RuleGroup_CaptchaConfig[any] `json:"CaptchaConfig,omitempty"`

	// ChallengeConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-rule.html#cfn-wafv2-rulegroup-rule-challengeconfig
	ChallengeConfig *RuleGroup_ChallengeConfig[any] `json:"ChallengeConfig,omitempty"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-rule.html#cfn-wafv2-rulegroup-rule-name
	Name T `json:"Name"`

	// Priority AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-rule.html#cfn-wafv2-rulegroup-rule-priority
	Priority T `json:"Priority"`

	// RuleLabels AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-rule.html#cfn-wafv2-rulegroup-rule-rulelabels
	RuleLabels []RuleGroup_Label[any] `json:"RuleLabels,omitempty"`

	// Statement AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-rule.html#cfn-wafv2-rulegroup-rule-statement
	Statement *RuleGroup_Statement[any] `json:"Statement"`

	// VisibilityConfig AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-rule.html#cfn-wafv2-rulegroup-rule-visibilityconfig
	VisibilityConfig *RuleGroup_VisibilityConfig[any] `json:"VisibilityConfig"`

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
func (r *RuleGroup_Rule[any]) AWSCloudFormationType() string {
	return "AWS::WAFv2::RuleGroup.Rule"
}
