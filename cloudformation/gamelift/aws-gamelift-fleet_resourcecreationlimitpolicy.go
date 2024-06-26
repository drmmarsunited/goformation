// Code generated by "go generate". Please don't change this file directly.

package gamelift

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Fleet_ResourceCreationLimitPolicy AWS CloudFormation Resource (AWS::GameLift::Fleet.ResourceCreationLimitPolicy)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-resourcecreationlimitpolicy.html
type Fleet_ResourceCreationLimitPolicy[T any] struct {

	// NewGameSessionsPerCreator AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-resourcecreationlimitpolicy.html#cfn-gamelift-fleet-resourcecreationlimitpolicy-newgamesessionspercreator
	NewGameSessionsPerCreator *T `json:"NewGameSessionsPerCreator,omitempty"`

	// PolicyPeriodInMinutes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-resourcecreationlimitpolicy.html#cfn-gamelift-fleet-resourcecreationlimitpolicy-policyperiodinminutes
	PolicyPeriodInMinutes *T `json:"PolicyPeriodInMinutes,omitempty"`

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
func (r *Fleet_ResourceCreationLimitPolicy[any]) AWSCloudFormationType() string {
	return "AWS::GameLift::Fleet.ResourceCreationLimitPolicy"
}
