// Code generated by "go generate". Please don't change this file directly.

package sagemaker

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// MonitoringSchedule_MonitoringInput AWS CloudFormation Resource (AWS::SageMaker::MonitoringSchedule.MonitoringInput)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringinput.html
type MonitoringSchedule_MonitoringInput[T any] struct {

	// BatchTransformInput AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringinput.html#cfn-sagemaker-monitoringschedule-monitoringinput-batchtransforminput
	BatchTransformInput *MonitoringSchedule_BatchTransformInput[any] `json:"BatchTransformInput,omitempty"`

	// EndpointInput AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringinput.html#cfn-sagemaker-monitoringschedule-monitoringinput-endpointinput
	EndpointInput *MonitoringSchedule_EndpointInput[any] `json:"EndpointInput,omitempty"`

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
func (r *MonitoringSchedule_MonitoringInput[any]) AWSCloudFormationType() string {
	return "AWS::SageMaker::MonitoringSchedule.MonitoringInput"
}
