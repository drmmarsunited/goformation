// Code generated by "go generate". Please don't change this file directly.

package datapipeline

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Pipeline_ParameterValue AWS CloudFormation Resource (AWS::DataPipeline::Pipeline.ParameterValue)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parametervalue.html
type Pipeline_ParameterValue struct {

	// Id AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parametervalue.html#cfn-datapipeline-pipeline-parametervalue-id
	Id string `json:"Id"`

	// StringValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parametervalue.html#cfn-datapipeline-pipeline-parametervalue-stringvalue
	StringValue string `json:"StringValue"`

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
func (r *Pipeline_ParameterValue) AWSCloudFormationType() string {
	return "AWS::DataPipeline::Pipeline.ParameterValue"
}
