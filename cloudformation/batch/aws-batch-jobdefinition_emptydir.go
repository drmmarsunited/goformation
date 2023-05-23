// Code generated by "go generate". Please don't change this file directly.

package batch

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// JobDefinition_EmptyDir AWS CloudFormation Resource (AWS::Batch::JobDefinition.EmptyDir)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-eksvolume-emptydir.html
type JobDefinition_EmptyDir struct {

	// Medium AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-eksvolume-emptydir.html#cfn-batch-jobdefinition-eksvolume-emptydir-medium
	Medium *string `json:"Medium,omitempty"`

	// SizeLimit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-eksvolume-emptydir.html#cfn-batch-jobdefinition-eksvolume-emptydir-sizelimit
	SizeLimit *string `json:"SizeLimit,omitempty"`

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
func (r *JobDefinition_EmptyDir) AWSCloudFormationType() string {
	return "AWS::Batch::JobDefinition.EmptyDir"
}
