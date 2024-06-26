// Code generated by "go generate". Please don't change this file directly.

package iot

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// JobTemplate_JobExecutionsRolloutConfig AWS CloudFormation Resource (AWS::IoT::JobTemplate.JobExecutionsRolloutConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-jobexecutionsrolloutconfig.html
type JobTemplate_JobExecutionsRolloutConfig[T any] struct {

	// ExponentialRolloutRate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-jobexecutionsrolloutconfig.html#cfn-iot-jobtemplate-jobexecutionsrolloutconfig-exponentialrolloutrate
	ExponentialRolloutRate *JobTemplate_ExponentialRolloutRate[any] `json:"ExponentialRolloutRate,omitempty"`

	// MaximumPerMinute AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-jobexecutionsrolloutconfig.html#cfn-iot-jobtemplate-jobexecutionsrolloutconfig-maximumperminute
	MaximumPerMinute *T `json:"MaximumPerMinute,omitempty"`

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
func (r *JobTemplate_JobExecutionsRolloutConfig[any]) AWSCloudFormationType() string {
	return "AWS::IoT::JobTemplate.JobExecutionsRolloutConfig"
}
