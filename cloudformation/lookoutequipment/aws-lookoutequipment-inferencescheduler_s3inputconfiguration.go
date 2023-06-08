// Code generated by "go generate". Please don't change this file directly.

package lookoutequipment

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// InferenceScheduler_S3InputConfiguration AWS CloudFormation Resource (AWS::LookoutEquipment::InferenceScheduler.S3InputConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutequipment-inferencescheduler-s3inputconfiguration.html
type InferenceScheduler_S3InputConfiguration[T any] struct {

	// Bucket AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutequipment-inferencescheduler-s3inputconfiguration.html#cfn-lookoutequipment-inferencescheduler-s3inputconfiguration-bucket
	Bucket string `json:"Bucket"`

	// Prefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutequipment-inferencescheduler-s3inputconfiguration.html#cfn-lookoutequipment-inferencescheduler-s3inputconfiguration-prefix
	Prefix *string `json:"Prefix,omitempty"`

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
func (r *InferenceScheduler_S3InputConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::LookoutEquipment::InferenceScheduler.S3InputConfiguration"
}
