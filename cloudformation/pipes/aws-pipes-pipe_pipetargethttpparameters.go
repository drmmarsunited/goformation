// Code generated by "go generate". Please don't change this file directly.

package pipes

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Pipe_PipeTargetHttpParameters AWS CloudFormation Resource (AWS::Pipes::Pipe.PipeTargetHttpParameters)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargethttpparameters.html
type Pipe_PipeTargetHttpParameters[T any] struct {

	// HeaderParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargethttpparameters.html#cfn-pipes-pipe-pipetargethttpparameters-headerparameters
	HeaderParameters map[string]T `json:"HeaderParameters,omitempty"`

	// PathParameterValues AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargethttpparameters.html#cfn-pipes-pipe-pipetargethttpparameters-pathparametervalues
	PathParameterValues []T `json:"PathParameterValues,omitempty"`

	// QueryStringParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargethttpparameters.html#cfn-pipes-pipe-pipetargethttpparameters-querystringparameters
	QueryStringParameters map[string]T `json:"QueryStringParameters,omitempty"`

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
func (r *Pipe_PipeTargetHttpParameters[any]) AWSCloudFormationType() string {
	return "AWS::Pipes::Pipe.PipeTargetHttpParameters"
}
