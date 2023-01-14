// Code generated by "go generate". Please don't change this file directly.

package iot

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// JobTemplate_RetryCriteria AWS CloudFormation Resource (AWS::IoT::JobTemplate.RetryCriteria)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-retrycriteria.html
type JobTemplate_RetryCriteria struct {

	// FailureType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-retrycriteria.html#cfn-iot-jobtemplate-retrycriteria-failuretype
	FailureType *string `json:"FailureType,omitempty"`

	// NumberOfRetries AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-retrycriteria.html#cfn-iot-jobtemplate-retrycriteria-numberofretries
	NumberOfRetries *int `json:"NumberOfRetries,omitempty"`

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
func (r *JobTemplate_RetryCriteria) AWSCloudFormationType() string {
	return "AWS::IoT::JobTemplate.RetryCriteria"
}
