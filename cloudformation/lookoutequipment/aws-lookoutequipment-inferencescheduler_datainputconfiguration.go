// Code generated by "go generate". Please don't change this file directly.

package lookoutequipment

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// InferenceScheduler_DataInputConfiguration AWS CloudFormation Resource (AWS::LookoutEquipment::InferenceScheduler.DataInputConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutequipment-inferencescheduler-datainputconfiguration.html
type InferenceScheduler_DataInputConfiguration[T any] struct {

	// InferenceInputNameConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutequipment-inferencescheduler-datainputconfiguration.html#cfn-lookoutequipment-inferencescheduler-datainputconfiguration-inferenceinputnameconfiguration
	InferenceInputNameConfiguration *InferenceScheduler_InputNameConfiguration[any] `json:"InferenceInputNameConfiguration,omitempty"`

	// InputTimeZoneOffset AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutequipment-inferencescheduler-datainputconfiguration.html#cfn-lookoutequipment-inferencescheduler-datainputconfiguration-inputtimezoneoffset
	InputTimeZoneOffset *T `json:"InputTimeZoneOffset,omitempty"`

	// S3InputConfiguration AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutequipment-inferencescheduler-datainputconfiguration.html#cfn-lookoutequipment-inferencescheduler-datainputconfiguration-s3inputconfiguration
	S3InputConfiguration *InferenceScheduler_S3InputConfiguration[any] `json:"S3InputConfiguration"`

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
func (r *InferenceScheduler_DataInputConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::LookoutEquipment::InferenceScheduler.DataInputConfiguration"
}
