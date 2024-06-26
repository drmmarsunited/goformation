// Code generated by "go generate". Please don't change this file directly.

package ecs

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Service_DeploymentAlarms AWS CloudFormation Resource (AWS::ECS::Service.DeploymentAlarms)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentalarms.html
type Service_DeploymentAlarms[T any] struct {

	// AlarmNames AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentalarms.html#cfn-ecs-service-deploymentalarms-alarmnames
	AlarmNames []T `json:"AlarmNames"`

	// Enable AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentalarms.html#cfn-ecs-service-deploymentalarms-enable
	Enable T `json:"Enable"`

	// Rollback AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentalarms.html#cfn-ecs-service-deploymentalarms-rollback
	Rollback T `json:"Rollback"`

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
func (r *Service_DeploymentAlarms[any]) AWSCloudFormationType() string {
	return "AWS::ECS::Service.DeploymentAlarms"
}
