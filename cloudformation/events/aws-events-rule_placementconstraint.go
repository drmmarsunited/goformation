// Code generated by "go generate". Please don't change this file directly.

package events

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Rule_PlacementConstraint AWS CloudFormation Resource (AWS::Events::Rule.PlacementConstraint)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-placementconstraint.html
type Rule_PlacementConstraint[T any] struct {

	// Expression AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-placementconstraint.html#cfn-events-rule-placementconstraint-expression
	Expression *string `json:"Expression,omitempty"`

	// Type AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-placementconstraint.html#cfn-events-rule-placementconstraint-type
	Type *string `json:"Type,omitempty"`

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
func (r *Rule_PlacementConstraint[any]) AWSCloudFormationType() string {
	return "AWS::Events::Rule.PlacementConstraint"
}
