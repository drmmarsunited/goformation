// Code generated by "go generate". Please don't change this file directly.

package entityresolution

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// MatchingWorkflow_Rule AWS CloudFormation Resource (AWS::EntityResolution::MatchingWorkflow.Rule)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-rule.html
type MatchingWorkflow_Rule[T any] struct {

	// MatchingKeys AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-rule.html#cfn-entityresolution-matchingworkflow-rule-matchingkeys
	MatchingKeys []T `json:"MatchingKeys"`

	// RuleName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-rule.html#cfn-entityresolution-matchingworkflow-rule-rulename
	RuleName T `json:"RuleName"`

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
func (r *MatchingWorkflow_Rule[any]) AWSCloudFormationType() string {
	return "AWS::EntityResolution::MatchingWorkflow.Rule"
}